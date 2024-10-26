package bruteforce

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"sync"
)

type Config struct {
	Algorithm string `yaml:"algorithm" default:"md5"`
	Workers   int    `yaml:"workers" default:"10"`
}

func worker(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, alg hash.Hash, candidatesChan <-chan string, goalHash string, result *string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return

		case candidate, ok := <-candidatesChan:
			if !ok {
				return
			}
			candidateHash := alg.Sum([]byte(candidate))
			if hex.EncodeToString(candidateHash[:]) == goalHash {
				*result = candidate
				cancel()
				return
			}
		}
	}
}

func Bruteforce(ctx context.Context, cancel context.CancelFunc, cfg *Config, candidatesChan <-chan string, goalHash string) {
	wg := &sync.WaitGroup{}

	var alg hash.Hash

	if cfg.Algorithm == "sha256" {
		alg = sha256.New()
	} else {
		alg = md5.New()
	}

	result := new(string)

	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go worker(ctx, cancel, wg, alg, candidatesChan, goalHash, result)
	}

	wg.Wait()

	if result != nil && *result != "" {
		fmt.Printf("password: %s\n", *result)
	} else {
		fmt.Printf("no password found\n")
	}
}
