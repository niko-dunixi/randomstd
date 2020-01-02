package main

import (
	"bytes"
	"github.com/paul-nelson-baker/concurrent-random/pool"
	"github.com/paul-nelson-baker/concurrent-random/safe"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	steps := 100
	baseSize := 1000
	stepSize := 1000
	for i := 0; i < steps; i++ {
		pooledConcurrentRandom := createPooledConcurrentRandom(2)
		b := benchmark(steps, baseSize+(i*stepSize), pooledConcurrentRandom)
		log.Printf("Pooled: %+v", b)
		b = benchmark(steps, baseSize+(i*stepSize), simpleConcurrentRandom)
		log.Printf("Simple: %+v", b)
	}
}

func benchmark(count, size int, action func(size int) int64) []int64 {
	var times = make([]int64, 0, count)
	for i := 0; i < count; i++ {
		millis := action(size)
		times = append(times, millis)
	}
	return times
}

func simpleConcurrentRandom(size int) int64 {
	random := rand.New(singleton.NewSource(time.Now().UnixNano()))
	wg := sync.WaitGroup{}
	wg.Add(size)
	start := time.Now()
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			lameUUID(random)
		}()
	}
	wg.Wait()
	end := time.Now()
	return end.Sub(start).Milliseconds()
}

func createPooledConcurrentRandom(poolSize int) func(size int) int64 {
	randomPool := pool.New(poolSize, pool.NaiveRandomConstructor)
	a := func(size int) int64 {
		wg := sync.WaitGroup{}
		wg.Add(size)
		start := time.Now()
		for i := 0; i < size; i++ {
			go func() {
				defer wg.Done()
				randomPool.Work(func(rand *rand.Rand) {
					lameUUID(rand)
				})
			}()
		}
		wg.Wait()
		end := time.Now()
		return end.Sub(start).Milliseconds()
	}
	return a
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
