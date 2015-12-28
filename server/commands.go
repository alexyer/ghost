package server

import "github.com/alexyer/ghost/protocol"

// PING command.
func (c *client) Ping() ([]string, error) {
	return []string{"Pong!"}, nil
}

// SET command.
// SET <key> <val>
func (c *client) Set(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 2 {
		return nil, GhostCmdError("SET", "wrong arguments")
	}

	c.collection.Set(cmd.Args[0], cmd.Args[1])
	return nil, nil
}

// GET command.
// GET <key>
func (c *client) Get(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, GhostCmdError("GET", "wrong arguments")
	}

	val, err := c.collection.Get(cmd.Args[0])

	if err != nil {
		return nil, GhostCmdError("GET", err.Error())
	}

	return []string{val}, nil
}

// DEL command.
// DEL <key>
func (c *client) Del(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, GhostCmdError("DEL", "wrong arguments")
	}

	c.collection.Del(cmd.Args[0])
	return nil, nil
}

// CGET command.
// CGET <collection name>
// Change user's collection.
func (c *client) CGet(cmd *protocol.Command) ([]string, error) {
	if len(cmd.Args) != 1 {
		return nil, GhostCmdError("CGET", "wrong arguments")
	}

	newCollection := c.Server.storage.GetCollection(cmd.Args[0])

	if newCollection == nil {
		return nil, GhostCmdError("CGET", "collection does not exist")
	}

	c.collection = newCollection

	return nil, nil
}
