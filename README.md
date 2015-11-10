[![Build Status](https://travis-ci.org/alexyer/ghost.svg?branch=master)](https://travis-ci.org/alexyer/ghost)
[![Coverage Status](https://coveralls.io/repos/alexyer/ghost/badge.svg?branch=master&service=github)](https://coveralls.io/github/alexyer/ghost?branch=master)

## Ghost
Yet another in-memory key/value storage written in Go.

## Description
Simple key/value storage based on naive implementation of hashmap data structure.
Yes, it has terrible performance, the point was to make something simple to get more comfortabe with Go language.
I hope to improve it one day...or not.

## Features
  * Concurrency unsafe
  * Slow
  * ACID - seriously, how do you think? ;)

## Benchmark
Ghost hashmap

```
BenchmarkSet-2                30          38032398 ns/op
BenchmarkGet-2                50          29517307 ns/op
BenchmarkDel-2               100          19910142 ns/op
BenchmarkMixed-2             300           5886845 ns/op
```

Native hashmap

```
BenchmarkNativeSet-2         100          14275454 ns/op
BenchmarkNativeGet-2         200           9778833 ns/op
BenchmarkNativeDel-2        3000            526531 ns/op
BenchmarkNativeMixed-2       500           2660463 ns/op
```

## Example
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
