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

type bucket []*node

type hashMap struct {
	Count   uint32 // Number of elements in hashmap
	Size    uint32 // Number of buckets in hashmap
	hash    hash.Hash32
	buckets []bucket
}

func NewHashMap() *hashMap {
	newTable := &hashMap{}

	newTable.buckets = make([]bucket, initSize)

	for i := range newTable.buckets {
		newTable.buckets[i] = make(bucket, 2)
	}

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
	bucketIndex := h.find(key, h.buckets[index])

	if bucketIndex < 0 {
		h.buckets[index] = append(h.buckets[index], &node{key, val})
		h.Count++
	} else {
		h.buckets[index][bucketIndex].Val = val
	}
}

// Get element from the hashmap.
// Return error if value is not found.
func (h *hashMap) Get(key string) (string, error) {
	index := h.getIndex(key)

	bucketIndex := h.find(key, h.buckets[index])

	if bucketIndex < 0 {
		return "", errors.New("No value")
	} else {
		return h.buckets[index][bucketIndex].Val, nil
	}
}

// Delete element from the hashmap.
func (h *hashMap) Del(key string) {
	index := h.getIndex(key)
	bucketIndex := h.find(key, h.buckets[index])

	if bucketIndex < 0 {
		return
	}

	h.buckets[index][bucketIndex] = h.buckets[index][len(h.buckets[index])-1]
	h.buckets[index][len(h.buckets[index])-1] = nil
	h.buckets[index] = h.buckets[index][:len(h.buckets[index])-1]

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

	for i := range newBuckets {
		newBuckets[i] = make(bucket, 2)
	}

	for n := range h.nodes() {
		index := h.getIndex(n.Key)
		newBuckets[index] = append(newBuckets[index], n)
	}

	h.buckets = newBuckets
}

// Navigate through all nodes
func (h *hashMap) nodes() <-chan *node {
	ch := make(chan *node)

	go func() {
		for _, b := range h.buckets {
			for _, n := range b {
				if n != nil {
					ch <- n
				}
			}
		}
		close(ch)
	}()

	return ch
}

// Find index of the node in bucket or -1.
func (h *hashMap) find(key string, b bucket) int {
	for i := range b {
		if b[i] != nil && b[i].Key == key {
			return i
		}
	}

	return -1
}

// Get index of bucket key belongs to.
func (h *hashMap) getIndex(key string) uint32 {
	h.hash.Reset()
	h.hash.Write([]byte(key))
	return h.hash.Sum32() % h.Size
}
