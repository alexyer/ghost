package ghost

import "testing"

func TestMsb2lsb(t *testing.T) {
	if msb2lsb(3) != 3221225472 {
		t.Fatalf("expected: 3221225472, got: %d", msb2lsb(3))
	}
}

func TestBsr(t *testing.T) {
	switch {
	case Bsr(4) != 2:
		t.Fatal("error")
	case Bsr(2) != 1:
		t.Fatal("error")
	case Bsr(5) != 2:
		t.Fatal("error")
	}
}
