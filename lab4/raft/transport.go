package raft

import (
	"fmt"
	"log"
	"net"
)

// Transport handles UDP communication for a Raft node.
type Transport struct {
	conn *net.UDPConn
}

// NewTransport initializes a new Transport.
func NewTransport(port int) *Transport {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP: %v", err)
	}
	return &Transport{conn: conn}
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

// Receive messages from other nodes.
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

func (t *Transport) Close() {
	if t.conn != nil {
		t.conn.Close()
	}
}
