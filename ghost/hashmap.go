// Implementation of Striped Hashmap data structure.
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
	Len     uint32    // Number of elements in hashmap
	Cap     uint32    // Number of buckets in hashmap
	buckets []*bucket // Array of indiviual buckets in hashmap
}

func NewHashMap(capacity int) *hashMap {
	newHash := &hashMap{
		Len:     0,
		Cap:     2,
		buckets: make([]*bucket, capacity),
	}

	tail := &node{
		Key:  ^uint32(0),
		next: nil,
	}

	head := &node{
		Key:  0,
		next: unsafe.Pointer(tail),
	}

	newHash.buckets[0] = &bucket{
		head: unsafe.Pointer(head),
	}
	return newHash
}

// Set or update key.
func (h *hashMap) Set(strKey, val string) {
	key := GetHash(strKey)

	node := &node{
		Key: Regularkey(key),
		Val: val,
	}

	bucketIndex := key & (atomic.LoadUint32(&h.Cap) - 1)

	if h.buckets[bucketIndex] == nil {
		h.initializeBucket(bucketIndex)
	}

	if h.buckets[bucketIndex].add(node) {
		if float64(atomic.AddUint32(&h.Len, 1))/float64(atomic.LoadUint32(&h.Cap)) > THRESHOLD {
			atomic.StoreUint32(&h.Cap, h.Cap<<1)
		}
	}
}

// Get element from the hashmap.
// Return error if value is not found.
func (h *hashMap) Get(strKey string) (string, error) {
	key := GetHash(strKey)

	bucketIndex := key & (atomic.LoadUint32(&h.Cap) - 1)

	if h.buckets[bucketIndex] == nil {
		h.initializeBucket(bucketIndex)
	}

	item := h.buckets[bucketIndex].get(Regularkey(key))

	if item == nil {
		return "", errors.New("ghost: no such key")
	}

	return item.Val, nil
}

// Delete element from the hashmap.
func (h *hashMap) Del(strKey string) {
	key := GetHash(strKey)

	bucketIndex := key & (atomic.LoadUint32(&h.Cap) - 1)

	if h.buckets[bucketIndex] == nil {
		h.initializeBucket(bucketIndex)
	}

	h.buckets[bucketIndex].remove(Regularkey(key))
}

// The role of initializeBucket is to direct the pointer
// in the array cell of the index bucket.
func (h *hashMap) initializeBucket(index uint32) {
	parentIndex := h.getParentIndex(index)

	if h.buckets[parentIndex] == nil {
		h.initializeBucket(parentIndex)
	}

	dummy := h.buckets[parentIndex].getDummy(index)

	if dummy != nil {
		h.buckets[index] = dummy
	}
}

func (h *hashMap) getParentIndex(bucketIndex uint32) uint32 {
	parentIndex := atomic.LoadUint32(&h.Cap)

	for parentIndex > bucketIndex {
		parentIndex = parentIndex >> 1
	}

	return bucketIndex - parentIndex
}
