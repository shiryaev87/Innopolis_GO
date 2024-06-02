package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// Добавляем флаги для определения операции и имени файла
	createFlag := flag.String("create", "", "Имя файла для создания")
	readFlag := flag.String("read", "", "Имя файла для чтения")
	deleteFlag := flag.String("delete", "", "Имя файла для удаления")
	flag.Parse()

	// Выполняем операцию в зависимости от установленных флагов
	if *createFlag != "" {
		content := []byte("Пример содержимого файла.\n")
		err := os.WriteFile(*createFlag, content, fs.FileMode(0644))
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
		} else {
			fmt.Println("Файл успешно создан.")
		}
	}

	if *readFlag != "" {
		data, err := os.ReadFile(*readFlag)
		if err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
		} else {
			fmt.Println("Содержимое файла:", string(data))
		}
	}

	if *deleteFlag != "" {
		err := os.Remove(*deleteFlag)
		if err != nil {
			fmt.Println("Ошибка при удалении файла:", err)
		} else {
			fmt.Println("Файл успешно удален.")
		}
	}
	fmt.Println("привет")
}
