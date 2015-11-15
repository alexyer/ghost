package client

import (
	"net"
	"time"
)

type Options struct {
	// Dialer creates new network connection and has priority
	// over network and addr options.
	dialer func() (net.Conn, error)

	// The network type.
	// Default: tcp.
	network string

	// host:port address.
	addr string

	// Collection name.
	// Default: "main".
	collectionName string

	// The maximum number of socket connections.
	// Default: 10 connections.
	poolSize int

	// Specifies amount of time client watis for connection if all
	// connections are busy before returning an error.
	// Default: 5 seconds.
	poolTimeout time.Duration

	// Specifies amount of time after which client closes idle connections.
	// Default: not close idle connections.
	idleTimeout time.Duration

	// Specifies the deadline for establishing new connections.
	// If reached, dial will fail with a timeout.
	dialTimeout time.Duration
}

func (opt *Options) getDialer() func() (net.Conn, error) {
	if opt.Dialer == nil {
		opt.Dialer = func() (net.Conn, error) {
			return net.DialTimeout(opt.GetNetwork(), opt.GetAddr(), opt.GetDialTimeout())
		}
	}
}

func (opt *Options) GetPoolSize() int {
	if opt.poolSize == 0 {
		return 10
	}
	return opt.poolSize
}

func (opt *Options) GetNetwork() string {
	if opt.network == "" {
		return "tcp"
	}

	return opt.network
}

func (opt *Options) GetAddr() string {
	return opt.addr
}

func (opt *Options) GetDialTimeout() time.Duration {
	if opt.dialTimeout == 0 {
		return 5 * time.Second
	}

	return opt.dialTimeout
}

func (opt *Options) GetCollectionName() string {
	if opt.collectionName == "" {
		return "main"
	}

	return opt.collectionName
}

func (opt *Options) GetPoolTimeout() time.Duration {
	if opt.poolTimeout == 0 {
		return 1 * time.Second
	}

	return opt.poolTimeout
}

func (opt *Options) GetIdleTimeout() time.Duration {
	return opt.idleTimeout
}
