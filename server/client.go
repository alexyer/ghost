package server

import (
	"fmt"
	"net"
)

type client struct {
	Conn   net.Conn
	Header []byte
	Buffer []byte
}

func newClient(conn net.Conn, bufSize int) *client {
	return &client{
		Conn:   conn,
		Buffer: make([]byte, bufSize),
	}
}

func (c *client) String() string {
	return fmt.Sprintf("Client<%s>", c.Conn.LocalAddr())
}
