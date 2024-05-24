package main

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	username string
}

func (c *client) readInput(s *server) {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		s.handleMsg(c, msg)
	}
}
