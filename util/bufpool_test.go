package util

import "testing"

func TestNewBufpool(t *testing.T) {
	bp := NewBufpool()

	if bp == nil {
		t.Fatal("Expected bufpool. Got <nil>")
	}

	if bp.maxSize != BUFPOOL_INIT_SIZE*(2<<uint(BUFPOOL_INIT_NUM-1)-1) {
		t.Error("Wrong maxSize")
	}
}

func TestGet(t *testing.T) {
	bp := NewBufpool()

	buf := bp.Get(42)

	if buf == nil {
		t.Fatal("Expected buffer. Got <nil>")
	}

	if len(buf) < 42 {
		t.Errorf("Wrong buffer length. Expected: 42, got: %d", len(buf))
	}
}

func TestPut(t *testing.T) {
	bp := NewBufpool()
	oldSize := bp.maxSize

	bp.Put(make([]byte, oldSize+1))

	if bp.maxSize <= oldSize {
		t.Errorf("Should grow")
	}
}

func TestGrow(t *testing.T) {
	bp := NewBufpool()
	oldSize := bp.maxSize

	buf := bp.Get(oldSize + 1)

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
	bp := NewBufpool()
	size := bp.maxSize

	go bp.Get(size + 1)
	go bp.Get(2 * size)
}

func TestPutRaces(t *testing.T) {
	bp := NewBufpool()
	size := bp.maxSize

	go bp.Put(make([]byte, size+1))
	go bp.Put(make([]byte, 2*size))
}

func TestPutBigBuf(t *testing.T) {
	bp := NewBufpool()
	bp.Put(make([]byte, 1024*1024))

	buf := bp.Get(1024 * 1024)

	if len(buf) < 1024*1024 {
		t.Errorf("Wrong buffer length. Expected: %d, got: %d", 1024*1024, len(buf))
	}
}
