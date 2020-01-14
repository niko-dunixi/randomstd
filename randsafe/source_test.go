package randsafe

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewSource(t *testing.T) {
	count := 1_000_000
	wg := sync.WaitGroup{}
	wg.Add(count)
	defer wg.Wait()

	source := NewSource(time.Now().UnixNano())
	random := rand.New(source)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			random.Int()
		}()
	}
}
