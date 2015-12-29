package server

import (
	"testing"

	"github.com/alexyer/ghost/ghost"
	"github.com/alexyer/ghost/protocol"
)

var (
	storage = ghost.GetStorage()
	c       = &client{
		Server:     &Server{storage: storage},
		collection: storage.GetCollection("main"),
	}
)

func testCmd(cmd *protocol.Command) ([]string, error) {
	return c.execCmd(cmd)
}

func TestPingCmd(t *testing.T) {
	cmdId := protocol.CommandId_PING

	result, err := c.execCmd(&protocol.Command{CommandId: &cmdId})
	if result[0] != "Pong!" {
		t.Errorf("Ping: wrong result: %v, %v", result, err)
	}
}

func TestSetCmd(t *testing.T) {
	c.collection.Del("test_key")
	cmdId := protocol.CommandId_SET

	_, err := c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"key1"},
	})

	if err == nil {
		t.Error("SET: wrong arguments: expected error, got <nil>")
	}

	_, err = c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_key", "test_val"},
	})

	if val, _ := c.collection.Get("test_key"); val != "test_val" {
		t.Error("SET: value hasn't been set")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestGetCmd(t *testing.T) {
	c.collection.Del("test_key")
	cmdId := protocol.CommandId_GET

	_, err := c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{},
	})

	if err == nil {
		t.Error("GET: wrong arguments: expected error, got <nil>")
	}

	_, err = c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_key"},
	})

	if err == nil {
		t.Error("GET: collection get: expected error, got <nil>")
	}

	c.collection.Set("test_key", "test_val")

	result, _ := c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_key"},
	})

	if result[0] != "test_val" {
		t.Error("GET: value hasn't been gotten")
	}
}

func TestDelCmd(t *testing.T) {
	c.collection.Set("test_key", "test_val")
	cmdId := protocol.CommandId_DEL

	_, err := c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{},
	})

	if err == nil {
		t.Error("SET: wrong arguments: expected error, got <nil>")
	}

	_, err = c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_key"},
	})

	if _, err := c.collection.Get("test_key"); err == nil {
		t.Error("SET: value hasn't been deleted")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestCGetCmd(t *testing.T) {
	cmdId := protocol.CommandId_CGET

	_, err := c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{},
	})

	if err == nil {
		t.Error("CGET: wrong arguments: expected error, got <nil>")
	}

	_, err = c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_collection"},
	})

	if err == nil {
		t.Error("CGET: get collection: expected error, got <nil>")
	}

	c.Server.storage.AddCollection("test_collection")

	_, err = c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_collection"},
	})

	if err != nil {
		t.Errorf("CGET: collection hasn't been changed: %s", err)
	}
}

func TestCAddCmd(t *testing.T) {
	c.Server.storage.DelCollection("test_collection")
	cmdId := protocol.CommandId_CADD

	_, err := c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{},
	})

	if err == nil {
		t.Error("CADD: wrong arguments: expected error, got <nil>")
	}

	_, err = c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_collection"},
	})

	if err != nil {
		t.Errorf("CADD: add collection: %s", err)
	}

	c.Server.storage.AddCollection("test_collection")

	_, err = c.execCmd(&protocol.Command{
		CommandId: &cmdId,
		Args:      []string{"test_collection"},
	})

	if err == nil {
		t.Errorf("CADD: add collection: %s", err)
	}
}
