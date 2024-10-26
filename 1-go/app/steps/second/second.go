package second

import (
	"bufio"
	"fmt"
	"ssdl1-os-go/app/steps"
)

func Step(r *bufio.Reader) {
	fileName := "second-stage.txt"

	fmt.Printf("Введите стоку для записи в файл:\n")

	input, err := r.ReadBytes('\n')

	if err != nil {
		panic(err)
	}

	steps.HandleFile(fileName, input)
}
