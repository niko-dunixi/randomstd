package main

import (
	"bytes"
	"github.com/paul-nelson-baker/randomstd"
	"github.com/paul-nelson-baker/randomstd/pool"
	"log"
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
		go randomPool.Work(func(r randomstd.Random) {
			defer wg.Done()
			uuidChannel <- lameUUID(r)
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

func lameUUID(r randomstd.Random) string {
	b := bytes.Buffer{}
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
