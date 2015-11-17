package client

import (
	"fmt"
	"log"
)

type GhostClient struct {
	connPool pool
	opt      *Options
}

func New(opt *Options) *GhostClient {
	return &GhostClient{
		connPool: newConnPool(opt),
		opt:      opt,
	}
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

func (c *GhostClient) process() {
	fmt.Println("process")
}

// Close the client, releasing any open resources.
// It is rare to Close a Client, as the Clientt is meant to be
// long-lived and shared between many goroutines.
func (c *GhostClient) Close() error {
	return c.connPool.Close()
}
