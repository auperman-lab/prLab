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
		ID:    ID,
		peers: peers,
		state: NewState(),
		log:   NewLog(),
		//transport:        NewTransport(port, "http://localhost:8080"),
		transport:        NewTransport(port, "http://manager:8080"),
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
			Payload: &AppendEntriesRequest{
				Term:         n.state.GetTerm(),
				LeaderID:     n.ID,
				PrevLogIndex: n.log.LastIndex(),
				PrevLogTerm:  n.log.LastTerm(),
				Entries:      make([]LogEntry, 0),
				LeaderCommit: n.state.commitIndex,
			},
		},
		n.peers,
	)

	n.heartbeatTimer.Reset(n.heartbeatTimeout)
}
func (n *Node) SendEntries(entries []LogEntry) {

	if n.state.GetLeader() != n.ID {
		return
	}

	prevLogIndex := n.log.LastIndex() - len(entries)

	prevLogTerm := 0
	if prevLogIndex >= 0 {
		prevLogTerm = n.log.LastTerm()
	} else {
		prevLogIndex = -1
	}

	m := AppendEntriesRequest{
		Term:         n.state.GetTerm(),
		LeaderID:     n.ID,
		PrevLogIndex: prevLogIndex,
		PrevLogTerm:  prevLogTerm,
		Entries:      entries,
		LeaderCommit: n.state.commitIndex,
	}
	fmt.Printf("\033[31m  message : %+v \033[0m\n", m)

	n.transport.Broadcast(
		&Message{
			Type:    AppendEntriesReq,
			Payload: m,
		},
		n.peers,
	)

}
func (n *Node) checkCommitIndex() {
	n.state.mu.Lock()
	defer n.state.mu.Unlock()

	matchIndices := make([]int, 0, len(n.state.matchIndex)+1)
	for _, matchIdx := range n.state.matchIndex {
		matchIndices = append(matchIndices, matchIdx)
	}
	matchIndices = append(matchIndices, n.log.LastIndex())

	sort.Ints(matchIndices)
	newCommitIndex := matchIndices[len(matchIndices)/2]

	if newCommitIndex > n.state.commitIndex {
		entry, _ := n.log.Get(newCommitIndex)
		if entry.Term == n.state.GetTerm() {
			n.state.commitIndex = newCommitIndex

			fmt.Printf("[%s] Log committed up to index %d\n", n.ID, newCommitIndex)
		}
	}
}

func (n *Node) retryAppendEntries(followerAddr string) {
	n.state.mu.Lock()
	defer n.state.mu.Unlock()

	nextIndex := n.state.nextIndex[followerAddr]
	prevLogIndex := nextIndex - 1
	prevLogTerm := 0
	if prevLogIndex >= 0 {
		entry, _ := n.log.Get(prevLogIndex)
		prevLogTerm = entry.Term
	}

	entries, _ := n.log.Get(nextIndex)

	n.transport.Send(
		followerAddr,
		&Message{
			Type: AppendEntriesReq,
			Payload: AppendEntriesRequest{
				Term:         n.state.GetTerm(),
				LeaderID:     n.ID,
				PrevLogIndex: prevLogIndex,
				PrevLogTerm:  prevLogTerm,
				Entries:      []LogEntry{entries},
				LeaderCommit: n.state.GetCommitIndex(),
			},
		},
	)
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

		if payload.Term > n.state.GetTerm() {
			n.state.SetTerm(payload.Term)
			n.convertToFollower()
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

		if payload.Term > n.state.GetTerm() {
			n.state.SetTerm(payload.Term)
			n.convertToFollower()
			return
		}

		if !payload.Success {
			followerAddr := addr.String()
			n.state.nextIndex[followerAddr] = max(n.state.nextIndex[followerAddr]-1, 1)
			fmt.Printf("[%s] Log inconsistency with %s, retrying...\n", n.ID, followerAddr)
			n.retryAppendEntries(followerAddr)
			return
		}

		followerAddr := addr.String()
		n.state.matchIndex[followerAddr] = n.state.nextIndex[followerAddr] - 1
		n.state.nextIndex[followerAddr] = n.state.matchIndex[followerAddr] + 1

		n.checkCommitIndex()

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

	for _, peer := range n.peers {
		n.state.matchIndex[peer] = 0
		n.state.nextIndex[peer] = n.log.LastIndex() + 1
	}

	n.transport.SendLeaderAddr(n.ID)

	n.sendHeartbeats()
}

func (n *Node) GetLeader() bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	fmt.Printf("\033[31m[%s]Leader Value: %s\033[0m\n", n.ID, n.state.leader)

	return n.state.leader == n.ID

}
func (n *Node) AppendLogs(data []byte) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	logEntry := LogEntry{
		Index: n.log.LastIndex() + 1,
		Term:  n.state.GetTerm(),
		Data:  data,
	}
	n.log.Append(logEntry)

	fmt.Printf("\033[31m node is sending the entries \033[0m\n")

	n.SendEntries([]LogEntry{logEntry})

	return nil
}
