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
	var r randomstd.Random = pool.New(150, randomstd.AtomicOffsetConstructor)
	for i := 0; i < count; i++ {
		go func() {
			uuidChannel <- lameUUID(r)
		}()
	}

	for i := 0; i < count; i++ {
		log.Println(<-uuidChannel)
	}
	close(uuidChannel)
}

func lameUUID(r randomstd.Random) string {
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
