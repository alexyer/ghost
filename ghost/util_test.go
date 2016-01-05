package ghost

import "testing"

func TestMsb2lsb(t *testing.T) {
	if msb2lsb(3) != 3221225472 {
		t.Fatalf("expected: 3221225472, got: %d", msb2lsb(3))
	}
}
