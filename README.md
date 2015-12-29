[![Build Status](https://travis-ci.org/alexyer/ghost.svg?branch=master)](https://travis-ci.org/alexyer/ghost)
[![Coverage Status](https://coveralls.io/repos/alexyer/ghost/badge.svg?branch=master&service=github)](https://coveralls.io/github/alexyer/ghost?branch=master)
[![GoDoc](https://godoc.org/github.com/alexyer/ghost?status.svg)](https://godoc.org/github.com/alexyer/ghost)

# Ghost
Yet another in-memory key/value storage written in Go.

### Description
Simple key/value storage based on implementation of striped hashmap data structure.
Yes, it has terrible performance, the point was to make something simple to get more comfortabe with Go language.
I hope to improve it one day...or not.

### Features
  * Concurrency safe
  * Slow
  * ACID - seriously, how do you think? ;)
  * Written in pure Go
  * Could be used as embedded storage
  * Could be run as standalone server

## Embedded

### Benchmark
Ghost hashmap

```
BenchmarkSet-2                    100000             10775 ns/op
BenchmarkGet-2                   5000000               277 ns/op
BenchmarkDel-2                   5000000               226 ns/op
```

Ghost concurrent hashmap

```
BenchmarkParallelSet-2            200000              9608 ns/op
BenchmarkParallelSet8-2           100000             11740 ns/op
BenchmarkParallelSet64-2          100000             11315 ns/op
BenchmarkParallelSet128-2         100000             11140 ns/op
BenchmarkParallelSet1024-2        100000             12608 ns/op

BenchmarkParallelGet-2           5000000               307 ns/op
BenchmarkParallelGet8-2          5000000               320 ns/op
BenchmarkParallelGet64-2         5000000               321 ns/op
BenchmarkParallelGet128-2        5000000               321 ns/op
BenchmarkParallelGet1024-2       3000000               359 ns/op

BenchmarkParallelDel-2          10000000               195 ns/op
BenchmarkParallelDel8-2         10000000               201 ns/op
BenchmarkParallelDel64-2        10000000               207 ns/op
BenchmarkParallelDel128-2       10000000               201 ns/op
BenchmarkParallelDel1024-2      10000000               212 ns/op
```

Native hashmap

```
BenchmarkNativeSet-2             1000000              2137 ns/op
BenchmarkNativeGet-2            10000000               116 ns/op
BenchmarkNativeDel-2            30000000                43.7 ns/op
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

## Server
Server is under development. The main limitation - server does not accept messages more that 4KB.
Will be fixed in future versions.

Run server:
```sh
ghost -host localhost -port 6869
```

### Commands

Hash commands:
  * PING -- Test command. Returns "Pong!".
  * SET &lt;key&gt; &lt;value&gt; -- Set create or update &lt;key&gt; with &lt;value&gt;.
  * GET &lt;key&gt; -- Get value of the &lt;key&gt;.
  * DEL &lt;key&gt; -- Delete key &lt;key&gt;.

Collection commands:
  * CGET <collection name> -- Change user's collection.
  * CADD <collection name> -- Create new collection.

### Client
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

## TODO
  * Implement CLI
  * Improve storage performance
  * Improve server and get rid of limitations

## Contributing
It's learing project, so there are possible a lot of issues, espesially in concurrent code,
so any improvements, critics or propsals are highly appretiated.
