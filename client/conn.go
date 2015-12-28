package client

import (
	"bufio"
	"net"
	"time"
)

const defaultBufSize = 4096

var zeroTime = time.Time{}

type conn struct {
	netcn net.Conn
	rd    *bufio.Reader
	buf   []byte

	usedAt       time.Time
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func newConnDialer(opt *Options) func() (*conn, error) {
	dialer := opt.GetDialer()
	return func() (*conn, error) {
		netcn, err := dialer()

		if err != nil {
			return nil, err
		}

		cn := &conn{
			netcn: netcn,
			buf:   make([]byte, defaultBufSize),
		}

		cn.rd = bufio.NewReader(cn)
		return cn, nil
	}
}

func (c *conn) Read(b []byte) (int, error) {
	if c.ReadTimeout != 0 {
		c.netcn.SetReadDeadline(time.Now().Add(c.ReadTimeout))
	} else {
		c.netcn.SetReadDeadline(zeroTime)
	}

	return c.netcn.Read(b)
}

func (c *conn) Write(b []byte) (int, error) {
	if c.ReadTimeout != 0 {
		c.netcn.SetWriteDeadline(time.Now().Add(c.ReadTimeout))
	} else {
		c.netcn.SetWriteDeadline(zeroTime)
	}

	return c.netcn.Write(b)
}

func (c *conn) Close() error {
	return c.netcn.Close()
}
