package client

import (
	"sync"
	"sync/atomic"
)

type connList struct {
	conns []*conn
	mu    sync.Mutex
	len   int32 // atomic
	size  int32
}

func newConnList(size int) *connList {
	return &connList{
		conns: make([]*conn, 0, size),
		size:  int32(size),
	}
}

func (cl *connList) Len() int {
	return int(atomic.LoadInt32(&cl.len))
}

// Reserve place in the list and return true on success.
// The caller must add or remove connection if place was reserved.
func (cl *connList) Reserver() bool {
	length := atomic.AddInt32(&cl.len, 1)
	reserved := length <= cl.size

	if !reserved {
		atomic.AddInt32(&cl.len, -1)
	}

	return reserved
}

// Add connection to the list.
// The caller must reserver place first.
func (cl *connList) Add(cn *conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	cl.conns = append(cl.conns, cn)
}

func (cl *connList) Remove(cn *conn) error {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if cn == nil {
		atomic.AddInt32(&cl.len, -1)
		return nil
	}

	for i, c := range cl.conns {
		if c == cn {
			cl.conns = append(cl.conns[:i], cl.conns[i+1:]...)
			atomic.AddInt32(&cl.len, -1)
			return cn.Close()
		}
	}

	if cl.closed() {
		return nil
	}

	panic("connection not found in the list")
}

func (cl *connList) Replace(cn, newcn *conn) error {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	for i, c := range cl.conns {
		if c == cn {
			cl.conns[i] = newcn
			return cn.Close()
		}
	}

	if cl.closed() {
		return newcn.Close()
	}

	panic("connection not found in the list")
}

func (cl *connList) Close() (retErr error) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	for _, c := range cl.conns {
		if err := c.Close(); err != nil {
			retErr = err
		}
	}

	cl.conns = nil
	atomic.StoreInt32(&cl.len, 0)

	return retErr
}

func (cl *connList) closed() bool {
	return cl.conns == nil
}
