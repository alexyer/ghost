package util

import (
	"sync"

	"github.com/alexyer/ghost/ghost"
)

const (
	BUFPOOL_INIT_NUM  = 8  // Initial number of pools
	BUFPOOL_INIT_SIZE = 32 // Size of the first buffer pool in bytes
)

// Buffers pool is an array of pools.
// Each pool contains buffers of specified size.
type Bufpool struct {
	poolsMu sync.RWMutex
	pools   []sync.Pool
	maxSize int
}

// Create new pools pool.
func NewBufpool() *Bufpool {
	return &Bufpool{
		pools:   make([]sync.Pool, BUFPOOL_INIT_NUM),
		maxSize: BUFPOOL_INIT_SIZE * (2<<uint(BUFPOOL_INIT_NUM) - 1),
	}
}

// Get a buffer capable to contain the given size from the pool.
func (bp *Bufpool) Get(size int) []byte {
	bp.poolsMu.RLock()

	if size > bp.maxSize {
		bp.poolsMu.RUnlock()
		bp.grow(size)
		bp.poolsMu.RLock()
	}

	poolIndex := bp.getPoolIndex(size)
	buffer := bp.pools[poolIndex].Get()

	bp.poolsMu.RUnlock()

	if buffer == nil {
		return make([]byte, BUFPOOL_INIT_SIZE*(2<<uint(poolIndex)-1))
	}

	return buffer.([]byte)
}

// Put buffer back into the pool.
func (bp *Bufpool) Put(buf []byte) {
	bp.poolsMu.RLock()

	size := len(buf)
	if size > bp.maxSize {
		bp.poolsMu.RUnlock()
		bp.grow(size)
		bp.poolsMu.RLock()
	}

	bp.poolsMu.RUnlock()

	bp.pools[bp.getPoolIndex(size)].Put(buf)
}

func (bp *Bufpool) grow(size int) {
	bp.poolsMu.Lock()

	// Somebody has been faster
	if size <= bp.maxSize {
		bp.poolsMu.Unlock()
		return
	}

	newLen := bp.getPoolIndex(size)
	newPools := make([]sync.Pool, newLen+1)

	copy(newPools, bp.pools)

	bp.pools = newPools
	bp.maxSize = BUFPOOL_INIT_SIZE * (2<<uint(newLen) - 1)

	bp.poolsMu.Unlock()
}

func (bp *Bufpool) getPoolIndex(size int) uint32 {
	return ghost.Bsr(uint32(size)+uint32(BUFPOOL_INIT_SIZE)) - ghost.Bsr(uint32(BUFPOOL_INIT_SIZE))
}
