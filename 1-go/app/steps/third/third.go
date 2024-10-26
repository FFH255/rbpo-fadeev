package third

import (
	"bufio"
	"encoding/json"
	"fmt"
	"ssdl1-os-go/app/steps"
)

func Step(r *bufio.Reader) {
	fileName := "third-stage.json"

	// Чтение JSON данных из консоли
	var inputData map[string]interface{}
	fmt.Println("Введите JSON данные:")

	// Чтение всего ввода с консоли
	inputBytes, err := r.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	// Парсинг JSON данных
	err = json.Unmarshal(inputBytes, &inputData)
	if err != nil {
		panic(err)
	}

	// Конвертируем данные обратно в JSON и записываем в файл
	outputBytes, err := json.MarshalIndent(inputData, "", "  ")
	if err != nil {
		panic(err)
	}

	steps.HandleFile(fileName, outputBytes)
}
