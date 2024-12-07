package raft

import "encoding/json"

type MessageType string

const (
	RequestVote   MessageType = "RequestVote"
	VoteGranted   MessageType = "VoteGranted"
	AppendEntries MessageType = "AppendEntries"
)

type Message struct {
	Type      MessageType `json:"type"`
	Term      int         `json:"term"`
	Candidate string      `json:"candidate"`
}

func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

func Deserialize(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}
