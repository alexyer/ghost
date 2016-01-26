package server

import (
	"log"
	"net"

	"github.com/alexyer/ghost/ghost"
	"github.com/alexyer/ghost/util"
)

type Server struct {
	bufpool *util.Bufpool
	opt     *Options
	storage *ghost.Storage
	logger  *log.Logger
}

func GhostRun(opt *Options) Server {
	s := Server{
		bufpool: util.NewBufpool(),
		opt:     opt,
		storage: ghost.GetStorage(),
		logger:  getLogger(opt.GetLogfileName()),
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

		go newClient(conn, s).handleCommand()
	}
}
