package bruteforce

import (
	"context"
	"encoding/hex"
	"hash"
	"math"
	"ssd-lab-pswd-go/src/generation"
	"sync"
)

func calculateButchSize(alphabet int, length int, buckets int) int {
	total := math.Pow(float64(alphabet), float64(length))
	return int(math.Floor(total / float64(buckets)))
}

func bruteforce(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, generator *generation.Generator, h hash.Hash, first int, last int, result *string, hash string) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		for i := first; i <= last; i++ {
			candidate := generator.Generate(i)
			candidateHash := hex.EncodeToString(h.Sum([]byte(candidate)))
			if candidateHash == hash {
				*result = candidate
				cancel()
				return
			}
		}
	}
}

func Bruteforce(generator *generation.Generator, h hash.Hash, goal string, alphabet int, length int, workers int) *string {
	result := new(string)

	wg := new(sync.WaitGroup)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	butchSize := calculateButchSize(alphabet, length, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		first := i * butchSize
		last := first + butchSize
		go bruteforce(ctx, cancel, wg, generator, h, first, last, result, goal)
	}

	wg.Wait()

	return result
}
