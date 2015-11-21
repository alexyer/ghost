package client

const (
	PONG int = iota
)

type processor struct {
	process func(cmd Cmder)
}

func (p *processor) Process(cmd Cmder) {
	p.process(cmd)
}

func (p *processor) Ping() {
	cmd := &Cmd{
		Cmd: PONG,
	}

	p.Process(cmd)
}
