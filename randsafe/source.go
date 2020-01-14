package randsafe

// Almost all this code is stolen from the Go rand package.
// Their concurrent code is package private because it is
// used with their global random. This blocks us from using
// it with:
//  - single random instances
//  - outside of their package
// Our copy allows us to create multiple individual instances
// that we can manage across different goroutines.

import (
	"math/rand"
	"sync"
)

func NewSource(seed int64) rand.Source {
	source := rand.NewSource(seed)
	return &lockedSource{
		lk:  sync.Mutex{},
		src: source.(rand.Source64),
	}
}

type lockedSource struct {
	lk  sync.Mutex
	src rand.Source64
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src.Uint64()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

// seedPos implements Seed for a lockedSource without a race condition.
func (r *lockedSource) seedPos(seed int64, readPos *int8) {
	r.lk.Lock()
	r.src.Seed(seed)
	*readPos = 0
	r.lk.Unlock()
}

// read implements Read for a lockedSource without a race condition.
func (r *lockedSource) read(p []byte, readVal *int64, readPos *int8) (n int, err error) {
	r.lk.Lock()
	n, err = read(p, r.src.Int63, readVal, readPos)
	r.lk.Unlock()
	return
}

func read(p []byte, int63 func() int64, readVal *int64, readPos *int8) (n int, err error) {
	pos := *readPos
	val := *readVal
	for n = 0; n < len(p); n++ {
		if pos == 0 {
			val = int63()
			pos = 7
		}
		p[n] = byte(val)
		val >>= 8
		pos--
	}
	*readPos = pos
	*readVal = val
	return
}
