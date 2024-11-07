package ws

import (
	"github.com/auperman-lab/lab2/cmd/ws/logic"
	"log"
	"log/slog"
	"net"
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

		c := s.NewClient(conn)
		go c.ReadInput()
	}
}
