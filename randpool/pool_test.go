package randpool

import (
	"github.com/paul-nelson-baker/randomstd"
	"testing"
	"time"
)

func TestNewInvalidSizeZero(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("this should have failed, but did not")
		}
	}()
	New(0, randomstd.NaiveConstructor)
}

func TestNewInvalidSizeNegative(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("this should have failed, but did not")
		}
	}()
	New(-1, randomstd.NaiveConstructor)
}

func TestNaiveRandomConstructor(t *testing.T) {
	count := 100
	intsChannel := make(chan int, count)
	randomPool := New(100, randomstd.NaiveConstructor)
	for i := 0; i < count; i++ {
		randomPool.Work(func(r randomstd.Random) {
			intsChannel <- r.Int()
		})
	}
	for i := 0; i < count; i++ {
		<-intsChannel
	}
}

func TestAtomicOffsetRandomConstructor(t *testing.T) {
	count := 100
	valuesChannel := make(chan int, count)
	randomPool := New(100, randomstd.AtomicOffsetConstructor)
	for i := 0; i < count; i++ {
		randomPool.Work(func(r randomstd.Random) {
			valuesChannel <- r.Int()
		})
	}
	for i := 0; i < count; i++ {
		<-valuesChannel
	}
}

func TestInterface(t *testing.T) {
	seedValue := time.Now().UnixNano()

	var randomPool randomstd.Random = New(1, func() randomstd.Random {
		return randomstd.NaiveSeededConstructor(seedValue)
	})
	var randomSingle = randomstd.NaiveSeededConstructor(seedValue)

	for i := 0; i < 1000; i++ {
		a := randomPool.Int()
		b := randomSingle.Int()
		if a != b {
			t.Errorf("Values diverged, but should not have.")
		}
	}
}
