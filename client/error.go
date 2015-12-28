package client

func isBadConn(c *conn, ei error) bool {
	if c.rd.Buffered() > 0 {
		return true
	}

	if ei == nil {
		return false
	}

	return true
}
