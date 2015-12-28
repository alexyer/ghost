package client

import (
	"testing"
	"time"

	"github.com/alexyer/ghost/server"
)

var (
	c *GhostClient
)

func init() {
	go server.GhostRun(&server.Options{Addr: "localhost:6869"})

	// Wait until server start
	// Yep, ugly. Tell if you know better solution
	time.Sleep(1 * time.Second)
	c = New(&Options{Addr: "localhost:6869"})
}

func TestPing(t *testing.T) {
	reply, _ := c.Ping()

	if reply.Values[0] != "Pong!" {
		t.Error("PING: error")
	}
}

func TestCollection(t *testing.T) {
	_, err := c.CGet("test_collection")

	if err == nil {
		t.Error("CGET: error")
	}

	_, err = c.CAdd("test_collection")

	if err != nil {
		t.Error(err)
	}

	_, err = c.CAdd("test_collection")

	if err == nil {
		t.Error("CGET: error")
	}

	_, err = c.CGet("test_collection")

	if err != nil {
		t.Error(err)
	}
}

func TestBasicOperations(t *testing.T) {
	_, err := c.Get("test_key")

	if err == nil {
		t.Error("GET: error")
	}

	c.Set("test_key", "test_val")

	res, _ := c.Get("test_key")

	if res != "test_val" {
		t.Errorf("wrong value set/get")
	}

	c.Del("test_key")

	res, _ = c.Get("test_key")

	if res == "test_val" {
		t.Errorf("wrong del")
	}
}
