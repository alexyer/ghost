package ghost

import "testing"

func BenchmarkSet(b *testing.B) {
	b.StopTimer()
	h := NewHashMap(1024 * 1024 * 10)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		h.Set(string(i), "Yarrr")
	}
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	h := NewHashMap(1024 * 1024 * 10)
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
	h := NewHashMap(1024 * 1024 * 10)
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
	h := NewHashMap(1024 * 1024 * 10)
	i := 0
	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet8(b *testing.B) {
	b.StopTimer()
	h := NewHashMap(1024 * 1024 * 10)
	i := 0

	b.SetParallelism(8)

	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet64(b *testing.B) {
	b.StopTimer()
	h := NewHashMap(1024 * 1024 * 10)
	i := 0

	b.SetParallelism(64)

	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet128(b *testing.B) {
	b.StopTimer()
	h := NewHashMap(1024 * 1024 * 10)
	i := 0

	b.SetParallelism(128)

	b.StartTimer()

	ParallelSet(b, i, h)
}

func BenchmarkParallelSet1024(b *testing.B) {
	b.StopTimer()
	h := NewHashMap(1024 * 1024 * 10)
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
	h := NewHashMap(1024 * 1024 * 10)

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.StartTimer()
	ParallelGet(b, i, h)
}

func BenchmarkParallelGet8(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap(1024 * 1024 * 10)

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
	h := NewHashMap(1024 * 1024 * 10)

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
	h := NewHashMap(1024 * 1024 * 10)

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
	h := NewHashMap(1024 * 1024 * 10)

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
	h := NewHashMap(1024 * 1024 * 10)

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.StartTimer()
	ParallelDel(b, i, h)
}

func BenchmarkParallelDel8(b *testing.B) {
	b.StopTimer()
	i := 0
	h := NewHashMap(1024 * 1024 * 10)

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
	h := NewHashMap(1024 * 1024 * 10)

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
	h := NewHashMap(1024 * 1024 * 10)

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
	h := NewHashMap(1024 * 1024 * 10)

	for i := 0; i < 5000; i++ {
		h.Set(string(i), "Yarrr")
	}

	b.SetParallelism(1024)

	b.StartTimer()
	ParallelDel(b, i, h)
}
