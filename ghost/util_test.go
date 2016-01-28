package ghost

import "testing"

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
