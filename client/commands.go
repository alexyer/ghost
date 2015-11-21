package client

import "github.com/alexyer/ghost/protocol"

type processor struct {
	process func(cmd *protocol.Command)
}

func (p *processor) Process(cmd *protocol.Command) {
	p.process(cmd)
}

func (p *processor) Ping() *protocol.Command {
	cmdId := protocol.CommandId_PING

	cmd := &protocol.Command{
		CommandId: &cmdId,
	}

	p.Process(cmd)

	return cmd
}
