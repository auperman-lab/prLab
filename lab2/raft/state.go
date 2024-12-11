package raft

import (
	"sync"
)

type State struct {
	mu sync.Mutex
	//persistent
	term     int
	votedFor string
	leader   string
	//volatile
	commitIndex int
	lastApplied int
	//volatile leader
	matchIndex map[string]int
	nextIndex  map[string]int
}

func NewState() *State {
	return &State{
		term:        0,
		votedFor:    "",
		leader:      "",
		commitIndex: 0,
		lastApplied: 0,
		matchIndex:  make(map[string]int),
		nextIndex:   make(map[string]int),
	}
}

func (s *State) GetTerm() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.term
}

func (s *State) SetTerm(term int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.term = term
}

func (s *State) GetVotedFor() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.votedFor
}

func (s *State) SetVotedFor(candidate string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.votedFor = candidate
}

func (s *State) GetLeader() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.leader
}

func (s *State) SetLeader(leader string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.leader = leader
}

func (s *State) GetCommitIndex() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.commitIndex
}

func (s *State) SetCommitIndex(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.commitIndex = index
}
