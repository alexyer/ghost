// Naive implementation of Hashmap data structure.
package ghost

import (
	"errors"
	"hash"
	"hash/fnv"
)

const (
	initSize  uint32  = 8    // Default number of buckets
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
	Size    uint32 // Number of buckets in hashmap
	hash    hash.Hash32
	buckets []bucket
}

func NewHashMap() *hashMap {
	newTable := &hashMap{}

	newTable.buckets = make([]bucket, initSize)

	newTable.Size = initSize
	newTable.hash = fnv.New32a()

	return newTable
}

// Set or update key.
func (h *hashMap) Set(key, val string) {
	if h.loadFactor() >= threshold {
		h.rehash()
	}

	index := h.getIndex(key)
	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		h.buckets[index].Push(node{key, val})
		h.Count++
	} else {
		h.buckets[index].Nodes[bucketIndex].Val = val
	}
}

// Get element from the hashmap.
// Return error if value is not found.
func (h *hashMap) Get(key string) (string, error) {
	index := h.getIndex(key)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		return "", errors.New("No value")
	} else {
		return h.buckets[index].Nodes[bucketIndex].Val, nil
	}
}

// Delete element from the hashmap.
func (h *hashMap) Del(key string) {
	index := h.getIndex(key)
	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		return
	}

	h.buckets[index].Pop(bucketIndex)
	h.Count--
}

// Get current load factor.
func (h *hashMap) loadFactor() float32 {
	return float32(h.Count) / float32(h.Size)
}

// Allocate new bigger hashmap and rehash all keys.
func (h *hashMap) rehash() {
	h.Size <<= 1
	newBuckets := make([]bucket, h.Size)

	for n := range h.nodes() {
		newBuckets[h.getIndex(n.Key)].Push(n)
	}

	h.buckets = newBuckets
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
	h.hash.Reset()
	h.hash.Write([]byte(key))
	return h.hash.Sum32() % h.Size
}
