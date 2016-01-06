package ghost

import (
	"sync/atomic"
	"unsafe"
)

const SEGMENT_SIZE = 32

type bucketSegments struct {
	segments unsafe.Pointer // Pointer to array of segments
	length   uint32
}

type segment struct {
	buckets unsafe.Pointer // Pointer to array of buckets
}

func newSegment() unsafe.Pointer {
	buckets := make([]*bucket, SEGMENT_SIZE)

	return unsafe.Pointer(&segment{
		buckets: unsafe.Pointer(&buckets),
	})
}

func NewBucketSegments() *bucketSegments {
	segments := make([]unsafe.Pointer, 1)
	segments[0] = newSegment()

	return &bucketSegments{
		segments: unsafe.Pointer(&segments),
		length:   1,
	}
}

func (bs *bucketSegments) getBucket(bucketIndex uint32) *bucket {
	segmentIndex := bucketIndex / SEGMENT_SIZE

	if segmentIndex >= bs.length {
		return nil
	}

	segments := *(*[]unsafe.Pointer)(bs.segments)

	if segments[segmentIndex] == nil {
		return nil
	}

	seg := (*segment)(segments[segmentIndex])
	buckets := *(*[]*bucket)(seg.buckets)

	return buckets[bucketIndex&(SEGMENT_SIZE-1)]
}

func (bs *bucketSegments) setBucket(bucketIndex uint32, head *node) {
	segmentIndex := bucketIndex / SEGMENT_SIZE

	for {
		if segmentIndex >= bs.length {
			oldSegments := *(*[]*segment)(bs.segments)

			newSegments := append(oldSegments, make([]*segment, bs.length<<1)...)
			if atomic.CompareAndSwapPointer(&bs.segments, unsafe.Pointer(bs.segments), unsafe.Pointer(&newSegments)) {
				atomic.AddUint32(&bs.length, 1)
				break
			}
		}
	}

	segments := *(*[]unsafe.Pointer)(bs.segments)

	if (segments[segmentIndex]) == nil {
		atomic.CompareAndSwapPointer(&segments[segmentIndex], nil, unsafe.Pointer(newSegment()))
	}

	seg := (*segment)(segments[segmentIndex])
	buckets := *(*[]*bucket)(seg.buckets)
	buckets[bucketIndex&(SEGMENT_SIZE-1)] = &bucket{
		head: unsafe.Pointer(head),
	}
}
