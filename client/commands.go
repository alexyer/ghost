package client

const (
	PONG int = iota
)

type processor struct {
	process func(cmd *Cmd)
}

func (p *processor) Process(cmd *Cmd) {
	p.process(cmd)
}

func (p *processor) Ping() *Cmd {
	cmd := &Cmd{
		Cmd: PONG,
	}

	p.Process(cmd)

	return cmd
}
