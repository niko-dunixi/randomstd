package randpool

import (
	"github.com/paul-nelson-baker/randomstd"
)

func (p *pool) ExpFloat64() float64 {
	c := make(chan float64, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.ExpFloat64()
	})
	return <-c
}

func (p *pool) Float32() float32 {
	c := make(chan float32, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Float32()
	})
	return <-c
}

func (p *pool) Float64() float64 {
	c := make(chan float64, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Float64()
	})
	return <-c
}

func (p *pool) Int() int {
	c := make(chan int, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Int()
	})
	return <-c
}

func (p *pool) Int31() int32 {
	c := make(chan int32, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Int31()
	})
	return <-c
}

func (p *pool) Int31n(n int32) int32 {
	c := make(chan int32, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Int31n(n)
	})
	return <-c
}

func (p *pool) Int63() int64 {
	c := make(chan int64, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Int63()
	})
	return <-c
}

func (p *pool) Int63n(n int64) int64 {
	c := make(chan int64, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Int63n(n)
	})
	return <-c
}

func (p *pool) Intn(n int) int {
	c := make(chan int, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Intn(n)
	})
	return <-c
}

func (p *pool) NormFloat64() float64 {
	c := make(chan float64, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.NormFloat64()
	})
	return <-c
}

func (p *pool) Perm(n int) []int {
	c := make(chan []int, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Perm(n)
	})
	return <-c
}

func (p *pool) Read(pBytes []byte) (n int, err error) {
	c := make(chan readResult, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		rN, rE := r.Read(pBytes)
		c <- readResult{n: rN, err: rE}
	})
	result := <-c
	return result.n, result.err
}

func (p *pool) Seed(seed int64) {
	panic("this is not supported on the random pool")
}

func (p *pool) Shuffle(n int, swap func(i, j int)) {
	p.Work(func(r randomstd.Random) {
		r.Shuffle(n, swap)
	})
}

func (p *pool) Uint32() uint32 {
	c := make(chan uint32, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- r.Uint32()
	})
	return <-c
}

type readResult struct {
	n   int
	err error
}
