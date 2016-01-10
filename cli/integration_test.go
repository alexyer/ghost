package cli

import (
	"testing"
	"time"

	"github.com/alexyer/ghost/client"
	"github.com/alexyer/ghost/server"
)

var c *client.GhostClient

func init() {
	go server.GhostRun(&server.Options{Addr: "localhost:6868"})

	time.Sleep(1 * time.Second)
	c, _ = ObtainClient("localhost", 6868)
}

// this test sequense shows steps of cli-package works
// here every part of package is used.

func TestPingParsingAndProcessing(t *testing.T) {
	commString := "ping"

	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	result, err := makeRequest(c, comm, args)
	if err != nil {
		t.Error("Error on PING command: ", err.Error())
	}

	if result != "Pong!" {
		t.Error("Wrong PING response: ", result)
	}
}

func TestSetGetDelParsingAndProcessing(t *testing.T) {
	commString := "set world earth"
	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	result, err := makeRequest(c, comm, args)
	if err != nil {
		t.Error("Error on value inserting: ", err.Error())
	}

	commString = "get world"
	comm, args, err = parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	result, err = makeRequest(c, comm, args)
	if err != nil {
		t.Error("Error on value retrieving: ", err.Error())
	}

	if result != "earth" {
		t.Error("Wrong value returned: ", result)
	}

	commString = "del world"
	comm, args, err = parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	_, err = makeRequest(c, comm, args)
	if err != nil {
		t.Error("Error on value deleting: ", err.Error())
	}
}

func TestSetGetCollectionParsingAndProcessing(t *testing.T) {
	commString := "cadd predators"
	comm, args, err := parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	_, err = makeRequest(c, comm, args)
	if err != nil {
		t.Error("Error on collection creation: ", err.Error())
	}

	commString = "cget predators"
	comm, args, err = parseCommand(commString)
	if err != nil {
		t.Error("Error on command parsing: ", err.Error())
	}

	_, err = makeRequest(c, comm, args)
	if err != nil {
		t.Error("Error on collection selection: ", err.Error())
	}
}
