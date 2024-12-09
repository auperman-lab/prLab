package raft

import (
	"encoding/json"
)

type MessageType string

const (
	VoteReq           MessageType = "VoteRequest"
	VoteResp          MessageType = "VoteGranted"
	AppendEntriesReq  MessageType = "AppendEntriesRequest"
	AppendEntriesResp MessageType = "AppendEntriesResponse"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}

type AppendEntriesRequest struct {
	Term         int        `json:"term"`           // Current term of the leader
	LeaderID     string     `json:"leader_id"`      // ID of the leader sending the request
	PrevLogIndex int        `json:"prev_log_index"` // Index of the log entry immediately preceding new entries
	PrevLogTerm  int        `json:"prev_log_term"`  // Term of the entry at PrevLogIndex
	Entries      []LogEntry `json:"entries"`        // New log entries to be appended
	LeaderCommit int        `json:"leader_commit"`  // Leader's commit index
}
type AppendEntriesResponse struct {
	Term    int  `json:"term"`
	Success bool `json:"success"`
}

type VoteRequest struct {
	Term         int    `json:"term"`
	CandidateID  string `json:"candidate_id"`
	LastLogIndex int    `json:"last_log_index"`
	LastLogTerm  int    `json:"last_log_term"`
}
type VoteResponse struct {
	Term        int  `json:"term"`         // Current term of the responding node
	VoteGranted bool `json:"vote_granted"` // Whether the vote was granted
}

func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

func Deserialize(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)

	return &msg, err
}

func DeserializeToVoteRequest(data interface{}) (*VoteRequest, error) {
	var voteReq VoteRequest
	payloadData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(payloadData, &voteReq)

	return &voteReq, err
}

func DeserializeToVoteResponse(data interface{}) (*VoteResponse, error) {
	var voteResp VoteResponse
	payloadData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(payloadData, &voteResp)

	return &voteResp, err
}

func DeserializeToAppendResponse(data interface{}) (*AppendEntriesResponse, error) {
	var appendEntriesResponse AppendEntriesResponse
	payloadData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(payloadData, &appendEntriesResponse)

	return &appendEntriesResponse, err
}

func DeserializeToAppendRequest(data interface{}) (*AppendEntriesRequest, error) {
	var appendEntriesRequest AppendEntriesRequest
	payloadData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(payloadData, &appendEntriesRequest)

	return &appendEntriesRequest, err
}
