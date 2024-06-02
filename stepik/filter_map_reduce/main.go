package main

import (
	"fmt"
)

// filter принимает слайс чисел и функцию, которая определяет критерий фильтрации
func filter(numbers []int, criteria func(int) bool) []int {
	var result []int
	for _, number := range numbers {
		if criteria(number) {
			result = append(result, number)
		}
	}
	return result
}

// Критерий фильтрации: оставить только четные числа
func isEven(number int) bool {
	return number%2 == 0
}

// Критерий фильтрации: оставить только числа больше 10
func isGreaterThanTen(number int) bool {
	return number > 10
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 10, 12, 15, 18}

	// Использование функции filter для получения только четных чисел
	evenNumbers := filter(numbers, isEven)
	fmt.Println("Even numbers:", evenNumbers)

	// Использование функции filter для получения чисел больше 10
	numbersGreaterThanTen := filter(numbers, isGreaterThanTen)
	fmt.Println("Numbers greater than 10:", numbersGreaterThanTen)
}
