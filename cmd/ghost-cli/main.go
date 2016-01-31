package main

import (
	"flag"
	"log"

	"github.com/alexyer/ghost/cli"
	"github.com/alexyer/ghost/client"
)

var (
	host   string
	port   int
	socket string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host to connect")
	flag.IntVar(&port, "port", 6869, "port to connect")
	flag.StringVar(&socket, "socket", "", "listen to unix socket")
}

func main() {
	var (
		c   *client.GhostClient
		err error
	)

	flag.Parse()

	if socket == "" {
		c, err = cli.ObtainClient(host, port)
	} else {
		c, err = cli.ObtainUnixSocketClient(socket)
	}

	if err != nil {
		log.Printf("Error: %s. Exiting.", err.Error())
		return
	}

	cli.StartCliSession(c)
}
