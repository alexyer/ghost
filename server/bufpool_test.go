package server

import "testing"

func TestNewBufpool(t *testing.T) {
	bp := newBufpool()

	if bp == nil {
		t.Fatal("Expected bufpool. Got <nil>")
	}

	if bp.maxSize != BUFPOOL_INIT_SIZE*(2<<uint(BUFPOOL_INIT_NUM)-1) {
		t.Error("Wrong maxSize")
	}
}

func TestGet(t *testing.T) {
	bp := newBufpool()

	buf := bp.get(42)

	if buf == nil {
		t.Fatal("Expected buffer. Got <nil>")
	}

	if len(buf) < 42 {
		t.Errorf("Wrong buffer length. Expected: 42, got: %d", len(buf))
	}
}

func TestPut(t *testing.T) {
	bp := newBufpool()
	oldSize := bp.maxSize

	bp.put(make([]byte, oldSize+1))

	if bp.maxSize <= oldSize {
		t.Errorf("Should grow")
	}
}

func TestGrow(t *testing.T) {
	bp := newBufpool()
	oldSize := bp.maxSize

	buf := bp.get(oldSize + 1)

	if buf == nil {
		t.Fatal("Expected buffer. Got <nil>")
	}

	if len(buf) < oldSize+1 {
		t.Errorf("Wrong buffer length. Expected: %d, got: %d", oldSize+1, len(buf))
	}

	if bp.maxSize <= oldSize {
		t.Errorf("Wrong maxSize")
	}
}

func TestGetRaces(t *testing.T) {
	bp := newBufpool()
	size := bp.maxSize

	go bp.get(size + 1)
	go bp.get(2 * size)
}

func TestPutRaces(t *testing.T) {
	bp := newBufpool()
	size := bp.maxSize

	go bp.put(make([]byte, size+1))
	go bp.put(make([]byte, 2*size))
}
