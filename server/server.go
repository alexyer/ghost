package server

import (
	"io"
	"log"
	"net"
)

type Server struct {
	opt *Options
}

func GhostRun(opt *Options) Server {
	s := Server{opt: opt}

	log.Printf("Starting Ghost server on %s", s.opt.GetAddr())

	s.handle()

	return s
}

func (s *Server) handle() {
	ln, err := net.Listen("tcp", s.opt.GetAddr())

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
		client := newClient(conn, s)
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

		// Read command to client buffer
		if err := s.read(c.Conn, c.Buffer); err != nil {
			log.Print(err)
			c.Conn.Close()
			return
		}

		res, err := c.Exec()

		if err != nil {
			log.Print(err)
			c.Conn.Close()
			return
		}

		if _, err := c.Conn.Write(res); err != nil {
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
