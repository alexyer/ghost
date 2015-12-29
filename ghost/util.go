package ghost

import "encoding/binary"

func IntToByteArray(val int64) []byte {
	buf := make([]byte, 8)
	binary.PutVarint(buf, val)
	return buf
}

func ByteArrayToUint64(bytes []byte) (int64, int) {
	return binary.Varint(bytes)
}
