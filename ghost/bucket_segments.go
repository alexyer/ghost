package ghost

import (
	"sync/atomic"
	"unsafe"
)

const (
	INITIAL_SEGMENT_SIZE uint32 = 8  // Size of the first segment
	SEGMENTS_NUM         uint32 = 30 // Maximum number of the segments
)

type bucketSegments struct {
	segments unsafe.Pointer // Pointer to array of segments
}

type segment struct {
	buckets unsafe.Pointer // Pointer to array of buckets
}

func newSegment(segmentSize uint32) unsafe.Pointer {
	buckets := make([]*bucket, segmentSize)

	return unsafe.Pointer(&segment{
		buckets: unsafe.Pointer(&buckets),
	})
}

func newBucketSegments() *bucketSegments {
	segments := make([]unsafe.Pointer, SEGMENTS_NUM)
	segments[0] = newSegment(INITIAL_SEGMENT_SIZE)

	return &bucketSegments{
		segments: unsafe.Pointer(&segments),
	}
}

func (bs *bucketSegments) getBucket(i uint32) *bucket {
	segmentIndex, bucketIndex := bs.at(i)

	segments := *(*[]unsafe.Pointer)(bs.segments)
	if segments[segmentIndex] == nil {
		return nil
	}

	seg := (*segment)(segments[segmentIndex])
	buckets := *(*[]*bucket)(seg.buckets)

	return buckets[bucketIndex]
}

func (bs *bucketSegments) setBucket(i uint32, newBucket *bucket) {
	segmentIndex, bucketIndex := bs.at(i)

	segments := *(*[]unsafe.Pointer)(bs.segments)

	if segments[segmentIndex] == nil {
		atomic.CompareAndSwapPointer(&segments[segmentIndex], nil, unsafe.Pointer(newSegment(INITIAL_SEGMENT_SIZE*(2<<segmentIndex-1))))
	}

	seg := (*segment)(segments[segmentIndex])
	buckets := *(*[]*bucket)(seg.buckets)
	buckets[bucketIndex] = newBucket
}

func (bs *bucketSegments) at(i uint32) (segmentIndex uint32, bucketIndex uint32) {
	pos := i + INITIAL_SEGMENT_SIZE
	hibit := bsr(pos)

	segmentIndex = hibit - bsr(INITIAL_SEGMENT_SIZE)
	bucketIndex = pos ^ (2<<hibit - 1)
	return
}
