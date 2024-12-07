package logic

import (
	"fmt"
	"log"
	"log/slog"
	"strings"
)

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
	CMD_READ
)

type command struct {
	id     commandID
	client *Client
	args   []string
}

func (s *Server) nick(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("nick is required. usage: /nick NAME")
		return
	}

	c.nick = args[1]
	c.msg(fmt.Sprintf("i guess u are %s", c.nick))
}

func (s *Server) join(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("room name is required. usage: /join ROOM_NAME")
		return
	}

	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		err, room := NewRoom(roomName)
		if err != nil {
			c.msg("failed to create a room")
			return
		}
		s.rooms[roomName] = room
		r = room
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r

	r.Broadcast(c, fmt.Sprintf("%s joined the room", c.nick))

	c.msg(fmt.Sprintf("welcome to %s", roomName))

}

func (s *Server) listRooms(c *Client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("available rooms: \n%s", strings.Join(rooms, ", \n")))
}

func (s *Server) msg(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("message is required, usage: /msg MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.room.Broadcast(c, c.nick+": "+msg)
}

func (s *Server) quit(c *Client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())

	if c.room != nil {
		c.msg(fmt.Sprintf("room %s will remember you", c.room.name))

		s.quitCurrentRoom(c)
		return
	}
	c.msg("sad to see you quit :^(")
	c.conn.Close()
}

func (s *Server) quitCurrentRoom(c *Client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.Broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
	c.room = nil
}

func (s *Server) readRoom(c *Client) {
	if c.room == nil {
		slog.Info("No room to read history")
		c.msg("First join a room")
		return
	}

	chatHistory := ReadChat(c.room.file)
	c.msg(chatHistory)
}
