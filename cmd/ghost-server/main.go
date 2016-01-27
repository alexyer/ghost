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
	logfile    string
	host       string
	port       int
	socket     string
)

func init() {
	flag.StringVar(&cpuprofile, "cpuprofile", "", "enable cpu profiling and write profile to file")
	flag.StringVar(&logfile, "logfile", "/tmp/ghost.log", "log file path")
	flag.StringVar(&host, "host", "localhost", "host")
	flag.IntVar(&port, "port", 6869, "port")
	flag.StringVar(&socket, "socket", "", "listen to unix socket")
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

	if socket != "" {
		os.Remove(socket)
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
	server.GhostRun(&server.Options{
		Addr:        fmt.Sprintf("%s:%d", host, port),
		LogfileName: logfile,
		Socket:      socket,
	})
}

func main() {
	initCPUProfile()
	go handleShutdown()
	initServer()
}
