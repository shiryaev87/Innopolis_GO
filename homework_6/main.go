/*
Домашняя работа N6
14 Горутины и синхронизация, ч. 3 / 07.06.2024

1. Напишите 2 функции: 1 функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
Вторая функция пишет эти данные в файл. Передайте в эти функции контекст.
Используйте context и waitgroup.

*/

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	//создаем контекст бэкграунд
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// Используем группу ожидания для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Создаем канала для передачи строк
	dataChannel := make(chan string)

	// Запуск горутины чтения ввода
	wg.Add(1)
	go readInput(ctx, &wg, dataChannel)

	// Запуск горутины функции записи в файл
	wg.Add(1)
	go writeFile(ctx, &wg, dataChannel)

	// Захват сигнала завершения программы
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнала прерывания
	fmt.Println("Нажмите Ctrl+C для отмены контекста...")
	<-sigChannel

	fmt.Println("Отменяем контекст")
	cancel()

	//ожидаем завершение вскех го рутин
	wg.Wait()
	fmt.Println("Завершение программы")
	//close(dataChannel)
}

// Функция для чтения ввода с консоли
func readInput(ctx context.Context, wg *sync.WaitGroup, dataChan chan<- string) {
	defer wg.Done() // Уменьшаем счетчик группы , чтобыкорректно завершилось
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение чтения")
			//закрываем канал
			close(dataChan)
			return
		default:
			fmt.Println("Введите текст:")
			input, _ := reader.ReadString('\n')
			dataChan <- input
		}

	}
}

// Функция для записи данных в файл
func writeFile(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan string) {
	defer wg.Done() //уменьшаем счетчик группы ожидания  при завершении го рутины
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

	for {
		select {

		case <-ctx.Done():
			fmt.Println("завершение записи в файл")
			return
		case inputText, ok := <-dataChan:
			if !ok {
				fmt.Println("Канал ввода закрыт")
				return
			}
			_, err := file.WriteString(inputText)
			if err != nil {
				fmt.Println("Ошибка записи в файл:", err)
				return
			}
		}
	}

}
