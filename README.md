# Concurrent Random
## Abstract
GoLang has first class concurrency, which is AWESOME!

A big problem that catches many beginners and new
users is the use of random. The philosophy of the
core language team is to NOT make things safe. The
reason for this is to allow application developers
to make choices about how to implement the safety
for their given use-case.

If you want to use Random, you have to create a
concurrent safe [`Source`](https://golang.org/src/math/rand/rand.go)
which is not difficult to do, but there is an on-going
discussion on whether to make this part of the 
core library: https://github.com/golang/go/issues/21393
(I've included this in the `safe` package for simplicity
but this is not what this library is for).

This is an implementation that takes advantage of first
class function pointers and concurrency and creates
a safe pool that allows end-developers to use random
concurrently without having to think about concurrency

## Usage

You can use it however you want. All you have to do is
pass in a function pointer that matches the following
signature, consuming a random pointer and returning void:

```go
type RandomPoolJob func(rand *rand.Rand)
```

A potentially time-expensive example is the need to generate
UUIDs randomly in a time efficient manor like this example
demonstrates:

```go
package main

import (
	"bytes"
	"github.com/paul-nelson-baker/concurrent-random/pool"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

func main() {
	count := 100_000
	wg := sync.WaitGroup{}
	wg.Add(count)
	uuidChannel := make(chan string)

	// The magic happens here, we kick off our `n` goroutines
	// and not worry about managing the random itself. We just
	// borrow it and let it return to the pool afterwards!
	randomPool := pool.New(150, pool.AtomicOffsetRandomConstructor)
	for i := 0; i < count; i++ {
		go randomPool.Work(func(rand *rand.Rand) {
			defer wg.Done()
			uuidChannel <- lameUUID(rand)
		})
	}

	go func() {
		for i := 0; i < count; i++ {
			log.Println(<-uuidChannel)
		}
	}()

	wg.Wait()
	close(uuidChannel)
}

func lameUUID(random *rand.Rand) string {
	b := bytes.Buffer{}
	for i := 0; i < 8; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 12; i++ {
		b.WriteString(value(random))
	}
	return b.String()
}

func value(random *rand.Rand) string {
	i := random.Int63n(16)
	return strconv.FormatInt(i, 16)
}
```