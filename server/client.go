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
	Conn       net.Conn
	Server     *Server
	MsgHeader  []byte
	MsgBuffer  []byte
	collection *ghost.Collection
}

func newClient(conn net.Conn, s *Server) *client {
	return &client{
		Conn:       conn,
		Server:     s,
		MsgHeader:  make([]byte, s.opt.GetMsgHeaderSize()),
		MsgBuffer:  make([]byte, s.opt.GetMsgBufferSize()),
		collection: s.storage.GetCollection("main"),
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

	cmdLen, _ := ghost.ByteArrayToUint64(c.MsgHeader)

	if err := proto.Unmarshal(c.MsgBuffer[:cmdLen], cmd); err != nil {
		return nil, err
	}

	switch *cmd.CommandId {
	case protocol.CommandId_PING:
		result, err = c.Ping()
	case protocol.CommandId_SET:
		result, err = c.Set(cmd)
	case protocol.CommandId_GET:
		result, err = c.Get(cmd)
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
