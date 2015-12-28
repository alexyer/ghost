package client

import "github.com/alexyer/ghost/protocol"

type processor struct {
	process func(cmd *protocol.Command) (*protocol.Reply, error)
}

func (p *processor) Process(cmd *protocol.Command) {
	p.process(cmd)
}

func (p *processor) Ping() (*protocol.Reply, error) {
	cmdId := protocol.CommandId_PING

	cmd := &protocol.Command{
		CommandId: &cmdId,
	}

	reply, err := p.process(cmd)
	return reply, err
}

func (p *processor) Set(key, val string) (*protocol.Reply, error) {
	cmdId := protocol.CommandId_SET

	cmd := &protocol.Command{
		CommandId: &cmdId,
		Args:      []string{key, val},
	}

	reply, err := p.process(cmd)
	return reply, err
}
