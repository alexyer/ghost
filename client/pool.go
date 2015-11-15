package client

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
	conns    *connList
	freeCons chan *conn
}
