package cache

/*
import "github.com/dgraph-io/ristretto"

func InitRistrettoCache() (*ristretto.Cache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e4,     // 10x number of items you expect
		MaxCost:     1 << 20, // 1 MB or arbitrary cost limit
		BufferItems: 64,      // recommended
	})
	if err != nil {
		panic(err)
	}
	return cache, err
} */
