package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s\n", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on port: 8888\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s\n", err.Error())
			continue
		}

		c := s.newClient(conn)
		go c.readInput(s)
	}
}
