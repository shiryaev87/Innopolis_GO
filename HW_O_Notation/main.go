package main

import (
	"fmt"
	"sync"
)

func MinEl2(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	//создаем wait group
	var wg sync.WaitGroup
	// добавляем 2 потока , т.к.  вызываем две горутины
	wg.Add(2)

	var t1, t2 int

	go func() {
		defer wg.Done() //закрываем
		t1 = MinEl2(a[:len(a)/2])
	}()

	go func() {
		defer wg.Done() //закрываем
		t2 = MinEl2(a[len(a)/2:])
	}()

	wg.Wait()

	if t1 <= t2 {
		return t1
	}
	return t2
}

func main() {
	a := []int{3, 5, 1, 2, 4, 8, 7}

	// Засекаем начало времени выполнения
	//start := time.Now()

	// Выполнение функции
	result := MinEl2(a)
	//end := time.Now()
	//time.Sleep(1 * time.Second)
	//duration := end.Sub(start)
	//duration := time.Since(start)
	//	duration := duration.Nanoseconds()
	fmt.Println("Минимальный элемент:", result)
	//fmt.Println("Время выполнения", duration.Nanoseconds())
}
