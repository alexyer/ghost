package server

import (
	"github.com/alexyer/ghost/protocol"
	"github.com/alexyer/ghost/util"
)

// PING command.
func (c *client) Ping() ([]string, error) {
	return []string{"Pong!"}, nil
}

// SET command.
// SET <key> <val>
func (c *client) Set(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 2 {
		return nil, util.GhostCmdError("SET", "wrong arguments")
	}

	c.collection.Set(cmd.Args[0], cmd.Args[1])
	return nil, nil
}

// GET command.
// GET <key>
func (c *client) Get(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdError("GET", "wrong arguments")
	}

	val, err := c.collection.Get(cmd.Args[0])

	if err != nil {
		return nil, util.GhostCmdError("GET", err.Error())
	}

	return []string{val}, nil
}

// DEL command.
// DEL <key>
func (c *client) Del(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdError("DEL", "wrong arguments")
	}

	c.collection.Del(cmd.Args[0])
	return nil, nil
}

// CGET command.
// CGET <collection name>
// Change user's collection.
func (c *client) CGet(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdError("CGET", "wrong arguments")
	}

	newCollection := c.Server.storage.GetCollection(cmd.Args[0])

	if newCollection == nil {
		return nil, util.GhostCmdError("CGET", "collection does not exist")
	}

	c.collection = newCollection

	return nil, nil
}

// CADD command.
// CADD <collection name>
// Add new collection.
func (c *client) CAdd(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdError("CADD", "wrong arguments")
	}

	_, err := c.Server.storage.AddCollection(cmd.Args[0])

	if err != nil {
		return nil, util.GhostCmdError("CADD", err.Error())
	}

	return nil, nil
}
