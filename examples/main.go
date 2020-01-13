package main

import (
	"bytes"
	"github.com/paul-nelson-baker/randomstd/pool"
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
