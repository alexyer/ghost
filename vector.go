package ghost

const initVectorSize = 2

type vector struct {
	Nodes []node
	count int
}

// Push element to vector.
func (v *vector) Push(n node) {
	if v.Nodes == nil {
		v.Nodes = make([]node, initSize)
	}

	if v.count == len(v.Nodes) {
		v.grow()
	}

	v.Nodes[v.count] = n
	v.count++
}

// Pop element from vector.
func (v *vector) Pop(i int) {
	copy(v.Nodes[i:], v.Nodes[i+1:])
	v.count--
}

// Find node index in vector.
func (v *vector) Find(key string) int {
	if v != nil {
		for i := 0; i < v.count; i++ {
			if v.Nodes[i].Key == key {
				return i
			}
		}
	}

	return -1
}

// Grow vector if limit achieved.
func (v *vector) grow() {
	newNodes := make([]node, len(v.Nodes)<<1)
	copy(newNodes, v.Nodes)
	v.Nodes = newNodes
}
