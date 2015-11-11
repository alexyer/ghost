package ghost

import (
	"fmt"
	"testing"
)

const MAXSIZE = 82

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

func BenchmarkSet(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		h.Set(string(i), "Yarrr")
	}
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	j := 0

	for i := 0; i < 100000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if j < 100000 {
			h.Get(string(i))
		}
	}
}

func BenchmarkDel(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	j := 0

	for i := 0; i < 100000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if j < 100000 {
			h.Del(string(i))
		}
	}
}

func BenchmarkNativeSet(b *testing.B) {
	b.StopTimer()
	h := make(map[string]string)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		h[string(i)] = "Yarrr"
	}
}

func BenchmarkNativeGet(b *testing.B) {
	b.StopTimer()
	var r string
	h := make(map[string]string)
	j := 0

	for i := 0; i < 1000000; i++ {
		h[string(i)] = "Yarrr"
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if j < 1000000 {
			r = h[string(i)]
		}
	}

	result = r
}

func BenchmarkNativeDel(b *testing.B) {
	b.StopTimer()
	h := make(map[string]string)
	j := 0

	for i := 0; i < 1000000; i++ {
		h[string(i)] = "Yarrr"
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if j < 1000000 {
			delete(h, string(i))
		}
	}
}

func ParallelSet(b *testing.B, i int, h *hashMap) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h.Set(string(i), "Yarrr")
			i++
		}
	})
}

func BenchmarkParallelSet(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	i := 0
	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet8(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	i := 0

	b.SetParallelism(8)

	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet64(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	i := 0

	b.SetParallelism(64)

	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet128(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	i := 0

	b.SetParallelism(128)

	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet1024(b *testing.B) {
	b.StopTimer()
	h := NewHashMap()
	i := 0

	b.SetParallelism(1024)

	b.StartTimer()

	ParallelSet(b, i, h)
}

func ParallelGet(b *testing.B, i int, h *hashMap) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h.Get(string(i))
			i++
		}
	})
}

func BenchmarkParallelGet(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.StartTimer()
	ParallelGet(b, i, h)
}

func BenchmarkParallelGet8(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(8)

	b.StartTimer()
	ParallelGet(b, i, h)
}

func BenchmarkParallelGet64(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(64)

	b.StartTimer()
	ParallelGet(b, i, h)
}

func BenchmarkParallelGet128(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(128)

	b.StartTimer()
	ParallelGet(b, i, h)
}

func BenchmarkParallelGet1024(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(1024)

	b.StartTimer()
	ParallelGet(b, i, h)
}

func ParallelDel(b *testing.B, i int, h *hashMap) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h.Del(string(i))
			i++
		}
	})
}

func BenchmarkParallelDel(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.StartTimer()
	ParallelDel(b, i, h)
}

func BenchmarkParallelDel8(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(8)

	b.StartTimer()
	ParallelDel(b, i, h)
}

func BenchmarkParallelDel64(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(64)

	b.StartTimer()
	ParallelDel(b, i, h)
}

func BenchmarkParallelDel128(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(128)

	b.StartTimer()
	ParallelDel(b, i, h)
}

func BenchmarkParallelDel1024(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap()

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(1024)

	b.StartTimer()
	ParallelDel(b, i, h)
}
