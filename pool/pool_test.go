package pool

import (
	"math/rand"
	"testing"
)

func TestNewInvalidSizeZero(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("this should have failed, but did not")
		}
	}()
	New(0, NaiveRandomConstructor)
}

func TestNewInvalidSizeNegative(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("this should have failed, but did not")
		}
	}()
	New(-1, NaiveRandomConstructor)
}

func TestNaiveRandomConstructor(t *testing.T) {
	count := 100
	intsChannel := make(chan int, count)
	randomPool := New(100, NaiveRandomConstructor)
	for i := 0; i < count; i++ {
		randomPool.Work(func(rand *rand.Rand) {
			intsChannel <- rand.Int()
		})
	}
	for i := 0; i < count; i++ {
		<-intsChannel
	}
}

func TestAtomicOffsetRandomConstructor(t *testing.T) {
	count := 100
	intsChannel := make(chan int, count)
	randomPool := New(100, AtomicOffsetRandomConstructor)
	for i := 0; i < count; i++ {
		randomPool.Work(func(rand *rand.Rand) {
			intsChannel <- rand.Int()
		})
	}
	for i := 0; i < count; i++ {
		<-intsChannel
	}
}
