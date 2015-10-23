package ghost

import (
	"errors"
	"hash/fnv"
)

const initSize uint32 = 2  // Default number of buckets
var currentSize uint32 = 0 // Current size of the hasmap

type node struct {
	Next *node
	Key  string
	Val  string
}

type hashMap struct {
	buckets []*node
}

func newHashMap() hashMap {
	currentSize = initSize

	newTable := hashMap{}
	newTable.buckets = make([]*node, currentSize)

	return newTable
}

// Set or update key
func (h *hashMap) set(key, val string) {
	index := getIndex(key)
	currentNode := h.find(key, h.buckets[index])

	if currentNode == nil {
		h.buckets[index] = &node{h.buckets[index], key, val}
	} else {
		currentNode.Val = val
	}
}

// Get element from the hashmap.
// Return error if value is not found.
func (h *hashMap) get(key string) (string, error) {
	tmp := h.find(key, h.buckets[getIndex(key)])

	if tmp != nil {
		return tmp.Val, nil
	} else {
		return "", errors.New("No value")
	}
}

// Delete element from the hashmap.
func (h *hashMap) del(key string) {
	index := getIndex(key)
	currentNode := h.buckets[index]
	var prev *node = nil

	for currentNode != nil {
		if currentNode.Key == key {
			if prev == nil {
				h.buckets[index] = currentNode.Next
			} else {
				prev.Next = currentNode.Next
			}

			return
		}

		prev, currentNode = currentNode, currentNode.Next
	}
}

func (h *hashMap) find(key string, node *node) *node {
	for node != nil {
		if node.Key == key {
			return node
		}
		node = node.Next
	}

	return nil
}

func getIndex(key string) uint32 {
	h := fnv.New32()
	h.Write([]byte(key))
	return h.Sum32() % currentSize
}
