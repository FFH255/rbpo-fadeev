package steps

import (
	"fmt"
	"os"
)

func HandleFile(fileName string, input []byte) {

	fmt.Printf("Создание файла: %s\n", fileName)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
		fmt.Printf("Удаление файла: %s\n", fileName)
		if err := os.Remove(fileName); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Запись в файл: %s\n", fileName)
	if _, err := f.Write(input); err != nil {
		panic(err)
	}

	fmt.Printf("Чтение из файла: %s\n", fileName)
	b, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Содержимое файла: %s\n", string(b))
}
