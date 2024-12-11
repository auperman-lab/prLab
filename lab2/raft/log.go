package raft

import (
	"fmt"
	"sync"
)

type Log struct {
	mu   sync.Mutex
	logs []LogEntry
}

type LogEntry struct {
	Index int
	Term  int
	Data  []byte
}

func NewLog() *Log {
	return &Log{
		logs: make([]LogEntry, 0),
	}
}

func (l *Log) Append(entry LogEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Printf("\033[31m  appended log : %+v \033[0m\n", entry)

	l.logs = append(l.logs, entry)
}

func (l *Log) Get(index int) (LogEntry, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if index < 0 || index >= len(l.logs) {
		return LogEntry{}, false
	}
	return l.logs[index], true
}

func (l *Log) LastIndex() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.logs) == 0 {
		return -1
	}
	return l.logs[len(l.logs)-1].Index
}

func (l *Log) LastTerm() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.logs) == 0 {
		return 0
	}
	return l.logs[len(l.logs)-1].Term
}

func (l *Log) CommittedLogs(commitIndex int) []LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()

	if commitIndex >= len(l.logs) {
		return l.logs
	}

	return l.logs[:commitIndex+1]
}

func (l *Log) UncommittedLogs(commitIndex int) []LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()

	if commitIndex >= len(l.logs)-1 {
		return []LogEntry{}
	}

	return l.logs[commitIndex+1:]
}
