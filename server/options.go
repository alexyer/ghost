package server

type Options struct {
	// host:port address.
	Addr string

	// Unix socket filename.
	Socket string

	// Log file location.
	LogfileName string
}

func (opt *Options) GetAddr() string {
	if opt.Addr == "" {
		opt.Addr = "localhost:6869"
	}

	return opt.Addr
}

func (opt *Options) GetLogfileName() string {
	if opt.LogfileName == "" {
		opt.LogfileName = "/tmp/ghost.log"
	}

	return opt.LogfileName
}

func (opt *Options) GetSocket() string {
	return opt.Socket
}
