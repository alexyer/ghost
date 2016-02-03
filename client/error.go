package client

import (
	"errors"

	"github.com/alexyer/ghost/protocol"
)

func isBadConn(c *conn, ei error) bool {
	if c.rd.Buffered() > 0 {
		return true
	}

	if ei == nil {
		return false
	}

	return true
}

func getReplyErrors(reply *protocol.Reply, err error) (*protocol.Reply, error) {
	if err != nil {
		return reply, err
	}

	if *reply.Error != "" {
		return reply, errors.New(*reply.Error)
	}

	return reply, err
}
