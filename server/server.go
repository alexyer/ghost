package server

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	Host             string
	Port             int
	ClientHeaderSize int
	ClientBufSize    int
}

type GhostServerConfig struct {
	Host             string
	Port             int
	ClientHeaderSize int
	ClientBufSize    int
}

func GhostRun(config *GhostServerConfig) Server {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 6869
	}

	if config.ClientBufSize == 0 {
		config.ClientBufSize = 4096
	}

	if config.ClientHeaderSize == 0 {
		config.ClientHeaderSize = 8
	}

	s := Server{config.Host, config.Port, config.ClientHeaderSize, config.ClientBufSize}

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
		client := newClient(conn, s, s.ClientHeaderSize, s.ClientBufSize)

		// TODO(alexyer): Debug
		fmt.Printf("New client: %v\n", client)

		go s.handleCommand(client)
	}()
}

func (s *Server) handleCommand(c *client) {
	for {
		if err := s.read(c.Conn, c.Header); err != nil {
			log.Print(err)
			c.Conn.Close()
			return
		}

		if err := s.read(c.Conn, c.Buffer); err != nil {
			log.Print(err)
			c.Conn.Close()
			return
		}

		res, _ := c.Exec()

		if _, err := c.Conn.Write([]byte(res)); err != nil {
			c.Conn.Close()
			return
		}
	}
}

func (s *Server) read(conn net.Conn, buf []byte) error {
	// TODO(alexyer): Implement proper error handling
	if _, err := conn.Read(buf); err != nil {
		if err != io.EOF {
			return err
		}
	}

	return nil
}
