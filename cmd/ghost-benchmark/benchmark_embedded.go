package main

import (
	"sync"
	"time"

	"github.com/alexyer/ghost/ghost"
)

func obtainCollection() *ghost.Collection {
	return ghost.GetStorage().GetCollection("main")
}

func benchmarkEmbeddedSet(gh *ghost.Collection) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("set", requests, size, keyrange)

	start := time.Now()

	for i := requests; i >= 0; i -= clients {
		for j := 0; j < clients; j++ {
			wg.Add(1)
			go func(i int) {
				gh.Set(keys[j], vals[j])
				wg.Done()
			}(i)
		}
		wg.Wait()
	}

	latency := time.Since(start)

	return result{
		totTime: latency,
		reqSec:  float64(requests) / latency.Seconds(),
	}
}

func benchmarkEmbeddedGet(gh *ghost.Collection) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("get", requests, size, keyrange)
	populateTestDataEmbedded(gh, keys, vals)

	start := time.Now()
	for i := requests; i >= 0; i -= clients {
		for j := 0; j < clients; j++ {
			wg.Add(1)
			go func(i int) {
				gh.Get(keys[j])
				wg.Done()
			}(i)
		}
		wg.Wait()
	}

	latency := time.Since(start)

	return result{
		totTime: latency,
		reqSec:  float64(requests) / latency.Seconds(),
	}
}

func benchmarkEmbeddedDel(gh *ghost.Collection) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("get", requests, size, keyrange)
	populateTestDataEmbedded(gh, keys, vals)

	start := time.Now()
	for i := requests; i >= 0; i -= clients {
		for j := 0; j < clients; j++ {
			wg.Add(1)
			go func(i int) {
				gh.Del(keys[j])
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
	wg.Wait()

	latency := time.Since(start)

	return result{
		totTime: latency,
		reqSec:  float64(requests) / latency.Seconds(),
	}
}

func populateTestDataEmbedded(gh *ghost.Collection, keys, vals []string) {
	var wg sync.WaitGroup
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(i int) {
			gh.Set(keys[i], vals[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
