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
