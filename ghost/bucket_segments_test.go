package ghost

import (
	"testing"
	"unsafe"
)

func TestBucketSegments(t *testing.T) {
	bs := NewBucketSegments()

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
	bs.setBucket(SEGMENT_SIZE+1, n)

	bucket := bs.getBucket(SEGMENT_SIZE + 1)

	if bucket.head != unsafe.Pointer(n) {
		t.Fatalf("wrong bucket. got: %p, expected %p", bucket.head, unsafe.Pointer(n))
	}
}
