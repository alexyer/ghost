package server

import (
	"log"
	"os"

	"github.com/alexyer/ghost/util"
)

// Create file logger.
func getLogger(filename string) *log.Logger {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(util.GhostErrorf("cannot open log file: %s\n", err))
	}

	return log.New(file, "", log.LstdFlags|log.Lshortfile)
}
