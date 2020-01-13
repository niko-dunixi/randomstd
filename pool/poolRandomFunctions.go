package pool

import "math/rand"

func (p *pool) Seed(seed int64) {
	panic("this is not supported on the random pool")
}

func (p *pool) Int63() int64 {
	c := make(chan int64, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Int63()
	})
	return <-c
}

func (p *pool) Uint32() uint32 {
	c := make(chan uint32, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Uint32()
	})
	return <-c
}

func (p *pool) Int31() int32 {
	c := make(chan int32, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Int31()
	})
	return <-c
}

func (p *pool) Int() int {
	c := make(chan int, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Int()
	})
	return <-c
}

func (p *pool) Int63n(n int64) int64 {
	c := make(chan int64, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Int63n(n)
	})
	return <-c
}

func (p *pool) Int31n(n int32) int32 {
	c := make(chan int32, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Int31n(n)
	})
	return <-c
}

func (p *pool) Intn(n int) int {
	c := make(chan int, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Intn(n)
	})
	return <-c
}

func (p *pool) Float64() float64 {
	c := make(chan float64, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Float64()
	})
	return <-c
}

func (p *pool) Float32() float32 {
	c := make(chan float32, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Float32()
	})
	return <-c
}

func (p *pool) Perm(n int) []int {
	c := make(chan []int, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Perm(n)
	})
	return <-c
}

func (p *pool) Shuffle(n int, swap func(i, j int)) {
	p.Work(func(rand *rand.Rand) {
		rand.Shuffle(n, swap)
	})
}

type readResult struct {
	n   int
	err error
}

func (p *pool) Read(pBytes []byte) (n int, err error) {
	c := make(chan readResult, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		rN, rE := rand.Read(pBytes)
		c <- readResult{n: rN, err: rE}
	})
	result := <-c
	return result.n, result.err
}
