package logic

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"strings"
)

type Client struct {
	conn     net.Conn
	nick     string
	room     *Room
	commands chan<- command
}

func (c *Client) ReadInput() {
	c.msg("\n\n" +
		"/nick [name] - set up your nickname\n" +
		"/join [room_name] - join a room\n" +
		"/rooms - list all rooms\n" +
		"/msg [message] - write a message\n" +
		"/q - quit the server" +
		"\n\n")

	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			slog.Error("error reading data", "error", err)
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		slog.Info("user message \n", "message", args)

		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/q":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}

	}
}

func (c *Client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *Client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
