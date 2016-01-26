package server

import (
	"log"
	"os"
)

var ghostLogger *log.Logger

func init() {
	file, err := os.Create("/tmp/ghost.log")
	if err != nil {
		log.Fatalf("ghost: cannot open log file: %s\n", err)
	}

	ghostLogger = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
