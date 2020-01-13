package randomstd

type Seed interface {
	Int63() (n int64)
	Uint64() (n uint64)
	Seed(seed int64)
}
