package raft

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

const (
	Follower  = "Follower"
	Candidate = "Candidate"
	Leader    = "Leader"
)

type Node struct {
	mu               sync.Mutex
	id               string
	state            string
	term             int
	votedFor         string
	votes            int
	peers            []string
	transport        *Transport
	electionTimeout  time.Duration
	heartbeatTimeout time.Duration
	electionTimer    *time.Timer
	heartbeatTimer   *time.Timer
}

func NewRaftNode(id string, port int, peers []string) *Node {
	node := &Node{
		id:               id,
		state:            Follower,
		term:             0,
		peers:            peers,
		transport:        NewTransport(port),
		electionTimeout:  time.Duration(rand.Intn(150)+300) * time.Millisecond,
		heartbeatTimeout: 150 * time.Millisecond,
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
			if n.state == Follower || n.state == Candidate {
				n.startElection()
			}
			n.mu.Unlock()
		case <-n.heartbeatTimer.C:
			n.mu.Lock()
			if n.state == Leader {
				n.heartbeat()
			}
			n.mu.Unlock()

		}

	}
}

func (n *Node) startElection() {
	n.state = Candidate
	n.term++
	n.votedFor = n.id
	n.votes = 1
	fmt.Printf("[%s] Starting election for term %d\n", n.id, n.term)
	n.broadcast(&Message{
		Type:      RequestVote,
		Term:      n.term,
		Candidate: n.id,
	})
	n.electionTimer.Reset(n.electionTimeout)
}

func (n *Node) heartbeat() {
	if n.state != Leader {
		return
	}

	n.broadcast(&Message{
		Type: AppendEntries,
		Term: n.term,
	})

	n.heartbeatTimer.Reset(n.heartbeatTimeout)

}

func (n *Node) listen() {
	n.transport.Receive(func(msg *Message, addr *net.UDPAddr) {
		n.handleMessage(msg, addr)
	})
}

func (n *Node) handleMessage(msg *Message, addr *net.UDPAddr) {
	n.mu.Lock()
	defer n.mu.Unlock()

	switch msg.Type {
	case RequestVote:
		fmt.Printf("[%s] Received RequestVote from %s\n", n.id, addr)
		if msg.Term > n.term {

			n.term = msg.Term
			n.state = Follower
			n.votedFor = ""
		}
		if (n.votedFor == "" || n.votedFor == msg.Candidate) && msg.Term >= n.term {
			n.votedFor = msg.Candidate
			n.transport.Send(addr.String(), &Message{Type: VoteGranted, Term: n.term})
		}
	case VoteGranted:
		fmt.Printf("[%s] Received VoteGranted from %s\n", n.id, addr)
		if n.state == Candidate && msg.Term == n.term {
			n.votes++
			if n.votes > len(n.peers)/2 {
				fmt.Printf("[%s] Became leader for term %d\n", n.id, n.term)
				n.state = Leader
				n.heartbeat()
			}
		}
	case AppendEntries:
		fmt.Printf("[%s] Received AppendEntries from %s\n", n.id, addr)

		if msg.Term >= n.term {
			n.state = Follower
			n.term = msg.Term
			n.electionTimer.Reset(n.electionTimeout)
		}
	}

}

func (n *Node) broadcast(msg *Message) {
	for _, peer := range n.peers {
		n.transport.Send(peer, msg)
	}
}

func (n *Node) Shutdown() {
	n.mu.Lock()
	defer n.mu.Unlock()

	fmt.Printf("[%s] Shutting down\n", n.id)

	n.electionTimer.Stop()
	n.heartbeatTimer.Stop()

	n.transport.Close()

	n.state = "Shutdown"
}
