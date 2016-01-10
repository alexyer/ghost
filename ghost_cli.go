package main

import (
	"flag"
	"log"

	"github.com/alexyer/ghost/cli"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host to connect")
	flag.IntVar(&port, "port", 6869, "port to connect")
}

func main() {
	flag.Parse()

	c, err := cli.ObtainClient(host, port)
	if err != nil {
		log.Printf("Error: %s. Exiting.", err.Error())
		return
	}

	cli.StartCliSession(c)
}
