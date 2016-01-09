// Implementation of Lock-free split ordered hashmap.
// For details refer to Shalev & Shavit "Split-Ordered Lists - Lock-free Resizable Hash Tables" work.
package ghost

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

const (
	THRESHOLD float64 = 0.75 // Threshold load factor to rehash table
)

type hashMap struct {
	Len            uint32          // Number of elements in hashmap
	Cap            uint32          // Number of buckets in hashmap
	bucketSegments *bucketSegments // Array of indiviual buckets in hashmap
}

func NewHashMap() *hashMap {
	newHash := &hashMap{
		Len:            0,
		Cap:            2,
		bucketSegments: newBucketSegments(),
	}

	tail := &node{
		Key:  ^uint32(0),
		next: nil,
	}

	head := &node{
		Key:  0,
		next: unsafe.Pointer(tail),
	}

	newHash.bucketSegments.setBucket(0, &bucket{
		head: unsafe.Pointer(head),
	})
	return newHash
}

// Set or update key.
func (h *hashMap) Set(strKey, val string) {
	key := GetHash(strKey)

	node := &node{
		Key: Regularkey(key),
		Val: val,
	}

	bucket := h.getBucket(key)

	if bucket.add(node) {
		if float64(atomic.AddUint32(&h.Len, 1))/float64(atomic.LoadUint32(&h.Cap)) > THRESHOLD {
			atomic.StoreUint32(&h.Cap, h.Cap<<1)
		}
	}
}

// Get element from the hashmap.
// Return error if value is not found.
func (h *hashMap) Get(strKey string) (string, error) {
	key := GetHash(strKey)

	bucket := h.getBucket(key)

	item := bucket.get(Regularkey(key))

	if item == nil {
		return "", errors.New("ghost: no such key")
	}

	return item.Val, nil
}

// Delete element from the hashmap.
func (h *hashMap) Del(strKey string) {
	key := GetHash(strKey)
	bucket := h.getBucket(key)
	bucket.remove(Regularkey(key))
}

// The role of initializeBucket is to direct the pointer
// in the array cell of the index bucket.
func (h *hashMap) initializeBucket(index uint32) *bucket {
	parentIndex := h.getParentIndex(index)

	if h.bucketSegments.getBucket(parentIndex) == nil {
		h.initializeBucket(parentIndex)
	}

	dummy := h.bucketSegments.getBucket(parentIndex).getDummy(index)

	if dummy != nil {
		h.bucketSegments.setBucket(index, dummy)
	}

	return dummy
}

func (h *hashMap) getParentIndex(bucketIndex uint32) uint32 {
	parentIndex := atomic.LoadUint32(&h.Cap)

	for parentIndex > bucketIndex {
		parentIndex = parentIndex >> 1
	}

	return bucketIndex - parentIndex
}

func (h *hashMap) getBucket(key uint32) *bucket {
	bucketIndex := key & (atomic.LoadUint32(&h.Cap) - 1)
	bucket := h.bucketSegments.getBucket(bucketIndex)

	if bucket == nil {
		bucket = h.initializeBucket(bucketIndex)
	}

	return bucket
}
