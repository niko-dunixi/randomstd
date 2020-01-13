package randomstd

type RandomVintage interface {
	Seed(seed int64)
	Int63() int64
	Uint32() uint32
	Int31() int32
	Int() int
	Int63n(n int64)
	Int31n(n int32)
	Intn(n int) int
	Float64() float64
	Float32() float32
	Perm(n int)
	Shuffle(n int, swap func(i, j int))
	Read(p []byte) (n int, err error)
}

type Random interface {
	RandomVintage
	Uint64() uint64
}

// https://golang.org/doc/go1.8#math_rand

type Seed interface {
	Int63() (n int64)
	Uint64() (n uint64)
	Seed(seed int64)
}
