package ghost

import "testing"

var r uint32

func BenchmarkMsb2lsb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = msb2lsb(uint32(i))
	}
}
