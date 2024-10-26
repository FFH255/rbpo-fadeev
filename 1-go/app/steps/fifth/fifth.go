package fifth

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/sqweek/dialog"
	"io"
	"os"
	"path/filepath"
)

// zipFile создает zip-архив и добавляет в него указанный файл
func zipFile(filename string, archiveName string) error {
	archive, err := os.Create(archiveName)
	if err != nil {
		return fmt.Errorf("ошибка при создании архива: %v", err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Добавление файла в архив
	w, err := zipWriter.Create(filepath.Base(filename))
	if err != nil {
		return fmt.Errorf("ошибка при добавлении файла в архив: %v", err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		return fmt.Errorf("ошибка при копировании данных в архив: %v", err)
	}

	fmt.Printf("Файл %s успешно добавлен в архив %s\n", filename, archiveName)
	return nil
}

// unzipFile разархивирует файл и выводит его содержимое
func unzipFile(archiveName string, outputDir string) error {
	archive, err := zip.OpenReader(archiveName)
	if err != nil {
		return fmt.Errorf("ошибка при открытии архива: %v", err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		fmt.Printf("Распаковка файла: %s\n", f.Name)

		outFile, err := os.Create(filepath.Join(outputDir, f.Name))
		if err != nil {
			return fmt.Errorf("ошибка при создании распакованного файла: %v", err)
		}
		defer outFile.Close()

		archiveFile, err := f.Open()
		if err != nil {
			return fmt.Errorf("ошибка при чтении файла из архива: %v", err)
		}
		defer archiveFile.Close()

		_, err = io.Copy(outFile, archiveFile)
		if err != nil {
			return fmt.Errorf("ошибка при записи распакованного файла: %v", err)
		}

		fmt.Printf("Файл %s успешно распакован\n", f.Name)
	}
	return nil
}

// removeFiles удаляет исходный файл и архив
func removeFiles(files ...string) error {
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			return fmt.Errorf("ошибка при удалении файла %s: %v", file, err)
		}
		fmt.Printf("Файл %s успешно удален\n", file)
	}
	return nil
}

func Step(_ *bufio.Reader) {
	// Открываем проводник для выбора файла
	filename, err := dialog.File().Title("Выберите файл для архивации").Load()
	if err != nil {
		fmt.Println("Файл не выбран:", err)
		return
	}

	// Имя архива
	archiveName := "archive.zip"

	// Создаем архив и добавляем в него выбранный файл
	err = zipFile(filename, archiveName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Распаковываем файл и выводим его данные
	err = unzipFile(archiveName, ".")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Удаляем исходный файл и архив
	err = removeFiles(filename, archiveName)
	if err != nil {
		fmt.Println(err)
	}
}
