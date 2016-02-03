package server

import (
	"strconv"

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
		return nil, util.GhostCmdWrongArgsError("SET")
	}

	c.collection.Set(cmd.Args[0], cmd.Args[1])
	return nil, nil
}

// GET command.
// GET <key>
func (c *client) Get(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdWrongArgsError("GET")
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
		return nil, util.GhostCmdWrongArgsError("DEL")
	}

	c.collection.Del(cmd.Args[0])
	return nil, nil
}

// CGET command.
// CGET <collection name>
// Change user's collection.
func (c *client) CGet(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdWrongArgsError("CGET")
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
		return nil, util.GhostCmdWrongArgsError("CADD")
	}

	_, err := c.Server.storage.AddCollection(cmd.Args[0])

	if err != nil {
		return nil, util.GhostCmdError("CADD", err.Error())
	}

	return nil, nil
}

// EXPIRE command.
// EXPIRE <key> <seconds>
// Set expiration time of the key.
func (c *client) Expire(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 2 {
		return nil, util.GhostCmdWrongArgsError("EXPIRE")
	}

	ttl, err := strconv.Atoi(cmd.Args[1])
	if err != nil {
		return nil, util.GhostCmdError("EXPIRE", err.Error())
	}

	if err := c.collection.Expire(cmd.Args[0], ttl); err != nil {
		return nil, util.GhostCmdError("EXPIRE", err.Error())
	} else {
		return nil, nil
	}
}

// TTL command.
// TTL <key>
// Get expiration time of the key.
func (c *client) TTL(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdWrongArgsError("TTL")
	}

	if ttl, err := c.collection.TTL(cmd.Args[0]); err != nil {
		return []string{strconv.Itoa(ttl)}, util.GhostCmdError("TTL", err.Error())
	} else {
		return []string{strconv.Itoa(ttl)}, nil
	}
}

// PERSIST command.
// PERSIST <key>
// Remove the existing timeout of the key.
func (c *client) Persist(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, util.GhostCmdWrongArgsError("PERSIST")
	}

	if err := c.collection.Persist(cmd.Args[0]); err != nil {
		return nil, util.GhostCmdError("PERSIST", err.Error())
	} else {
		return nil, nil
	}
}
