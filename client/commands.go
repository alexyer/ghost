package client

import (
	"errors"

	"github.com/alexyer/ghost/protocol"
)

type processor struct {
	process func(cmd *protocol.Command) (*protocol.Reply, error)
}

func (p *processor) Process(cmd *protocol.Command) {
	p.process(cmd)
}

// PING command.
func (p *processor) Ping() (*protocol.Reply, error) {
	cmdId := protocol.CommandId_PING

	cmd := &protocol.Command{
		CommandId: &cmdId,
	}

	reply, err := p.process(cmd)
	return reply, err
}

// SET command.
// SET <key> <val>
func (p *processor) Set(key, val string) {
	cmdId := protocol.CommandId_SET

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key, val},
	}

	p.process(cmd)
	return
}

// GET command.
// GET <key>
func (p *processor) Get(key string) (string, error) {
	cmdId := protocol.CommandId_GET

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key},
	}

	reply, err := p.process(cmd)

	if err != nil {
		return "", err
	}

	if *reply.Error != "" {
		return "", errors.New(*reply.Error)
	}

	return reply.Values[0], nil
}
