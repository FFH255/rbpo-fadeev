package main

import (
	"context"
	"fmt"
	"ssd-lab-pswd-go/src/bruteforce"
	"ssd-lab-pswd-go/src/config"
	"ssd-lab-pswd-go/src/generator"
	"time"
)

func main() {
	fmt.Println("Программа запушена...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	start := time.Now()

	cfg := config.MustLoad()

	candidateChannel := generator.GeneratePasswords(&cfg.Generator)

	bruteforce.Bruteforce(ctx, cancel, &cfg.Bruteforce, candidateChannel, cfg.HashGoal)

	end := time.Since(start)

	fmt.Printf("Generated in %s\n", end)
	fmt.Println("Программа завершена")
}
