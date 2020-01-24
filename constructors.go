package randomstd

import (
	"math/rand"
	"sync"
	"time"
)

type Constructor func() Random
type RandomSupplier Constructor
type RandomConsumer func(r Random)

var (
	rootSeed        = time.Now().UnixNano()
	offset    int64 = 0
	mutexSeed       = sync.Mutex{}
)

// If we call UnixNano fast enough, we could potentially get the same seed.
// This constructor will call UnixNano once, and then offset it by one for
// every call thereafter creating divergent randomness.
func AtomicOffsetConstructor() Random {
	mutexSeed.Lock()
	seed := rootSeed + offset
	offset += 1
	mutexSeed.Unlock()
	return NaiveSeededConstructor(seed)
}

// The most basic creation of a random that utilizes UnixNano
func NaiveConstructor() Random {
	return NaiveSeededConstructor(time.Now().UnixNano())
}

func NaiveSeededConstructor(seed int64) Random {
	return rand.New(rand.NewSource(seed))
}
