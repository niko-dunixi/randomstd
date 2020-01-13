// +build go1.8

package pool

import "math/rand"

func (p *pool) Uint64() uint64 {
	c := make(chan uint64, 1)
	defer close(c)
	p.Work(func(rand *rand.Rand) {
		c <- rand.Uint64()
	})
	return <-c
}
