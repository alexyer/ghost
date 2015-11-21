package client

type Cmder interface {
	Reset()
	Result() (string, error)
}

type Cmd struct {
	Cmd  int
	Args []string
	Err  error
	Val  string
}

func (c *Cmd) Reset() {
	c.Err = nil
	c.Val = ""
}

func (c *Cmd) Result() (string, error) {
	return c.Val, c.Err
}
