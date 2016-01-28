package server

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/alexyer/ghost/ghost"
	"github.com/alexyer/ghost/protocol"
	"github.com/alexyer/ghost/util"
	"github.com/golang/protobuf/proto"
)

type client struct {
	Conn       net.Conn
	Server     *Server
	MsgHeader  []byte
	collection *ghost.Collection
}

func newClient(conn net.Conn, s *Server) *client {
	return &client{
		Conn:       conn,
		Server:     s,
		MsgHeader:  make([]byte, MSG_HEADER_SIZE),
		collection: s.storage.GetCollection("main"),
	}
}

func (c *client) String() string {
	return fmt.Sprintf("Client<%s>", c.Conn.LocalAddr())
}

func (c *client) Exec() (reply []byte, err error) {
	var (
		cmd = new(protocol.Command)
	)

	// Read header
	if read, err := util.ReadData(c.Conn, c.MsgHeader, MSG_HEADER_SIZE); err != nil {
		if err != io.EOF {
			return nil, util.GhostErrorf("error when trying to read header. actually read: %d. underlying error: %s", read, err)
		} else {
			return nil, err
		}
	}

	cmdLen, _ := ghost.ByteArrayToUint64(c.MsgHeader)
	iCmdLen := int(cmdLen)
	msgBuf := c.Server.bufpool.Get(iCmdLen)

	// Read command to client buffer
	cmdRead, cmdReadErr := util.ReadData(c.Conn, msgBuf, iCmdLen)
	if cmdReadErr != nil {
		if cmdReadErr != io.EOF {
			return nil, util.GhostErrorf("Failure to read from connection. was told to read %d, actually read: %d. underlying error: %s",
				int(iCmdLen), cmdRead, cmdReadErr)
		} else {
			return nil, cmdReadErr
		}
	}

	if cmdRead > 0 && cmdReadErr == nil {
		if err := proto.Unmarshal(msgBuf[:iCmdLen], cmd); err != nil {
			c.Server.bufpool.Put(msgBuf)
			return nil, cmdReadErr
		}
	} else {
		return nil, cmdReadErr
	}

	c.Server.bufpool.Put(msgBuf)

	result, err := c.execCmd(cmd)
	return c.encodeReply(result, err)
}

func (c *client) handleCommand() {
	for {
		res, err := c.Exec()

		if err != nil {
			if err != io.EOF {
				log.Print(err)
				c.Server.logger.Print(err)
			}
			c.Conn.Close()
			return
		}

		replySize := ghost.UintToByteArray(uint64(len(res)))

		if _, err := c.Conn.Write(append(replySize, res...)); err != nil {
			log.Print(err)
			c.Server.logger.Print(err)
			c.Conn.Close()
			return
		}
	}
}

func (c *client) execCmd(cmd *protocol.Command) (result []string, err error) {
	switch *cmd.CommandId {
	case protocol.CommandId_PING:
		result, err = c.Ping()
	case protocol.CommandId_SET:
		result, err = c.Set(cmd)
	case protocol.CommandId_GET:
		result, err = c.Get(cmd)
	case protocol.CommandId_DEL:
		result, err = c.Del(cmd)
	case protocol.CommandId_CGET:
		result, err = c.CGet(cmd)
	case protocol.CommandId_CADD:
		result, err = c.CAdd(cmd)
	default:
		err = errors.New("ghost: unknown command")
	}

	return result, err
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
