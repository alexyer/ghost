package ghost

import "testing"

const MAXSIZE = 42

var (
	empty   hashMap
	one     hashMap
	several hashMap
	many    hashMap
)

func init() {
	empty = newHashMap()
	one = newHashMap()
	several = newHashMap()
	many = newHashMap()

	one.set("One", "one")

	several.set("One", "one")
	several.set("Two", "two")

	for i := 0; i < MAXSIZE; i++ {
		many.set(string(i), string(i))
	}
}

func TestHashGet(t *testing.T) {
	if i, _ := one.get("One"); i != "one" {
		t.Errorf("One hash get failed")
	}

	first, _ := several.get("One")
	second, _ := several.get("Two")
	if first != "one" || second != "two" {
		t.Errorf("Several hash get failed")
	}

	for i := 0; i < MAXSIZE; i++ {
		if i, _ := many.get(string(i)); string(i) != string(i) {
			t.Errorf("Many hash get failed")
		}
	}
}

func TestHashGetNotInTable(t *testing.T) {
	if _, err := empty.get("One"); err == nil {
		t.Errorf("Wrong behavior")
	}

	if _, err := one.get("Two"); err == nil {
		t.Errorf("Wrong behavior")
	}

	if _, err := several.get("three"); err == nil {
		t.Errorf("Wrong behavior")
	}
}

func TestHashSetChangeValue(t *testing.T) {
	one.set("One", "ten")

	if val, _ := one.get("One"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}

	several.set("Two", "ten")

	if val, _ := several.get("Two"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}
}

func TestHashSetValue(t *testing.T) {
	one.set("Two", "ten")

	if val, _ := one.get("Two"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}

	several.set("Ten", "ten")

	if val, _ := several.get("Ten"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}
}

func TestHashDelete(t *testing.T) {
	one.del("One")

	if _, err := one.get("One"); err == nil {
		t.Errorf("Wrong delete behavior")
	}

	several.del("Two")

	if _, err := several.get("Two"); err == nil {
		t.Errorf("Wrong delete behavior")
	}
}
