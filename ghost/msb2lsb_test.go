package ghost

import "testing"

var r uint32

func TestMsb2lsb(t *testing.T) {
	if msb2lsb(3) != 3221225472 {
		t.Fatalf("expected: 3221225472, got: %d", msb2lsb(3))
	}
}

func BenchmarkMsb2lsb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = msb2lsb(uint32(i))
	}
}
