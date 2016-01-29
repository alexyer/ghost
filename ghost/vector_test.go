package ghost

import "testing"

func TestVectorPush(t *testing.T) {
	vec := vector{}
	vec.Push(node{
		Key: "key",
		Val: "val",
	})
	vec.Push(node{
		Key: "key1",
		Val: "val1",
	})
	vec.Push(node{
		Key: "key2",
		Val: "val2",
	})
	vec.Push(node{
		Key: "key3",
		Val: "val3",
	})

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
