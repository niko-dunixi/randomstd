package pool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RandomPoolJob func(rand *rand.Rand)

type RandomPool interface {
	Work(job RandomPoolJob)
}

type pool struct {
	workers chan *rand.Rand
}

type RandomConstructor func() *rand.Rand

var (
	rootSeed       = time.Now().UnixNano()
	offset   int64 = 0
	mutex          = sync.Mutex{}
)

// The most basic creation of a random that utilizes UnixNano
func NaiveRandomConstructor() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// If we call UnixNano fast enough, we could potentially get the same seed.
// This constructor will call UnixNano once, and then offset it by one for
// every call thereafter creating divergent randomness.
func AtomicOffsetRandomConstructor() *rand.Rand {
	mutex.Lock()
	seed := rootSeed + offset
	offset += 1
	mutex.Unlock()

	return rand.New(rand.NewSource(seed))
}

func New(size int, rc RandomConstructor) RandomPool {
	if size <= 0 {
		err := fmt.Errorf("must be a positive initger, but was provided size: %d", size)
		panic(err)
	}

	pool := pool{
		workers: make(chan *rand.Rand, size),
	}
	wg := sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			pool.workers <- rc()
		}()
	}
	wg.Wait()
	return pool
}

func (p pool) Work(job RandomPoolJob) {
	r := <-p.workers
	defer func() {
		p.workers <- r
	}()
	job(r)
}
