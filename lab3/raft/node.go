package raft

import (
	"fmt"
	"math/rand"
	"net"
	"sort"
	"sync"
	"time"
)

type Node struct {
	ID               string
	peers            []string
	state            *State
	nodeType         NodeType
	log              *Log
	transport        *Transport
	heartbeatTimer   *time.Timer
	electionTimer    *time.Timer
	electionTimeout  time.Duration
	heartbeatTimeout time.Duration
	votes            int
	mu               sync.Mutex
}

type NodeType string

const (
	Follower  NodeType = "Follower"
	Candidate NodeType = "Candidate"
	Leader    NodeType = "Leader"
)

func NewNode(ID string, port int, peers []string) *Node {
	node := &Node{
		ID:               ID,
		peers:            peers,
		state:            NewState(),
		log:              NewLog(),
		transport:        NewTransport(port),
		electionTimeout:  time.Duration(rand.Intn(150)+300) * time.Millisecond,
		heartbeatTimeout: 150 * time.Millisecond,
		votes:            0,
		nodeType:         Follower,
	}
	node.electionTimer = time.NewTimer(node.electionTimeout)
	node.heartbeatTimer = time.NewTimer(node.heartbeatTimeout)

	go node.startTimer()
	go node.listen()

	return node
}

func (n *Node) startTimer() {
	for {
		select {
		case <-n.electionTimer.C:
			n.mu.Lock()
			if n.nodeType == Follower || n.nodeType == Candidate {
				n.convertToCandidate()
			}
			n.mu.Unlock()
		case <-n.heartbeatTimer.C:
			if n.nodeType == Leader {
				n.sendHeartbeats()
			}

		}

	}
}

func (n *Node) startElection() {

	n.transport.Broadcast(
		&Message{
			Type: VoteReq,
			Payload: VoteRequest{
				Term:         n.state.GetTerm(),
				CandidateID:  n.ID,
				LastLogIndex: n.log.LastIndex(),
				LastLogTerm:  n.log.LastTerm(),
			},
		},
		n.peers,
	)
}

func (n *Node) listen() {
	n.transport.Receive(func(msg *Message, addr *net.UDPAddr) {
		n.handleMessage(msg, addr)
	})
}

func (n *Node) RequestVote(candidateID string, candidateTerm int, lastLogIndex int, lastLogTerm int) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	if candidateTerm < n.state.GetTerm() {
		return false
	}

	if n.state.GetVotedFor() == "" || n.state.GetVotedFor() == candidateID {
		if lastLogTerm > n.log.LastTerm() || (lastLogTerm == n.log.LastTerm() && lastLogIndex >= n.log.LastIndex()) {
			n.state.SetVotedFor(candidateID)
			return true
		}
	}

	return false
}

func (n *Node) AppendEntries(term int, leaderID string, prevLogIndex int, prevLogTerm int, entries []LogEntry, leaderCommit int) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	if term < n.state.GetTerm() {
		return false
	}

	n.state.SetTerm(term)
	n.state.SetLeader(leaderID)
	n.nodeType = Follower
	n.electionTimer.Reset(n.electionTimeout)

	if prevLogIndex >= 0 {
		entry, exists := n.log.Get(prevLogIndex)
		if !exists || entry.Term != prevLogTerm {
			return false
		}
	}

	for _, entry := range entries {
		n.log.Append(entry)
	}

	if leaderCommit > n.log.LastIndex() {
		commitIndex := min(leaderCommit, n.log.LastIndex())
		n.state.SetCommitIndex(commitIndex)
	}

	return true
}

func (n *Node) sendHeartbeats() {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.state.GetLeader() != n.ID {
		return
	}

	n.transport.Broadcast(
		&Message{
			Type: AppendEntriesReq,
			Payload: AppendEntriesRequest{
				Term:         n.state.GetTerm(),
				LeaderID:     n.ID,
				PrevLogIndex: n.log.LastIndex(),
				PrevLogTerm:  n.log.LastTerm(),
				Entries:      make([]LogEntry, 0),
				LeaderCommit: 0,
			},
		},
		n.peers,
	)

	n.heartbeatTimer.Reset(n.heartbeatTimeout)
}

func (n *Node) handleMessage(msg *Message, addr *net.UDPAddr) {

	switch msg.Type {
	case VoteReq:
		fmt.Printf("[%s] Received RequestVote from %s\n", n.ID, addr)
		payload, err := DeserializeToVoteRequest(msg.Payload)
		if err != nil {
			fmt.Printf("[%s] Error deserializing payload: %s\n", n.ID, err)
		}
		ok := n.RequestVote(payload.CandidateID, payload.Term, payload.LastLogIndex, payload.LastLogTerm)
		n.transport.Send(
			addr.String(),
			&Message{
				Type: VoteResp,
				Payload: VoteResponse{
					Term:        payload.Term,
					VoteGranted: ok,
				},
			},
		)
	case VoteResp:
		fmt.Printf("[%s] Received Vote Response from %s\n", n.ID, addr)
		payload, err := DeserializeToVoteResponse(msg.Payload)
		if err != nil {
			fmt.Printf("[%s] Error deserializing payload: %s\n", n.ID, err)
		}

		if n.nodeType == Candidate && payload.Term == n.state.term && payload.VoteGranted {
			fmt.Printf("[%s] VoteGranted from %s\n", n.ID, addr)

			n.votes++
			if n.votes > len(n.peers)/2 {
				fmt.Printf("[%s] Became leader for term %d\n", n.ID, n.state.term)
				n.convertToLeader()
			}
		}
	case AppendEntriesReq:
		fmt.Printf("[%s] Received AppendEntries Request from %s\n", n.ID, addr)
		payload, err := DeserializeToAppendRequest(msg.Payload)
		if err != nil {
			fmt.Printf("[%s] Error deserializing payload: %s\n", n.ID, err)
		}

		ok := n.AppendEntries(payload.Term, payload.LeaderID, payload.PrevLogIndex, payload.PrevLogTerm, payload.Entries, payload.LeaderCommit)
		if !ok {
			n.convertToFollower()
		}

		n.transport.Send(
			addr.String(),
			&Message{
				Type: AppendEntriesResp,
				Payload: AppendEntriesResponse{
					Term:    payload.Term,
					Success: ok,
				},
			},
		)
	case AppendEntriesResp:
		fmt.Printf("[%s] Received AppendEntries Response from %s\n", n.ID, addr)
		payload, err := DeserializeToAppendResponse(msg.Payload)
		if err != nil {
			fmt.Printf("[%s] Error deserializing payload: %s\n", n.ID, err)
		}

		if !payload.Success {
			fmt.Printf("[%s] Error from  %s\n", n.ID, addr)
		}

	}

}

func (n *Node) advanceCommitIndex() {
	matchIndices := make([]int, 0, len(n.peers)+1)
	for _, index := range n.state.matchIndex {
		matchIndices = append(matchIndices, index)
	}
	matchIndices = append(matchIndices, n.log.LastIndex())

	sort.Ints(matchIndices)
	newCommitIndex := matchIndices[len(matchIndices)/2]

	if newCommitIndex > n.state.GetCommitIndex() {
		entry, exists := n.log.Get(newCommitIndex)

		if exists && entry.Term == n.state.GetTerm() {
			n.state.SetCommitIndex(newCommitIndex)
		}
	}
}

func (n *Node) convertToCandidate() {
	n.nodeType = Candidate
	n.state.SetTerm(n.state.GetTerm() + 1)
	n.votes = 1
	n.state.SetVotedFor(n.ID)
	n.electionTimer.Reset(n.electionTimeout)
	n.startElection()

}
func (n *Node) convertToFollower() {
	n.nodeType = Follower
	n.electionTimer.Reset(n.electionTimeout)
	n.votes = 0

}
func (n *Node) convertToLeader() {
	n.nodeType = Leader
	n.state.SetLeader(n.ID)

	n.sendHeartbeats()
}
