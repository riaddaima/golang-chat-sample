package main

import (
	"fmt"
	"log"
	"net"
)

type server struct {
	clients map[net.Addr]*client
}

func newServer() *server {
	return &server{
		clients: make(map[net.Addr]*client),
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("A new client has joined the chat: %s\n", conn.RemoteAddr().String())
	conn.Write([]byte("Who are you?\n"))
	newClient := &client{
		conn:     conn,
		username: "",
	}
	s.clients[conn.RemoteAddr()] = newClient
	return newClient
}

func (s *server) handleMsg(sender *client, msg string) {
	if sender.username == "" {
		sender.username = msg
		for _, c := range s.clients {
			welcomeMsg := fmt.Sprintf("Welcome, %s!\n", sender.username)
			c.conn.Write([]byte(welcomeMsg))
		}
	} else {
		for addr, c := range s.clients {
			// Prevent duplicate message for the author.
			if sender.conn.RemoteAddr() != addr {
				c.conn.Write([]byte(sender.username + ": " + msg + "\n"))
			}
		}
	}
}
