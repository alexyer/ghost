package cli

import (
	"testing"
	"time"

	"github.com/alexyer/ghost/server"
)

func TestFailedConnect(t *testing.T) {
	_, err := ObtainClient("somehost", 3000)

	if err == nil {
		t.Error("connection doesn't tell about failure.")
	}
}

func TestSuccessfullConnect(t *testing.T) {
	go server.GhostRun(&server.Options{Addr: "localhost:6870"})
	time.Sleep(1 * time.Second)

	c, err := ObtainClient("localhost", 6870)
	if err != nil {
		t.Error("Can't connect to test localhost ", err.Error())
	}

	_, err = c.Ping()
	if err != nil {
		t.Error("Error on ping ", err.Error())
	}
}
