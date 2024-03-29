# randomstd
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
type RandomPoolJob func(r randomstd.Random)
```

A potentially time-expensive example is the need to generate
UUIDs randomly in a time efficient manor like this example
demonstrates:

```go
package main

import (
	"github.com/paul-nelson-baker/randomstd"
	"github.com/paul-nelson-baker/randomstd/pool"
	"log"
	"strconv"
	"strings"
)

func main() {
	count := 100_000
	uuidChannel := make(chan string)

	// The magic happens here, we kick off our `n` goroutines
	// and not worry about managing the random itself. We just
	// call random's methods like normal and let the concurrency
	// magic happen on its own!
	var r randomstd.Random = pool.New(150, pool.AtomicOffsetRandomConstructor)
	for i := 0; i < count; i++ {
		go func() {
			uuidChannel <- naiveUUID(r)
		}()
	}

	for i := 0; i < count; i++ {
		log.Println(<-uuidChannel)
	}
	close(uuidChannel)
}

func naiveUUID(r randomstd.Random) string {
	b := strings.Builder{}
	for i := 0; i < 8; i++ {
		b.WriteString(value(r))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(r))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(r))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(r))
	}
	b.WriteString("-")
	for i := 0; i < 12; i++ {
		b.WriteString(value(r))
	}
	return b.String()
}

func value(r randomstd.Random) string {
	i := r.Int63n(16)
	return strconv.FormatInt(i, 16)
}

```
