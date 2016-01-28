[![Build Status](https://travis-ci.org/alexyer/ghost.svg?branch=master)](https://travis-ci.org/alexyer/ghost)
[![Coverage Status](https://coveralls.io/repos/alexyer/ghost/badge.svg?branch=master&service=github)](https://coveralls.io/github/alexyer/ghost?branch=master)
[![GoDoc](https://godoc.org/github.com/alexyer/ghost?status.svg)](https://godoc.org/github.com/alexyer/ghost)

# Ghost
Yet another in-memory key/value storage written in Go.

### Description
Simple key/value storage.
Uses implementation of lock-free hashmap based on
"Split-Ordered Lists - Lock-free Resizable Hash Tables" by Shalev & Shavit and
"Lock-free Dynamically Resizable Arrays" by Dechev, Pirkelbauer, Stroustrup works
to provide good performance and concurrency-safe access.

### Features
  * Concurrency safe
  * Fast
  * Written in pure Go, means could be used in any place where Go could be run
  * Could be used as embedded storage
  * Could be run as standalone server

## Commands

Server commands:
  * PING -- Test command. Returns "Pong!".

Hash commands:
  * SET &lt;key&gt; &lt;value&gt; -- Set create or update &lt;key&gt; with &lt;value&gt;.
  * GET &lt;key&gt; -- Get value of the &lt;key&gt;.
  * DEL &lt;key&gt; -- Delete key &lt;key&gt;.

Collection commands:
  * CGET &lt;collection name&gt; -- Change user's collection.
  * CADD &lt;collection name&gt; -- Create new collection.

## Server

####Build server:
```sh
make ghost-server
```

####Run server:
```sh
ghost -host localhost -port 6869
```
### Benchmark
####Build:
```sh
make ghost-benchmark
```

```sh
Usage of ./ghost-benchmark:
  -clients int
        Number of paralel connections (default 50)
  -host string
        Server hostname (default "localhost")
  -keyrange int
        Use random keys for SET/GET (default 100)
  -port int
        Server port (default 6869)
  -requests int
        Total number of requests (default 10000)
  -size int
        Data size of SET/GET value in bytes (default 2)
  -socket string
        listen to unix socket
```

## Client
```go
package main

import (
	"fmt"

	"github.com/alexyer/ghost/client"
)

func main() {
    // Create new client and connect to the Ghost server.
	ghost := client.New(&client.Options{
		Addr: "localhost:6869",
	})

	ghost.Set("key1", "val2")      // Set key
	res, err := ghost.Get("key1")  // Get key
	ghost.Del("key1")              // Del key

	ghost.CAdd("new-collection")  // Create new collection
	ghost.CGet("new-collection")  // Change client collection
}
```

## CLI
Now you can use a simple cli to test or play with data. All the current commands
are supported. Cli works only if ghost-server exists on provided address.

####Build:
```sh
make ghost-cli
```

####Run cli:
```sh
ghost-cli -host localhost -port 6869
```

#### Example session:
```sh
> ping # will test the connection
Pong!
> set hello world # will set value "world" to key "hello"
OK
> get hello # will get the value stored with key "hello"
world
> del hello # will delete the value stored with key "hello"
OK
> cadd mars # will add new "mars" collection
OK
> cget mars # will select "mars" collection
OK
> set "few words key" "few words value" # if few words in value or keys is needed surround it with quotes
OK
> get "few words key"
few words value
> set song "riders on the storm" # only one argument could be in quotes if needed
OK
> get song
riders on the storm
> set "another song" stairway # other order is possible
OK
> get "another song"
stairway
```

## Embedded
### Benchmark
Ghost hashmap

```
BenchmarkSet-4                 3000000       383 ns/op
BenchmarkGet-4                10000000       114 ns/op
BenchmarkDel-4                10000000       106 ns/op
```

Ghost concurrent hashmap

```
BenchmarkParallelSet-4       5000000           382 ns/op
BenchmarkParallelSet8-4      2000000           530 ns/op
BenchmarkParallelSet64-4     3000000           392 ns/op
BenchmarkParallelSet128-4    5000000           312 ns/op
BenchmarkParallelSet1024-4   1000000         33503 ns/op

BenchmarkParallelGet-4      20000000            75.7 ns/op
BenchmarkParallelGet8-4     20000000            58.9 ns/op
BenchmarkParallelGet64-4    20000000            62.8 ns/op
BenchmarkParallelGet128-4   20000000            61.1 ns/op
BenchmarkParallelGet1024-4  20000000            63.1 ns/op

BenchmarkParallelDel-4      30000000            46.2 ns/op
BenchmarkParallelDel8-4     30000000            46.1 ns/op
BenchmarkParallelDel64-4    30000000            45.7 ns/op
BenchmarkParallelDel128-4   30000000            45.5 ns/op
BenchmarkParallelDel1024-4  30000000            45.6 ns/op
```

Native hashmap

```
BenchmarkNativeSet-4           5000000       220 ns/op
BenchmarkNativeGet-4          30000000       41.7 ns/op
BenchmarkNativeDel-4          100000000      15.7 ns/op
```

### Example
```go
package main

import (
        "fmt"

        "github.com/alexyer/ghost"
)

func main() {
        //Storage
        storage := ghost.GetStorage() // Get storage instance

        storage.AddCollection("newcollection")          // Create new collection
        mainCollection := storage.GetCollection("main") // Get existing collection
        storage.DelCollection("newcollection")          // Delete collection

        // Collections
        mainCollection.Set("somekey", "42") // Set item

        val, _ := mainCollection.Get("somekey") // Get item from Collection
        fmt.Println(val)

        mainCollection.Del("somekey") // Delete item
}
```

## TODO
  * Improve CLI
  * Improve server and get rid of limitations
  * Improve documentation
  * Properly comment sources

## Contributing
It's learing project, so there are possible a lot of issues, espesially in concurrent code,
so any improvements, critics or propsals are highly appretiated.
