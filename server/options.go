package server

type Options struct {
	Addr string
}

func (opt *Options) GetAddr() string {
	if opt.Addr == "" {
		opt.Addr = "localhost:6869"
	}

	return opt.Addr
}
