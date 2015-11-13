package client

import (
	"fmt"
	"net"
)

type connection struct {
	Conn       net.Conn
	Host       string
	Port       int
	Collection string
}

type GhostConnectionConfig struct {
	Host       string
	Port       int
	Collection string
}

// Try to connect to Ghost server.
func GhostConnect(config *GhostConnectionConfig) (connection, error) {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 6869
	}
	if config.Collection == "" {
		config.Collection = "main"
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	return connection{conn, config.Host, config.Port, config.Collection}, err
}

// Close server connection.
func (c *connection) Close() {
	c.Conn.Close()
}
