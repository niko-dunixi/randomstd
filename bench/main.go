package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/paul-nelson-baker/randomstd"
	"github.com/paul-nelson-baker/randomstd/randpool"
	"github.com/paul-nelson-baker/randomstd/randsafe"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	benchmarkResults := map[string][]bench{}

	steps := 30
	baseSize := 1000
	stepSize := 1000

	var simpleBench []bench
	for i := 0; i < steps; i++ {
		b := benchmark(steps, baseSize+(i*stepSize), simpleConcurrentRandom)
		simpleBench = append(simpleBench, b)
	}
	benchmarkResults["simple_concurrency"] = simpleBench
	poolSizes := []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100,}
	for _, poolSize := range poolSizes {
		log.Printf("Pool Size: %d", poolSize)
		pooledConcurrentRandom := createPooledConcurrentRandom(poolSize)
		var pooledBench []bench
		for i := 0; i < steps; i++ {
			b := benchmark(steps, baseSize+(i*stepSize), pooledConcurrentRandom)
			pooledBench = append(pooledBench, b)
		}
		name := fmt.Sprintf("pool-%d-concurrency", poolSize)
		benchmarkResults[name] = pooledBench
	}

	if err := saveBenchmarks("concurrency-benchmark", benchmarkResults); err != nil {
		log.Fatalln(err)
	}
}

func saveBenchmarks(filename string, benchmarks map[string][]bench) error {
	if !strings.HasSuffix(filename, ".json") {
		filename += ".json"
	}
	allBytes, err := json.MarshalIndent(benchmarks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, allBytes, 0644)
}

type bench struct {
	Size  int     `json:"size"`
	Times []int64 `json:"times"`
}

func benchmark(count, size int, action func(size int) int64) bench {
	log.Printf("Count: %d, Size: %d", count, size)
	var times = make([]int64, 0, count)
	for i := 0; i < count; i++ {
		millis := action(size)
		times = append(times, millis)
	}
	return bench{
		Size:  size,
		Times: times,
	}
}

func simpleConcurrentRandom(size int) int64 {
	random := rand.New(randsafe.NewSource(time.Now().UnixNano()))
	wg := sync.WaitGroup{}
	wg.Add(size)
	start := time.Now()
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			lameUUID(random)
		}()
	}
	wg.Wait()
	end := time.Now()
	return end.Sub(start).Milliseconds()
}

func createPooledConcurrentRandom(poolSize int) func(size int) int64 {
	randomPool := randpool.New(poolSize, randomstd.NaiveConstructor)
	a := func(size int) int64 {
		wg := sync.WaitGroup{}
		wg.Add(size)
		start := time.Now()
		for i := 0; i < size; i++ {
			go func() {
				defer wg.Done()
				randomPool.Work(func(rand randomstd.Random) {
					lameUUID(rand)
				})
			}()
		}
		wg.Wait()
		end := time.Now()
		return end.Sub(start).Milliseconds()
	}
	return a
}

func lameUUID(random randomstd.Random) string {
	b := bytes.Buffer{}
	for i := 0; i < 8; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 4; i++ {
		b.WriteString(value(random))
	}
	b.WriteString("-")
	for i := 0; i < 12; i++ {
		b.WriteString(value(random))
	}
	return b.String()
}

func value(random randomstd.Random) string {
	i := random.Int63n(16)
	return strconv.FormatInt(i, 16)
}
