package randpool

import (
	"fmt"
	"github.com/paul-nelson-baker/randomstd"
	"sync"
)

type RandomPool interface {
	randomstd.Random
	Work(task randomstd.RandomConsumer)
}

type pool struct {
	workers chan randomstd.Random
}

func New(size int, rc randomstd.Constructor) RandomPool {
	if size <= 0 {
		err := fmt.Errorf("must be a positive initger, but was provided size: %d", size)
		panic(err)
	}

	pool := pool{
		workers: make(chan randomstd.Random, size),
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
	return &pool
}

func (p pool) Work(task randomstd.RandomConsumer) {
	r := <-p.workers
	defer func() {
		p.workers <- r
	}()
	task(r)
}
