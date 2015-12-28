package client

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"gopkg.in/bsm/ratelimit.v1"
)

var (
	errClosed      = errors.New("ghost: client is closed")
	errPoolTimeout = errors.New("ghost: connection pool timeout")
)

type pool interface {
	First() *conn
	Get() (*conn, bool, error)
	Put(*conn) error
	Remove(*conn) error
	Len() int
	FreeLen() int
	Close() error
}

type connPool struct {
	dialer   func() (*conn, error)
	rl       *ratelimit.RateLimiter
	opt      *Options
	conns    *connList
	freeCons chan *conn
	_closed  int32
}

func newConnPool(opt *Options) *connPool {
	cp := &connPool{
		dialer:   newConnDialer(opt),
		rl:       ratelimit.New(2*opt.GetPoolSize(), time.Second),
		opt:      opt,
		conns:    newConnList(opt.GetPoolSize()),
		freeCons: make(chan *conn, opt.GetPoolSize()),
	}

	if cp.opt.GetIdleTimeout() > 0 {
		go cp.reaper()
	}

	return cp
}

func (cp *connPool) closed() bool {
	return atomic.LoadInt32(&cp._closed) == 1
}

func (cp *connPool) isIdle(c *conn) bool {
	return cp.opt.GetIdleTimeout() > 0 && time.Since(c.usedAt) > cp.opt.GetIdleTimeout()
}

func (cp *connPool) reaper() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for _ = range ticker.C {
		if cp.closed() {
			break
		}
	}

	if c := cp.First(); c != nil {
		cp.Put(c)
	}
}

// Wait for free non-idle connection.
// Return nil on timeout.
func (cp *connPool) wait() *conn {
	deadline := time.After(cp.opt.GetPoolTimeout())

	for {
		select {
		case c := <-cp.freeCons:
			if cp.isIdle(c) {
				cp.Remove(c)
				continue
			}
			return c
		case <-deadline:
			return nil
		}
	}
}

//Establish a new connection.
func (cp *connPool) new() (*conn, error) {
	if cp.rl.Limit() {
		return nil, fmt.Errorf("ghost: you open connections too fast")
	}

	c, err := cp.dialer()

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Return first non-idle connection from the pool or nil
// if there are no connections.
func (cp *connPool) First() *conn {
	for {
		select {
		case c := <-cp.freeCons:
			if cp.isIdle(c) {
				cp.conns.Remove(c)
				continue
			}
			return c
		default:
			return nil
		}
	}
}

// Return existed connection from the pool or create a new one.
func (cp *connPool) Get() (c *conn, isNew bool, err error) {
	if cp.closed() {
		err = errClosed
		return
	}

	// Fetch first non-idle connection, if available.
	if c = cp.First(); c != nil {
		return
	}

	// Try to create a new one.
	if cp.conns.Reserve() {
		c, err = cp.new()
		if err != nil {
			cp.conns.Remove(nil)
			return
		}

		cp.conns.Add(c)
		isNew = true
		return
	}

	// Wait for available connection.
	if c = cp.wait(); c != nil {
		return
	}

	err = errPoolTimeout
	return
}

func (cp *connPool) Put(c *conn) error {
	if c.rd.Buffered() != 0 {
		b, _ := c.rd.Peek(c.rd.Buffered())
		log.Printf("ghost: connection has unread data: %q", b)
		return cp.Remove(c)
	}

	if cp.opt.GetIdleTimeout() > 0 {
		c.usedAt = time.Now()
	}

	cp.freeCons <- c

	return nil
}

func (cp *connPool) Remove(c *conn) error {
	newc, err := cp.new()

	if err != nil {
		log.Printf("ghost: new failed: %s", err)
		return cp.conns.Remove(c)
	}

	err = cp.conns.Replace(c, newc)
	cp.freeCons <- newc

	return err
}

// Return total number of connections.
func (cp *connPool) Len() int {
	return cp.conns.Len()
}

// Return number of free connections.
func (cp *connPool) FreeLen() int {
	return len(cp.freeCons)
}

func (cp *connPool) Close() (retErr error) {
	if !atomic.CompareAndSwapInt32(&cp._closed, 0, 1) {
		return errClosed
	}

	// Wait for app to free connections, but don't close them immediately.
	for i := 0; i < cp.Len(); i++ {
		if c := cp.wait(); c != nil {
			break
		}
	}

	// Close all connections.
	if err := cp.conns.Close(); err != nil {
		retErr = err
	}

	return retErr
}
