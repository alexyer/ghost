package cli

import (
	"errors"
	"fmt"
	"log"

	"github.com/alexyer/ghost/client"
)

// Function returns ghost-client on successfull connection to server
// and error with description otherwise.
func ObtainClient(host string, port int) (*client.GhostClient, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	return connect(addr, "tcp")
}

// Obtain client instance working over unix file socket.
func ObtainUnixSocketClient(socket string) (*client.GhostClient, error) {
	return connect(socket, "unix")
}

func connect(addr, network string) (*client.GhostClient, error) {
	c := client.New(&client.Options{
		Addr:    addr,
		Network: network,
	})

	if _, err := c.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("cli-ghost: cannot obtain connection to %s", addr))
	}

	log.Printf("Connection to %s is successfull.", addr)
	return c, nil
}
