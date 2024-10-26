package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"hash"
	"os"
	"ssd-lab-pswd-go/src/bruteforce"
	"ssd-lab-pswd-go/src/config"
	"ssd-lab-pswd-go/src/generation"
	"ssd-lab-pswd-go/src/pkg/io"
	"time"
)

func askPasswordHash() string {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("Введите хеш пароля: ")
	passwordHash := io.ReadLine(r)
	return passwordHash
}

func main() {
	cfg := config.MustLoad()

	var h hash.Hash

	if cfg.Alg == "sha256" {
		h = sha256.New()
	} else {
		h = md5.New()
	}

	passwordHash := askPasswordHash()

	generator := generation.New(cfg.Chars, cfg.Length)

	fmt.Printf("Поиск пароля...\n")

	start := time.Now()
	password := bruteforce.Bruteforce(generator, h, passwordHash, len(cfg.Chars), cfg.Length, cfg.Workers)
	end := time.Since(start)

	if password != nil && *password != "" {
		fmt.Printf("Пароль: %s\n", *password)
	} else {
		fmt.Printf("Пароль не найден\n")
	}

	fmt.Printf("Время выполнения: %s\n", end)
}
