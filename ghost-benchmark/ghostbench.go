// Benchmark for Ghost server.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/alexyer/ghost/client"
)

var (
	host     string
	port     int
	clients  int
	requests int
	size     int
	keyrange int
)

type result struct {
	totTime time.Duration
	reqSec  float64
}

func init() {
	flag.StringVar(&host, "host", "localhost", "Server hostname")
	flag.IntVar(&port, "port", 6869, "Server port")
	flag.IntVar(&clients, "clients", 50, "Number of paralel connections")
	flag.IntVar(&requests, "requests", 10000, "Total number of requests")
	flag.IntVar(&size, "size", 2, "Data size of SET/GET value in bytes")
	flag.IntVar(&keyrange, "keyrange", 100, "Use random keys for SET/GET")
	flag.Parse()

}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := make([]byte, n)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

func obtainClient() *client.GhostClient {
	return client.New(&client.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		PoolSize: clients,
	})
}

func initTestData(prefix string, n, size, keyrange int) ([]string, []string) {
	keys := make([]string, n)
	vals := make([]string, n)

	for i := 0; i < n; i++ {
		keys[i] = fmt.Sprintf("%s_key:%d", prefix, rand.Intn(keyrange))
		vals[i] = randString(size)
	}
	return keys, vals
}

func benckmarkSet(c *client.GhostClient) result {
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

func benckmarkGet(c *client.GhostClient) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("get", requests, size, keyrange)
	populateTestData(c, keys, vals)

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

func benckmarkDel(c *client.GhostClient) result {
	var wg sync.WaitGroup
	keys, vals := initTestData("get", requests, size, keyrange)
	populateTestData(c, keys, vals)

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

func populateTestData(c *client.GhostClient, keys, vals []string) {
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

func printResults(name string, res result) {
	fmt.Printf("========= %s =========\n", name)
	fmt.Printf("Total time: %v\n", res.totTime)
	fmt.Printf("Requests completed: %d\n", requests)
	fmt.Printf("Requests per second: %.2f\n", res.reqSec)
	fmt.Printf("Parallel clients: %d\n", clients)
	fmt.Printf("Payload: %d\n", size)
}

func main() {
	c := obtainClient()

	printResults("SET", benckmarkSet(c))
	printResults("GET", benckmarkGet(c))
	printResults("DEL", benckmarkDel(c))
}
