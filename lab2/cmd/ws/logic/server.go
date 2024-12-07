package logic

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	rooms    map[string]*Room
	commands chan command
}

func NewServer() *Server {
	roomsFiles := ReadRooms()
	rooms := make(map[string]*Room, len(roomsFiles))
	for _, file := range roomsFiles {
		roomName := strings.TrimSuffix(file, ".json")
		r := &Room{
			name:    roomName,
			members: make(map[net.Addr]*Client),
			file:    GetFile(file),
		}
		rooms[roomName] = r
	}

	fmt.Println(rooms)

	return &Server{
		rooms:    rooms,
		commands: make(chan command),
	}
}

func (s *Server) Run() {
	for cmd := range s.commands {
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
		case CMD_READ:
			s.readRoom(cmd.client)
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
