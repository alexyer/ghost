package server

import "net"

type client struct {
	Conn  net.Conn
	Wchan chan []byte
	Rchan chan []byte
}

func newClient(conn net.Conn) *client {
	return &client{
		Conn:  conn,
		Wchan: make(chan []byte),
		Rchan: make(chan []byte),
	}
}
