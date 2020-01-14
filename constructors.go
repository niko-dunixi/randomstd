package randomstd

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type Constructor func() Random

var (
	rootSeed       = time.Now().UnixNano()
	offset   int64 = 0
)

// If we call UnixNano fast enough, we could potentially get the same seed.
// This constructor will call UnixNano once, and then offset it by one for
// every call thereafter creating divergent randomness.
func AtomicOffsetConstructor() Random {
	seed := rootSeed + offset
	atomic.AddInt64(&offset, 1)
	return NaiveSeededConstructor(seed)
}

// The most basic creation of a random that utilizes UnixNano
func NaiveConstructor() Random {
	return NaiveSeededConstructor(time.Now().UnixNano())
}

func NaiveSeededConstructor(seed int64) Random {
	return rand.New(rand.NewSource(seed))
}
