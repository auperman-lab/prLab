package main

import (
	"fmt"
	"github.com/auperman-lab/lab4/raft"
	"os"
	"strconv"
	"strings"
)

func main() {
	nodeID := os.Getenv("NODE_ID")
	portStr := os.Getenv("PORT")
	peersEnv := os.Getenv("PEERS")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Invalid port number: %s\n", portStr)
		return
	}
	// Split peers from environment variable
	peers := strings.Split(peersEnv, ",")

	go raft.NewRaftNode(nodeID, port, peers)

	fmt.Printf("Node %s started on port %s with peers: %v\n", nodeID, port, peers)

	select {}
}
