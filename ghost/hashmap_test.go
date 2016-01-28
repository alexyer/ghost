package ghost

import (
	"fmt"
	"sync"
	"testing"
)

const (
	PARALLEL_TEST_TRY_NUM = 10000
	PARALLEL_TEST_RETRIES = 1
	MAXSIZE               = 82
)

var (
	empty   *hashMap
	one     *hashMap
	several *hashMap
	many    *hashMap
	result  string
)

func init() {
	empty = NewHashMap()
	one = NewHashMap()
	several = NewHashMap()
	many = NewHashMap()

	one.Set("One", "one")

	several.Set("One", "one")
	several.Set("Two", "two")

	for i := 0; i < MAXSIZE; i++ {
		many.Set(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
	}
}

func TestHashGet(t *testing.T) {
	if i, _ := one.Get("One"); i != "one" {
		t.Errorf("One hash Get failed")
	}

	first, _ := several.Get("One")
	second, _ := several.Get("Two")
	if first != "one" || second != "two" {
		t.Errorf("Several hash Get failed")
	}

	for i := 0; i < MAXSIZE; i++ {
		if j, _ := many.Get(fmt.Sprintf("%d", i)); fmt.Sprintf("%d", i) != j {
			t.Errorf("Many hash Get failed. Expected: %d, Got: %d", i, j)
		}
	}
}

func TestHashGetNotInTable(t *testing.T) {
	if _, err := empty.Get("One"); err == nil {
		t.Errorf("Wrong behavior")
	}

	if _, err := one.Get("Two"); err == nil {
		t.Errorf("Wrong behavior")
	}

	if _, err := several.Get("three"); err == nil {
		t.Errorf("Wrong behavior")
	}
}

func TestHashSetChangeValue(t *testing.T) {
	one.Set("One", "ten")

	if val, _ := one.Get("One"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}

	several.Set("Two", "ten")

	if val, _ := several.Get("Two"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}
}

func TestHashSetValue(t *testing.T) {
	one.Set("Two", "ten")

	if val, _ := one.Get("Two"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}

	several.Set("Ten", "ten")

	if val, _ := several.Get("Ten"); val != "ten" {
		t.Errorf("Wrong update behavior")
	}
}

func TestHashDelete(t *testing.T) {
	one.Del("One")

	if res, err := one.Get("One"); err == nil {
		t.Errorf("Wrong delete behavior. Got: %s", res)
	}

	several.Del("Two")

	if _, err := several.Get("Two"); err == nil {
		t.Errorf("Wrong delete behavior")
	}
}

func TestHashDeleteParallel(t *testing.T) {
	for retry := 0; retry < PARALLEL_TEST_RETRIES; retry++ {
		var wg sync.WaitGroup
		h := NewHashMap()

		for i := 0; i < PARALLEL_TEST_TRY_NUM; i++ {
			h.Set(string(i), "val"+string(i))
		}

		for i := 0; i < PARALLEL_TEST_TRY_NUM; i++ {
			wg.Add(1)
			go func() {
				h.Del(string(i))
				wg.Done()
			}()
		}

		wg.Wait()

		for i := 0; i < PARALLEL_TEST_TRY_NUM; i++ {
			if val, err := h.Get(string(i)); err == nil && val == "val"+string(i) {
				t.Fatal("Wrong delete behavior.")
			}
		}
	}
}
