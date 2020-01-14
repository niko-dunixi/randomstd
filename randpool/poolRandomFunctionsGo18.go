// +build go1.8

package randpool

import (
	"github.com/paul-nelson-baker/randomstd"
	"math/rand"
)

func (p *pool) Uint64() uint64 {
	c := make(chan uint64, 1)
	defer close(c)
	p.Work(func(r randomstd.Random) {
		c <- rand.Uint64()
	})
	return <-c
}
