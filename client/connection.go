package client

import (
	"fmt"
	"net"
)

type connection struct {
	Conn net.Conn
}

// Try to connect to Ghost server.
func ConnectGhost(host string, port int) (connection, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	return connection{conn}, err
}

// Close server connection.
func (c *connection) Close() {
	c.Conn.Close()
}
