package server

import (
	"errors"
	"fmt"
	"net"

	"github.com/alexyer/ghost/ghost"
	"github.com/alexyer/ghost/protocol"
	"github.com/golang/protobuf/proto"
)

type client struct {
	Conn   net.Conn
	Server *Server
	Header []byte
	Buffer []byte
}

func newClient(conn net.Conn, s *Server, headerSize, bufSize int) *client {
	return &client{
		Conn:   conn,
		Server: s,
		Header: make([]byte, headerSize),
		Buffer: make([]byte, bufSize),
	}
}

func (c *client) String() string {
	return fmt.Sprintf("Client<%s>", c.Conn.LocalAddr())
}

func (c *client) Exec() (result string, err error) {
	cmd := new(protocol.Command)

	cmdLen, _ := ghost.ByteArrayToUint64(c.Header)

	if err := proto.Unmarshal(c.Buffer[:cmdLen], cmd); err != nil {
		return "", err
	}

	switch *cmd.CommandId {
	case protocol.CommandId_PING:
		result, err = c.Server.Ping()
	default:
		err = errors.New("ghost: unknown command")
	}

	return
}
