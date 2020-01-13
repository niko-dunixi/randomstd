// +build !go1.8

package randomstd

type Random interface {
	Seed(seed int64)
	Int63() int64
	Uint32() uint32
	Int31() int32
	Int() int
	Int63n(n int64) int64
	Int31n(n int32) int32
	Intn(n int) int
	Float64() float64
	Float32() float32
	Perm(n int) []int
	Shuffle(n int, swap func(i, j int))
	Read(p []byte) (n int, err error)
}
