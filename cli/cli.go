package main

import (
	"flag"
	"log"
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

	c, err := obtainClient(host, port)
	if err != nil {
		log.Printf("Error: %s. Exiting.", err.Error())
		return
	}

	startCliSession(c)
}
