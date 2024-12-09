package main

import (
	"github.com/auperman-lab/lab3/raft"
)

func main() {
	peers1 := []string{"localhost:9002", "localhost:9003", "localhost:9004", "localhost:9005"}
	peers2 := []string{"localhost:9001", "localhost:9003", "localhost:9004", "localhost:9005"}
	peers3 := []string{"localhost:9001", "localhost:9002", "localhost:9004", "localhost:9005"}
	peers4 := []string{"localhost:9001", "localhost:9002", "localhost:9003", "localhost:9005"}
	peers5 := []string{"localhost:9001", "localhost:9002", "localhost:9003", "localhost:9004"}

	//node1 := raft.NewNode("node1", 9001, peers1)
	//node2 := raft.NewNode("node2", 9002, peers2)
	//node3 := raft.NewNode("node3", 9003, peers3)
	//node4 := raft.NewNode("node4", 9004, peers4)
	//node5 := raft.NewNode("node5", 9005, peers5)

	go raft.NewNode("node1", 9001, peers1)
	go raft.NewNode("node2", 9002, peers2)
	go raft.NewNode("node3", 9003, peers3)
	go raft.NewNode("node4", 9004, peers4)
	go raft.NewNode("node5", 9005, peers5)

	//nodes := []*raft.Node{node1, node2, node3, node4, node5}
	//
	//nodes = nodes[1:]

	select {}
}
