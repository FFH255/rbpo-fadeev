package main

import (
	"bufio"
	"fmt"
	"os"
	"ssdl1-os-go/app/steps/fifth"
	"ssdl1-os-go/app/steps/first"
	"ssdl1-os-go/app/steps/forth"
	"ssdl1-os-go/app/steps/second"
	"ssdl1-os-go/app/steps/third"
)

type Step func(r *bufio.Reader)

func pressEnterToContinue(r *bufio.Reader) {
	fmt.Println("\nНажмите Enter, чтобы продолжить...")
	_, _ = r.ReadBytes('\n')
}

func main() {
	stages := []Step{first.Step, second.Step, third.Step, forth.Step, fifth.Step}

	r := bufio.NewReader(os.Stdin)

	fmt.Print("Начало выполнения программы\n")

	for i, stage := range stages {
		fmt.Printf("Переход в шагу %d:\n\n", i+1)
		stage(r)
		pressEnterToContinue(r)
	}

	fmt.Print("Завершение выполнения программы\n")
}
