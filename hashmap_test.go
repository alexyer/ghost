package ghost

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const MAXSIZE = 82

var (
	empty     *hashMap
	one       *hashMap
	several   *hashMap
	many      *hashMap
	wordsHash *hashMap
	nativeMap map[string]string
	words     []string
	result    string
)

func init() {
	empty = NewHashMap()
	one = NewHashMap()
	several = NewHashMap()
	many = NewHashMap()
	wordsHash = NewHashMap()

	one.Set("One", "one")

	several.Set("One", "one")
	several.Set("Two", "two")

	for i := 0; i < MAXSIZE; i++ {
		many.Set(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
	}

	raw, err := ioutil.ReadFile("/usr/share/dict/cracklib-small")

	if err == nil {
		data := string(raw)
		words = strings.Split(data, "\n")
	}

	for _, w := range words {
		wordsHash.Set(string(w), string(w))
	}

	nativeMap = make(map[string]string)

	for _, w := range words {
		nativeMap[string(w)] = string(w)
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

	if _, err := one.Get("One"); err == nil {
		t.Errorf("Wrong delete behavior")
	}

	several.Del("Two")

	if _, err := several.Get("Two"); err == nil {
		t.Errorf("Wrong delete behavior")
	}
}

func BenchmarkSet(b *testing.B) {
	h := NewHashMap()

	for i := 0; i < b.N; i++ {
		for _, w := range words {
			h.Set(string(w), string(w))
		}
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, w := range words {
			wordsHash.Get(string(w))
		}
	}
}

func BenchmarkDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, w := range words {
			wordsHash.Del(string(w))
		}
	}
}

func BenchmarkMixed(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			op := r.Intn(100)

			switch {
			case op <= 75:
				wordsHash.Get(words[r.Intn(2500)])
			case op <= 90:
				wordsHash.Set(words[r.Intn(2500)], words[r.Intn(2500)])
			default:
				wordsHash.Del(words[r.Intn(2500)])
			}
		}
	}
}

func BenchmarkNativeSet(b *testing.B) {
	h := make(map[string]string)

	for i := 0; i < b.N; i++ {
		for _, w := range words {
			h[string(w)] = string(w)
		}
	}
}

func BenchmarkNativeGet(b *testing.B) {
	var r string

	for i := 0; i < b.N; i++ {
		for _, w := range words {
			r = nativeMap[string(w)]
		}
	}

	result = r
}

func BenchmarkNativeDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, w := range words {
			delete(nativeMap, string(w))
		}
	}
}

func BenchmarkNativeMixed(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var res string

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			op := r.Intn(100)

			switch {
			case op < 75:
				res = nativeMap[words[r.Intn(2500)]]
			case op < 90:
				nativeMap[words[r.Intn(2500)]] = words[r.Intn(2500)]
			default:
				delete(nativeMap, words[r.Intn(2500)])
			}
		}
	}

	result = res
}
