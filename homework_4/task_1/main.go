/*
Тема : 11 Основы многопоточности, ч.2 31.05.24
Домашняя работа 4. задание 1
*/
/*1. Напишите 2 функции:
	Первая функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
	Вторая функция пишет эти данные в файл. Свяжите эти функции каналом.
Работа приложения должна завершится при нажатии клавиш ctrl+c с кодом 0. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Создание канала для передачи строк
	dataChannel := make(chan string)

	// Захват сигнала завершения программы
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	// Запуск горутины чтения ввода
	go readInput(dataChannel)

	// Запуск горутины функции записи в файл
	go writeFile(dataChannel)

	// Ожидание сигнала завершения
	<-sigChannel
	fmt.Println("Завершение программы")
	close(dataChannel)
}

// Функция для чтения ввода с консоли
func readInput(dataChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите текст: ")
		if scanner.Scan() {
			text := scanner.Text()
			dataChan <- text
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			break
		}
	}
}

// Функция для записи данных в файл
func writeFile(dataChan chan string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	for data := range dataChan {
		_, err := file.WriteString(data + "\n")
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
	}
}
