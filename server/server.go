package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
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

	log.Printf("Starting Ghost server on %s:%d", s.Host, s.Port)

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
	go func() {
		client := newClient(conn)

		// TODO(alexyer): Debug
		fmt.Printf("New client: %v\n", client)

		go s.handleCommand(client)
	}()
}

func (s *Server) handleCommand(c *client) {
	buf := make([]byte, 8)

	for {
		if _, err := c.Conn.Read(buf); err != nil {
			c.Conn.Close()
			return
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(r.Intn(100)) * time.Millisecond)

		if _, err := c.Conn.Write([]byte("response")); err != nil {
			c.Conn.Close()
			return
		}
	}
}
