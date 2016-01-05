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

func GetHash(key string) uint32 {
	return FNV1a_32([]byte(key))
}

func Regularkey(key uint32) uint32 {
	return msb2lsb(key | 0x80000000)
}

func Dummykey(key uint32) uint32 {
	return msb2lsb(key)
}
