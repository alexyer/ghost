package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Host string
	Port int
}

func RunGhost(host string, port int) Server {
	s := Server{host, port}

	log.Printf("Starting ghost server on %s:%d", s.Host, s.Port)

	s.handle()

	return s
}

func (s *Server) handle() {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	fmt.Println(conn)
}
