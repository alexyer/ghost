package client

import (
	"net"
	"time"
)

type Options struct {
	// Dialer creates new network connection and has priority
	// over network and Addr options.
	Dialer func() (net.Conn, error)

	// The network type.
	// Default: tcp.
	Network string

	// host:port address.
	Addr string

	// Collection name.
	// Default: "main".
	CollectionName string

	// The maximum number of socket connections.
	// Default: 10 connections.
	PoolSize int

	// Specifies amount of time client watis for connection if all
	// connections are busy before returning an error.
	// Default: 5 seconds.
	PoolTimeout time.Duration

	// Specifies amount of time after which client closes idle connections.
	// Default: not close idle connections.
	IdleTimeout time.Duration

	// Specifies the deadline for establishing new connections.
	// If reached, dial will fail with a timeout.
	DialTiemout time.Duration

	// The maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int

	// Size of the message header
	MsgHeaderSize int

	// Size of the message buffer
	MsgBufferSize int
}

func (opt *Options) GetDialer() func() (net.Conn, error) {
	if opt.Dialer == nil {
		opt.Dialer = func() (net.Conn, error) {
			return net.DialTimeout(opt.GetNetwork(), opt.GetAddr(), opt.GetDialTimeout())
		}
	}

	return opt.Dialer
}

func (opt *Options) GetPoolSize() int {
	if opt.PoolSize == 0 {
		return 10
	}
	return opt.PoolSize
}

func (opt *Options) GetNetwork() string {
	if opt.Network == "" {
		return "tcp"
	}

	return opt.Network
}

func (opt *Options) GetAddr() string {
	return opt.Addr
}

func (opt *Options) GetDialTimeout() time.Duration {
	if opt.DialTiemout == 0 {
		return 5 * time.Second
	}

	return opt.DialTiemout
}

func (opt *Options) GetCollectionName() string {
	if opt.CollectionName == "" {
		return "main"
	}

	return opt.CollectionName
}

func (opt *Options) GetPoolTimeout() time.Duration {
	if opt.PoolTimeout == 0 {
		return 1 * time.Second
	}

	return opt.PoolTimeout
}

func (opt *Options) GetIdleTimeout() time.Duration {
	return opt.IdleTimeout
}

func (opt *Options) GetMaxRetries() int {
	return opt.MaxRetries
}

func (opt *Options) GetMsgHeaderSize() int {
	if opt.MsgHeaderSize == 0 {
		opt.MsgHeaderSize = 8
	}

	return opt.MsgHeaderSize
}

func (opt *Options) GetMsgBufferSize() int {
	if opt.MsgBufferSize == 0 {
		opt.MsgBufferSize = 1024 * 1024
	}

	return opt.MsgBufferSize
}
