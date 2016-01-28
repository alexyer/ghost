package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/alexyer/ghost/client"
)

func obtainClient() *client.GhostClient {
	return newClient(fmt.Sprintf("%s:%d", host, port), "tcp")
}

func obtainUnixSocketClient() *client.GhostClient {
	return newClient(fmt.Sprintf(socket), "unix")
}

func newClient(addr, network string) *client.GhostClient {
	return client.New(&client.Options{
		Addr:        addr,
		Network:     network,
		PoolSize:    clients,
		PoolTimeout: time.Duration(poolTimeout) * time.Second,
	})
}

func benchmarkServerSet(c *client.GhostClient) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("set", requests, size, keyrange)

	start := time.Now()

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(keys, vals []string, i int) {
			c.Set(keys[i], vals[i])
			wg.Done()
		}(keys, vals, i)
	}
	wg.Wait()

	latency := time.Since(start)

	return result{
		totTime: latency,
		reqSec:  float64(requests) / latency.Seconds(),
	}
}

func benchmarkServerGet(c *client.GhostClient) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("get", requests, size, keyrange)
	populateTestDataServer(c, keys, vals)

	start := time.Now()
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(i int) {
			c.Get(keys[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	latency := time.Since(start)

	return result{
		totTime: latency,
		reqSec:  float64(requests) / latency.Seconds(),
	}
}

func benchmarkServerDel(c *client.GhostClient) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("get", requests, size, keyrange)
	populateTestDataServer(c, keys, vals)

	start := time.Now()
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(i int) {
			c.Del(keys[i])
			wg.Done()
		}(i)
	}
	wg.Wait()

	latency := time.Since(start)

	return result{
		totTime: latency,
		reqSec:  float64(requests) / latency.Seconds(),
	}
}

func populateTestDataServer(c *client.GhostClient, keys, vals []string) {
	var wg sync.WaitGroup
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(i int) {
			c.Set(keys[i], vals[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
