package ghost

import "encoding/binary"

// Implementation of FNV-1a hash alghorithm.
func FNV1a_32(data []byte) uint32 {
	var (
		val uint32 = 2166136261 // Offset basis
		i   uint
	)

	hash := make([]byte, 4) // Little Endian Hash value

	for _, v := range data {
		val ^= uint32(v) // xor the bottom with the current octet
		val *= 16777619  // multiply by the 32 bit FNV magic prime mod 2^32
	}

	for ; i < 4; i++ {
		hash[i] = byte(val >> (i * 8))
	}

	return binary.LittleEndian.Uint32(hash)
}
