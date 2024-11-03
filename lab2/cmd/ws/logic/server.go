package logic

import (
	"log"
	"log/slog"
	"net"
)

type Server struct {
	rooms    map[string]*Room
	commands chan command
}

func NewServer() *Server {
	return &Server{
		rooms:    make(map[string]*Room),
		commands: make(chan command),
	}
}

func (s *Server) Run() {
	for cmd := range s.commands {
		slog.Info("user message %s\n", cmd)

		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *Server) NewClient(conn net.Conn) *Client {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())

	return &Client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}
}
