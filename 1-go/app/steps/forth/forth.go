package forth

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"ssdl1-os-go/app/steps"
)

type XML struct {
	XMLName    xml.Name
	Attributes []xml.Attr `xml:",any,attr"`
	Content    string     `xml:",chardata"`
	Children   []XML      `xml:",any"`
}

func Step(r *bufio.Reader) {
	fileName := "forth-stage.xml"

	// Чтение XML данных из консоли
	var inputData XML
	fmt.Println("Введите XML данные:")

	// Чтение всего ввода с консоли
	inputBytes, err := r.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	// Парсинг XML данных
	err = xml.Unmarshal(inputBytes, &inputData)
	if err != nil {
		panic(err)
	}

	// Конвертируем данные обратно в XML и записываем в файл
	outputBytes, err := xml.Marshal(inputData)
	if err != nil {
		panic(err)
	}

	steps.HandleFile(fileName, outputBytes)
}
