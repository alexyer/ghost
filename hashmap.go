// Naive implementation of Hashmap data structure.
package ghost

import (
	"errors"
	"hash/fnv"
)

const (
	initSize  uint32  = 2    // Default number of buckets
	threshold float32 = 0.75 // Threshold load factor to rehash table
)

type node struct {
	Next *node
	Key  string
	Val  string
}

type hashMap struct {
	Count   uint32 // Number of elements in hashmap
	Size    uint32 // Number of buckets in hashmap
	buckets []*node
}

func NewHashMap() *hashMap {
	newTable := &hashMap{}
	newTable.buckets = make([]*node, initSize)
	newTable.Size = initSize

	return newTable
}

// Set or update key.
func (h *hashMap) Set(key, val string) {
	if h.loadFactor() >= threshold {
		h.rehash()
	}

	index := h.getIndex(key)
	currentNode := h.find(key, h.buckets[index])

	if currentNode == nil {
		h.buckets[index] = &node{h.buckets[index], key, val}
		h.Count++
	} else {
		currentNode.Val = val
	}
}

// Get element from the hashmap.
// Return error if value is not found.
func (h *hashMap) Get(key string) (string, error) {
	tmp := h.find(key, h.buckets[h.getIndex(key)])

	if tmp != nil {
		return tmp.Val, nil
	} else {
		return "", errors.New("No value")
	}
}

// Delete element from the hashmap.
func (h *hashMap) Del(key string) {
	index := h.getIndex(key)
	currentNode := h.buckets[index]
	var prev *node = nil

	for currentNode != nil {
		if currentNode.Key == key {
			if prev == nil {
				h.buckets[index] = currentNode.Next
			} else {
				prev.Next = currentNode.Next
			}

			h.Count--
			return
		}

		prev, currentNode = currentNode, currentNode.Next
	}
}

// Get current load factor.
func (h *hashMap) loadFactor() float32 {
	return float32(h.Count) / float32(h.Size)
}

// Allocate new bigger hashmap and rehash all keys.
func (h *hashMap) rehash() {
	h.Size <<= 1
	newBuckets := make([]*node, h.Size)

	for n := range h.nodes() {
		index := h.getIndex(n.Key)

		if newBuckets[index] == nil {
			newBuckets[index] = &node{nil, n.Key, n.Val}
		} else {
			newBuckets[index] = &node{newBuckets[index], n.Key, n.Val}
		}
	}

	h.buckets = newBuckets
}

// Navigate through all nodes
func (h *hashMap) nodes() <-chan *node {
	ch := make(chan *node)

	go func() {
		for _, bucket := range h.buckets {
			currentNode := bucket

			for {
				if currentNode == nil {
					break
				}

				ch <- currentNode

				currentNode = currentNode.Next
			}
		}
		close(ch)
	}()

	return ch
}

// Find node.
func (h *hashMap) find(key string, node *node) *node {
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}

	return nil
}

// Get index of bucket key belongs to.
func (h *hashMap) getIndex(key string) uint32 {
	hash := fnv.New32()
	hash.Write([]byte(key))
	return hash.Sum32() % h.Size
}
