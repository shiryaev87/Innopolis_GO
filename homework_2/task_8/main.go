/*
 8. Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.

Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.
*/
package main

import (
	"fmt"
)

// IsEqualArrays
func IsEqualArrays[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[T]int)
	bMap := make(map[T]int)

	for _, v := range a {
		aMap[v]++
	}

	for _, v := range b {
		bMap[v]++
	}

	for k, v := range aMap {
		if bMap[k] != v {
			return false
		}
	}

	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 3, 2, 1}
	arr3 := []float64{1.0, 2.0, 3.0, 4.0}
	arr4 := []float64{4.0, 3.0, 2.0, 1.0}

	fmt.Println(IsEqualArrays(arr1, arr2))
	fmt.Println(IsEqualArrays(arr3, arr4))

}
