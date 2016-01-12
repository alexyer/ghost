package server

import (
	"log"
	"net"

	"github.com/alexyer/ghost/ghost"
)

type Server struct {
	bufpool *bufpool
	opt     *Options
	storage *ghost.Storage
}

func GhostRun(opt *Options) Server {
	s := Server{
		bufpool: newBufpool(),
		opt:     opt,
		storage: ghost.GetStorage(),
	}

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
		go newClient(conn, s).handleCommand()
	}()
}
