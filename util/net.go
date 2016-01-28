package util

import "net"

// Read data of the given size from connection.
func ReadData(conn net.Conn, buf []byte, size int) (int, error) {
	var (
		totalBytesRead = 0
		readErr        error
		bytesRead      = 0
	)

	for totalBytesRead < size && readErr == nil {
		bytesRead, readErr = conn.Read(buf[totalBytesRead:])
		totalBytesRead += bytesRead
	}

	return totalBytesRead, readErr
}
