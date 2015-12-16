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

func newClient(conn net.Conn, s *Server) *client {
	return &client{
		Conn:   conn,
		Server: s,
		Header: make([]byte, s.opt.GetClientHeaderSize()),
		Buffer: make([]byte, s.opt.GetClientBufSize()),
	}
}

func (c *client) String() string {
	return fmt.Sprintf("Client<%s>", c.Conn.LocalAddr())
}

func (c *client) Exec() (reply []byte, err error) {
	var (
		result []string
		cmd    = new(protocol.Command)
	)

	cmdLen, _ := ghost.ByteArrayToUint64(c.Header)

	if err := proto.Unmarshal(c.Buffer[:cmdLen], cmd); err != nil {
		return nil, err
	}

	switch *cmd.CommandId {
	case protocol.CommandId_PING:
		result, err = c.Server.Ping()
	default:
		err = errors.New("ghost: unknown command")
	}

	return c.encodeReply(result, err)
}

func (c *client) encodeReply(values []string, err error) ([]byte, error) {
	var errMsg string

	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = ""
	}

	return proto.Marshal(&protocol.Reply{
		Values: values,
		Error:  &errMsg,
	})
}
