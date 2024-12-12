package raft

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

type Transport struct {
	mu         sync.Mutex
	conn       *net.UDPConn
	managerURL string
	nodePort   int // Add nodeAddr to store the node's address

}

func NewTransport(port int, managerURL string) *Transport {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP: %v", err)
	}

	return &Transport{conn: conn, nodePort: port, managerURL: managerURL}
}

func (t *Transport) Send(addr string, msg *Message) {
	data, err := msg.Serialize()
	if err != nil {
		log.Printf("Error serializing message: %v", err)
		return
	}
	destAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Printf("Error resolving address %s: %v", addr, err)
		return
	}
	t.conn.WriteToUDP(data, destAddr)
}

func (t *Transport) Receive(handler func(msg *Message, addr *net.UDPAddr)) {
	buf := make([]byte, 1024)
	for {
		n, addr, err := t.conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error reading UDP message: %v", err)
			continue
		}
		msg, err := Deserialize(buf[:n])
		if err != nil {
			log.Printf("Error deserializing message: %v", err)
			continue
		}
		handler(msg, addr)
	}
}

func (t *Transport) Broadcast(msg *Message, peers []string) {
	for _, peer := range peers {
		go t.Send(peer, msg)
	}
}

// TODO failing to send the new leader
func (t *Transport) SendLeaderAddr(nodeId string) {
	reqBody := struct {
		NewLeaderURL string `json:"new_leader_url"`
	}{
		//NewLeaderURL: "http://"+nodeId+":"+t.nodePort,
		NewLeaderURL: fmt.Sprintf("http://%s:%d", nodeId, t.nodePort),
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("Error marshaling request: %s", err)
		return
	}

	log.Printf("Request body: %s", data)
	log.Printf("Manager URL: %s", t.managerURL)

	resp, err := http.Post(t.managerURL+"/changeLeader", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Failed to send Leader Address: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Manager responded with status: %d", resp.StatusCode)
	} else {
		log.Printf("Successfully sent leader address to manager")
	}
	return
}
