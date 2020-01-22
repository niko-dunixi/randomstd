// +build !go1.8

package randomstd

type Random interface {
	ExpFloat64() float64
	Float32() float32
	Float64() float64
	Int() int
	Int31() int32
	Int31n(n int32) int32
	Int63() int64
	Int63n(n int64) int64
	Intn(n int) int
	NormFloat64() float64
	Perm(n int) []int
	Read(p []byte) (n int, err error)
	Seed(seed int64)
	Shuffle(n int, swap func(i, j int))
	Uint32() uint32
}
