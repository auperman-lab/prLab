package logic

import (
	"net"
	"os"
	"sync"
)

type Room struct {
	name    string
	members map[net.Addr]*Client
	m       sync.RWMutex
	file    *os.File
}

func NewRoom(name string) (error, *Room) {
	r := &Room{}
	r.name = name
	r.members = make(map[net.Addr]*Client)
	err, f := CreateFile(name)
	if err != nil {
		return err, nil
	}
	r.file = f
	return nil, r

}

func (r *Room) Broadcast(sender *Client, msg string) {
	r.m.Lock()
	WriteMessage(r.file, msg)
	r.m.Unlock()
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}

func (r *Room) ReadChatHistory() string {
	r.m.RLock()
	chatHistory := ReadChat(r.file)
	r.m.RUnlock()

	return chatHistory

}
