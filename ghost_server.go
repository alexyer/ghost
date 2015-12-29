package main

import (
	"flag"
	"fmt"

	"github.com/alexyer/ghost/server"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host")
	flag.IntVar(&port, "port", 6869, "port")
}

func main() {
	flag.Parse()
	server.GhostRun(&server.Options{Addr: fmt.Sprintf("%s:%d", host, port)})
}
