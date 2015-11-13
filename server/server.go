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

type GhostServerConfig struct {
	Host string
	Port int
}

func GhostRun(config *GhostServerConfig) Server {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 6869
	}

	s := Server{config.Host, config.Port}

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
