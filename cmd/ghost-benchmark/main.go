// Benchmark for Ghost server.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	embedded bool
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
	flag.BoolVar(&embedded, "embedded", false, "Test embedded storage")
	flag.StringVar(&host, "host", "localhost", "Server hostname")
	flag.IntVar(&port, "port", 6869, "Server port")
	flag.IntVar(&clients, "clients", 50, "Number of paralel connections")
	flag.IntVar(&requests, "requests", 10000, "Total number of requests")
	flag.IntVar(&size, "size", 2, "Data size of SET/GET value in bytes")
	flag.IntVar(&keyrange, "keyrange", 100, "Use random keys for SET/GET")
	flag.Parse()

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
func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := make([]byte, n)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
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
	if clients >= requests {
		log.Fatal("Clients should be < requests")
	}

	if embedded {
		c := obtainCollection()
		printResults("SET", benchmarkEmbeddedSet(c))
		printResults("GET", benchmarkEmbeddedGet(c))
		printResults("DEL", benchmarkEmbeddedDel(c))
	} else {
		c := obtainClient()
		printResults("SET", benchmarkServerSet(c))
		printResults("GET", benchmarkServerGet(c))
		printResults("DEL", benchmarkServerDel(c))
	}
}
