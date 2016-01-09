package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/alexyer/ghost/client"
)

func obtainClient(host string, port int) (*client.GhostClient, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	return connect(addr)
}

func connect(addr string) (*client.GhostClient, error) {
	c := client.New(&client.Options{Addr: addr})

	if _, err := c.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("cli-ghost: cannot obtain connection to %s", addr))
	}

	log.Printf("Connection to %s is successfull.", addr)
	return c, nil
}
