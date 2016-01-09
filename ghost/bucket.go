package ghost

import (
	"unsafe"

	"github.com/alexyer/taggedptr"
)

type bucket struct {
	head unsafe.Pointer
}

// Add item to list.
func (b *bucket) add(newNode *node) bool {

	for {
		pred, curr := b.find(newNode.Key)

		if (*node)(curr).Key == newNode.Key {
			newNode.next = (*node)(curr).next
		} else {
			newNode.next = curr
		}
		return taggedptr.CompareAndSwap(&(*node)(pred).next, curr, unsafe.Pointer(newNode), 0, 0)
	}
}

func (b *bucket) getDummy(index uint32) *bucket {
	key := Dummykey(index)

	for {
		pred, curr := b.find(key)

		if (*node)(curr).Key == key {
			return &bucket{
				head: curr,
			}
		} else {
			dummy := &node{
				Key:  key,
				next: (*node)(pred).next,
			}

			if taggedptr.CompareAndSwap(&(*node)(pred).next, curr, unsafe.Pointer(dummy), 0, 0) {
				return &bucket{
					head: unsafe.Pointer(dummy),
				}
			} else {
				continue
			}
		}
	}
	return nil
}

func (b *bucket) get(key uint32) *node {
	_, curr := b.find(key)

	if (*node)(curr).Key == key {
		return (*node)(curr)
	}
	return nil
}

// Remove item from the list.
func (b *bucket) remove(key uint32) bool {
	for {
		pred, curr := b.find(key)

		if (*node)(curr).Key != key {
			return false
		} else {
			succ := taggedptr.GetPointer((*node)(curr).next)

			if !taggedptr.AttemptTag(&(*node)(curr).next, succ, 1) {
				continue
			}

			taggedptr.CompareAndSwap(&(*node)(pred).next, curr, succ, 0, 0)
			return true
		}
	}
}

func (b *bucket) find(key uint32) (unsafe.Pointer, unsafe.Pointer) {
	pred := b.head
	curr := taggedptr.GetPointer((*node)(b.head).next)

Retry:
	for {
		succ, tag := taggedptr.Get((*node)(curr).next)

		for tag != 0 {
			if !taggedptr.CompareAndSwap(&(*node)(pred).next, curr, succ, 0, 0) {
				continue Retry
			}
			curr = succ
			succ = taggedptr.GetPointer((*node)(curr).next)
		}
		if (*node)(curr).Key >= key {
			return pred, curr
		}
		pred = curr
		curr = succ
	}
}
