package ghost

import (
	"testing"
	"unsafe"
)

func TestBucketSegments(t *testing.T) {
	bs := newBucketSegments()

	if bs == nil {
		t.Fatal("wrong constructor")
	}

	i := 42
	ptr := unsafe.Pointer(&i)

	n := &node{
		Key:  42,
		Val:  "test",
		next: ptr,
	}

	b := &bucket{
		head: unsafe.Pointer(n),
	}

	bs.setBucket(SEGMENT_SIZE+1, b)

	bucket := bs.getBucket(SEGMENT_SIZE + 1)

	if bucket != b {
		t.Fatalf("wrong bucket. got: %p, expected %p", bucket, b)
	}
}
