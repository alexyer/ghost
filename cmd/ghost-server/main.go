package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/alexyer/ghost/server"
)

var (
	cpuprofile string
	host       string
	port       int
)

func init() {
	flag.StringVar(&cpuprofile, "cpuprofile", "", "enable cpu profiling and write profile to file")
	flag.StringVar(&host, "host", "localhost", "host")
	flag.IntVar(&port, "port", 6869, "port")
	flag.Parse()
}

func handleShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()
}

func cleanup() {
	if cpuprofile != "" {
		pprof.StopCPUProfile()
		log.Print("Stopped CPU profiling")
	}

	log.Print("Bye!")
}

func initCPUProfile() {
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Starting CPU profiling. Output: %s", cpuprofile)

		pprof.StartCPUProfile(f)
	}
}

func initServer() {
	server.GhostRun(&server.Options{Addr: fmt.Sprintf("%s:%d", host, port)})
}

func main() {
	initCPUProfile()
	go handleShutdown()
	initServer()
}