package ghost

import "testing"

func TestVectorPush(t *testing.T) {
	vec := vector{}
	vec.Push(node{"key", "val"})
	vec.Push(node{"key1", "val1"})
	vec.Push(node{"key2", "val2"})
	vec.Push(node{"key3", "val3"})

	index := vec.Find("key1")

	if index == -1 {
		t.Errorf("Expected index, got -1")
	}

	vec.Pop(index)

	index = vec.Find("key1")

	if index != -1 {
		t.Errorf("Wrong pop.")
	}
}
