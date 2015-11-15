package client

import (
	"bufio"
	"net"
	"time"
)

const defaultBufSize = 4096

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
