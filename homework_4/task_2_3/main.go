/*
Тема : 11 Основы многопоточности, ч.2 31.05.24
Домашняя работа 4. задания 2-3.

*/
/*2. Напишите функцию разделения массива чисел на массивы простых и составных чисел.
Для записи в массивы используйте два разных канала и го
рутины.
Важно, чтобы были использованы владельцы каналов.*/
/*3. Реализуйте функцию слияния двух каналов в один.*/

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("1) Исходный массив : \n", numbers)
	fmt.Println()
	simpleNumb, compositeNumb := separator1(numbers)

	fmt.Println("2) Массив простых чисел: \n", simpleNumb)
	fmt.Println()
	fmt.Println("3) Массив составных чисел: \n", compositeNumb)
	fmt.Println()

	fmt.Println("Результат слияния двух каналов в один")
	// два массива , перекладываем в каналы
	//time.Sleep(2 * time.Second)
	ch_1 := make(chan int)
	ch_2 := make(chan int)

	go func() {
		for _, num := range simpleNumb {
			ch_1 <- num
			//fmt.Println(num)
		}
		close(ch_1)
	}()

	/*for v := range ch_1 {
		fmt.Println(v)
	}
	*/
	go func() {
		for _, num := range compositeNumb {
			ch_2 <- num
		}
		close(ch_2)
	}()

	var merged []int
	for v := range merge(ch_1, ch_2) {
		//fmt.Println(v)
		merged = append(merged, v)
	}
	fmt.Println(merged)

}

func separator1(numbers1 []int) ([]int, []int) {
	simpleChan := make(chan int)    // канал для простых чисел
	compositeChan := make(chan int) // канал для составных чисел
	done := make(chan bool, 2)      // канал для закрытия, буферезированный, т.к. два канала ( простые и составные числа)

	var simpleNumbers []int
	var compositeNumbers []int

	//
	go func() { //создаем го рутину
		for _, num := range numbers1 { // проходим циклом по массиву с числами
			if isSimple(num) {
				simpleChan <- num // если протое число то добавляем в канал с простыми числам
			} else {
				compositeChan <- num // инчае добавляем в канала с  составными
			}
		}
		close(simpleChan) // закрываем оба канала
		close(compositeChan)

	}()

	//	го рутина для чтения простых чисел из канала
	go func() {
		for simple := range simpleChan {
			simpleNumbers = append(simpleNumbers, simple)

		}
		done <- true
	}()

	//го  рутина  для чтение составных чисел

	go func() {
		for composite := range compositeChan {
			compositeNumbers = append(compositeNumbers, composite)

		}
		done <- true

	}()

	// ожидание закрытия обоих каналов
	<-done
	<-done

	return simpleNumbers, compositeNumbers
}

func isSimple(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// функция объединения нескольких каналов в один
func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(c <-chan int) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()

	return out
}
