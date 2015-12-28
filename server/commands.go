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
