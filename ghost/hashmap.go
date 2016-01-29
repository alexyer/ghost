// Implementation of Striped Hashmap data structure.
package ghost

import (
	"sync"
	"time"
)

const (
	INIT_SIZE uint32  = 64   // Default number of buckets
	THRESHOLD float32 = 0.75 // Threshold load factor to rehash table
	LOCKS_NUM         = 1024 // Size of the lock array
)

type node struct {
	Key            string
	Val            string
	expire         bool // Expiration flag. Expire at expirationDate if 'true'
	expirationDate time.Time
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
	return &hashMap{
		buckets: make([]bucket, INIT_SIZE),
		locks:   make([]sync.Mutex, LOCKS_NUM),
		Size:    INIT_SIZE,
	}
}

// Set or update key.
func (h *hashMap) Set(key, val string) {
	if h.loadFactor() >= THRESHOLD {
		h.rehash()
	}

	index := h.getIndex(key)

	h.acquire(index)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		h.buckets[index].Push(node{
			Key: key,
			Val: val,
		})

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
		return "", NoValueErr
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
		h.release(index)
		return
	}

	h.buckets[index].Pop(bucketIndex)

	h.CountMu.Lock()
	h.Count--
	h.CountMu.Unlock()

	h.release(index)
}

// Set expiration date.
// ttl - time to live of the key in seconds.
func (h *hashMap) Expire(key string, ttl int) error {
	index := h.getIndex(key)

	h.acquire(index)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		h.release(index)
		return NoValueErr
	} else {
		h.buckets[index].Nodes[bucketIndex].expirationDate = time.Now().Add(time.Duration(ttl) * time.Second)
		h.buckets[index].Nodes[bucketIndex].expire = true
		h.release(index)

		return nil
	}
}

// Show ttl of the key.
func (h *hashMap) TTL(key string) (int, error) {
	index := h.getIndex(key)

	h.acquire(index)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		h.release(index)
		return -1, NoValueErr
	}

	if !h.buckets[index].Nodes[bucketIndex].expire {
		h.release(index)
		return -1, nil
	}

	ttl := int(h.buckets[index].Nodes[bucketIndex].expirationDate.Sub(time.Now()).Seconds())
	h.release(index)

	return ttl, nil
}

// Remove the existing timeout on key.
func (h *hashMap) Persist(key string) error {
	index := h.getIndex(key)

	h.acquire(index)

	bucketIndex := h.buckets[index].Find(key)

	if bucketIndex < 0 {
		h.release(index)
		return NoValueErr
	} else {
		h.buckets[index].Nodes[bucketIndex].expire = false
		h.release(index)
		return nil
	}
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

func (h *hashMap) acquireAll() {
	for i := 0; i < len(h.locks); i++ {
		h.locks[i].Lock()
	}
}

func (h *hashMap) releaseAll() {
	for i := len(h.locks) - 1; i >= 0; i-- {
		h.locks[i].Unlock()
	}
}

// Allocate new bigger hashmap and rehash all keys.
func (h *hashMap) rehash() {
	oldSize := h.Size

	h.acquireAll()

	if oldSize != h.Size {
		h.releaseAll()
		return // Someone beat us to it
	}

	h.Size <<= 1
	newBuckets := make([]bucket, h.Size)

	for n := range h.nodes() {
		newBuckets[h.getIndex(n.Key)].Push(n)
	}

	h.buckets = newBuckets

	h.releaseAll()
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
	return FNV1a_32([]byte(key)) & (h.Size - 1)
}
