package server

type Options struct {
	Addr             string
	ClientHeaderSize int
	ClientBufSize    int
}

func (opt *Options) GetAddr() string {
	if opt.Addr == "" {
		opt.Addr = "localhost:6869"
	}

	return opt.Addr
}

func (opt *Options) GetClientHeaderSize() int {
	if opt.ClientHeaderSize == 0 {
		opt.ClientHeaderSize = 8
	}

	return opt.ClientHeaderSize
}

func (opt *Options) GetClientBufSize() int {
	if opt.ClientBufSize == 0 {
		opt.ClientBufSize = 4096
	}

	return opt.ClientBufSize
}
