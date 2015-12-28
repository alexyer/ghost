package client

import (
	"errors"
	"fmt"
	"log"

	"github.com/alexyer/ghost/ghost"
	"github.com/alexyer/ghost/protocol"
	"github.com/golang/protobuf/proto"
)

type GhostClient struct {
	connPool pool
	opt      *Options
	processor
	msgHeader []byte
	msgBuffer []byte
}

func New(opt *Options) *GhostClient {
	newClient := &GhostClient{
		connPool:  newConnPool(opt),
		opt:       opt,
		msgHeader: make([]byte, opt.GetMsgHeaderSize()),
		msgBuffer: make([]byte, opt.GetMsgBufferSize()),
	}

	newClient.processor.process = newClient.process

	return newClient
}

func (c *GhostClient) String() string {
	return fmt.Sprintf("Ghost<%s collection: %s>", c.opt.GetAddr(), c.opt.GetCollectionName())
}

func (c *GhostClient) conn() (*conn, bool, error) {
	return c.connPool.Get()
}

func (c *GhostClient) putConn(cn *conn, ei error) {
	var err error

	if isBadConn(cn, ei) {
		err = c.connPool.Remove(cn)
	} else {
		err = c.connPool.Put(cn)
	}

	if err != nil {
		log.Printf("ghost: putConn failed: %s", err)
	}
}

// TODO(alexyer): Implement proper error handling and result return.
func (c *GhostClient) process(cmd *protocol.Command) (*protocol.Reply, error) {
	for i := 0; i <= c.opt.GetMaxRetries(); i++ {
		cn, _, err := c.conn()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		marshaledCmd, err := proto.Marshal(cmd)
		if err != nil {
			fmt.Println(err)
			c.putConn(cn, err)
			return nil, err
		}

		msgSize := ghost.IntToByteArray(int64(len(marshaledCmd)))

		if _, err := cn.Write(append(msgSize, marshaledCmd...)); err != nil {
			fmt.Println(err)
			c.putConn(cn, err)
			return nil, err
		}

		reply, err := c.getReply(cn)

		c.putConn(cn, err)
		return reply, err
	}

	return nil, errors.New("ghost: exceeded maximum number of retries")
}

// Get reply from the Ghost server and unmarshal.
func (c *GhostClient) getReply(cn *conn) (*protocol.Reply, error) {
	if _, err := cn.Read(c.msgHeader); err != nil {
		c.putConn(cn, err)
		return nil, err
	}

	if _, err := cn.Read(c.msgBuffer); err != nil {
		c.putConn(cn, err)
		return nil, err
	}

	cmdLen, _ := ghost.ByteArrayToUint64(c.msgHeader)
	reply := new(protocol.Reply)

	if err := proto.Unmarshal(c.msgBuffer[:cmdLen], reply); err != nil {
		return nil, err
	}

	return reply, nil
}

// Close the client, releasing any open resources.
// It is rare to Close a Client, as the Client is meant to be
// long-lived and shared between many goroutines.
func (c *GhostClient) Close() error {
	return c.connPool.Close()
}
