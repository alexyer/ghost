package server

type Options struct {
	Addr          string
	MsgHeaderSize int
	MsgBufferSize int
}

func (opt *Options) GetAddr() string {
	if opt.Addr == "" {
		opt.Addr = "localhost:6869"
	}

	return opt.Addr
}

func (opt *Options) GetMsgHeaderSize() int {
	if opt.MsgHeaderSize == 0 {
		opt.MsgHeaderSize = 8
	}

	return opt.MsgHeaderSize
}

func (opt *Options) GetMsgBufferSize() int {
	if opt.MsgBufferSize == 0 {
		opt.MsgBufferSize = 4096
	}

	return opt.MsgBufferSize
}
