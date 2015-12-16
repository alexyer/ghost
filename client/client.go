package client

import (
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
}

func New(opt *Options) *GhostClient {
	newClient := &GhostClient{
		connPool: newConnPool(opt),
		opt:      opt,
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
func (c *GhostClient) process(cmd *protocol.Command) {
	for i := 0; i <= c.opt.GetMaxRetries(); i++ {
		cn, _, err := c.conn()
		if err != nil {
			fmt.Println(err)
			return
		}

		marshaledCmd, err := proto.Marshal(cmd)
		if err != nil {
			fmt.Println(err)
			c.putConn(cn, err)
			return
		}

		msgSize := ghost.IntToByteArray(int64(len(marshaledCmd)))

		if _, err := cn.Write(append(msgSize, marshaledCmd...)); err != nil {
			fmt.Println(err)
			c.putConn(cn, err)
			return
		}

		resp := make([]byte, 4096)
		if _, err := cn.Read(resp); err != nil {
			fmt.Println(err)
			c.putConn(cn, err)
			return
		}

		fmt.Println(string(resp))

		c.putConn(cn, err)
		return
	}
}

// Close the client, releasing any open resources.
// It is rare to Close a Client, as the Client is meant to be
// long-lived and shared between many goroutines.
func (c *GhostClient) Close() error {
	return c.connPool.Close()
}
