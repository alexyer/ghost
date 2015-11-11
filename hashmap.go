// Naive implementation of Hashmap data structure.
package ghost

import (
	"errors"
	"sync"
)

const (
	initSize  uint32  = 64   // Default number of buckets
	threshold float32 = 0.75 // Threshold load factor to rehash table
)

type node struct {
	Key string
	Val string
}

type bucket struct {
	vector
}

type hashMap struct {
	Count   uint32 // Number of elements in hashmap
	CountMu sync.Mutex
	Size    uint32       // Number of buckets in hashmap
	buckets []bucket     // Array of indiviual buckets in hashmap
	locks   []sync.Mutex // Array of locks. Used to syncronize bucket access
}

func NewHashMap() *hashMap {
	newTable := &hashMap{}

	newTable.buckets = make([]bucket, initSize)

	newTable.locks = make([]sync.Mutex, initSize)

	newTable.Size = initSize

	return newTable
}

// Set or update key.
func (h *hashMap) Set(key, val string) {
	if h.loadFactor() >= threshold {
		h.rehash()
	}

	index := h.getIndex(key)

	h.acquire(index)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		h.buckets[index].Push(node{key, val})

		h.CountMu.Lock()
		h.Count++
		h.CountMu.Unlock()
	} else {
		h.buckets[index].Nodes[bucketIndex].Val = val
	}

	h.release(index)
}

// Get element from the hashmap.
// Return error if value is not found.
func (h *hashMap) Get(key string) (string, error) {
	index := h.getIndex(key)

	h.acquire(index)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		h.release(index)
		return "", errors.New("No value")
	} else {
		val := h.buckets[index].Nodes[bucketIndex].Val
		h.release(index)

		return val, nil
	}
}

// Delete element from the hashmap.
func (h *hashMap) Del(key string) {
	index := h.getIndex(key)

	h.acquire(index)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		return
	}

	h.buckets[index].Pop(bucketIndex)

	h.CountMu.Lock()
	h.Count--
	h.CountMu.Unlock()

	h.release(index)
}

// Get current load factor.
func (h *hashMap) loadFactor() float32 {
	h.CountMu.Lock()
	factor := float32(h.Count) / float32(h.Size)
	h.CountMu.Unlock()

	return factor
}

// Acquire control on the bucket.
func (h *hashMap) acquire(index uint32) {
	h.locks[index%uint32(len(h.locks))].Lock()
}

// Release control on the bucket.
func (h *hashMap) release(index uint32) {
	h.locks[index%uint32(len(h.locks))].Unlock()
}

// Allocate new bigger hashmap and rehash all keys.
func (h *hashMap) rehash() {
	oldSize := h.Size

	for i := 0; i < len(h.locks); i++ {
		h.locks[i].Lock()
	}

	if oldSize != h.Size {
		return // Someone beat us to it
	}

	h.Size <<= 1
	newBuckets := make([]bucket, h.Size)

	for n := range h.nodes() {
		newBuckets[h.getIndex(n.Key)].Push(n)
	}

	h.buckets = newBuckets

	for i := len(h.locks) - 1; i >= 0; i-- {
		h.locks[i].Unlock()
	}
}

// Navigate through all nodes
func (h *hashMap) nodes() <-chan node {
	ch := make(chan node)

	go func() {
		for _, b := range h.buckets {
			for i := 0; i < b.count; i++ {
				ch <- b.Nodes[i]
			}
		}
		close(ch)
	}()

	return ch
}

// Get index of bucket key belongs to.
func (h *hashMap) getIndex(key string) uint32 {
	return FNV1a_32([]byte(key)) % h.Size
}
