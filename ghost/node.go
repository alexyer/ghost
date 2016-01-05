package ghost

import "unsafe"

type node struct {
	Key  uint32
	Val  string
	next unsafe.Pointer // Pointer to the next node
}
