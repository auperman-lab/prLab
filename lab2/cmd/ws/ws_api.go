package ws

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/auperman-lab/lab2/cmd/ws/logic"
	"log"
	"log/slog"
	"net"
	"net/http"
)

type WsServer struct {
	addr string
}

func NewWsServer(addr string) *WsServer {
	return &WsServer{
		addr: addr,
	}
}

func (ws WsServer) Run() {
	s := logic.NewServer()
	go s.Run()

	listener, err := net.Listen("tcp", ws.addr)
	if err != nil {
		log.Fatalf("unable to start websocket server: %s", err.Error())
	}

	defer listener.Close()
	slog.Info("server started on ", "address", ws.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		// Perform WebSocket handshake
		if !performWebSocketHandshake(conn) {
			conn.Close()
			continue
		}

		c := s.NewClient(conn)
		go c.ReadInput()
	}
}

func performWebSocketHandshake(conn net.Conn) bool {
	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Printf("failed to read request: %s", err.Error())
		return false
	}

	// Check WebSocket headers
	if req.Header.Get("Upgrade") != "websocket" || req.Header.Get("Connection") != "Upgrade" {
		log.Printf("invalid WebSocket headers")
		return false
	}

	// Generate the Sec-WebSocket-Accept header
	acceptKey := computeAcceptKey(req.Header.Get("Sec-WebSocket-Key"))

	// Write the WebSocket handshake response
	_, err = fmt.Fprintf(conn, "HTTP/1.1 101 Switching Protocols\r\n"+
		"Upgrade: websocket\r\n"+
		"Connection: Upgrade\r\n"+
		"Sec-WebSocket-Accept: %s\r\n\r\n", acceptKey)
	if err != nil {
		log.Printf("failed to send handshake response: %s", err.Error())
		return false
	}

	return true
}

// computeAcceptKey generates the Sec-WebSocket-Accept value based on Sec-WebSocket-Key
func computeAcceptKey(key string) string {
	// GUID is defined by the WebSocket protocol
	GUID := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	hash := sha1.New()
	hash.Write([]byte(key + GUID))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
