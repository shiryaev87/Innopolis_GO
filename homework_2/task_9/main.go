/*
 9. Реализуйте тип-дженерик Numbers, который является слайсом численных типов.

Реализуйте следующие методы для этого типа:
* суммирование всех элементов,
* произведение всех элементов,
* сравнение с другим слайсом на равность,
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента,
* удаление элемента массива по значению,
* удаление элемента массива по индексу.
*/
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Numbers[T constraints.Ordered] []T

// метод суммирования элементов
func (n Numbers[T]) Sum() T {
	var sum T
	for _, v := range n {
		sum += v
	}
	return sum
}

// метод умножения элементов
func (n Numbers[T]) Product() T {
	if len(n) == 0 {
		var zero T
		return zero
	}
	product := n[0]
	for _, v := range n[1:] {
		product += v
	}
	return product
}

// сравнение слайсов
func (n Numbers[T]) Equals(other Numbers[T]) bool {
	if len(n) != len(other) {
		return false
	}
	for i, v := range n {
		if v != other[i] {
			return false
		}
	}
	return true
}

// проверка является ли элементов слайса
func (n Numbers[T]) IndexOf(element T) (int, bool) {
	for i, v := range n {
		if v == element {
			return i, true
		}
	}
	return -1, false
}

// удаление элемента слайса
func (n *Numbers[T]) RemoveValue(element T) {
	for i, v := range *n {
		if v == element {
			*n = append((*n)[:i], (*n)[i+1:]...)
			return
		}
	}
}

// уадаление элемента по индексу
func (n *Numbers[T]) RemoveIndex(index int) {
	if index < 0 || index >= len(*n) {
		return
	}
	*n = append((*n)[:index], (*n)[index+1:]...)
}

func main() {
	nums := Numbers[int]{1, 2, 3, 4, 5}
	fmt.Println("Сумма:", nums.Sum())
	fmt.Println("Произведение:", nums.Product())
	fmt.Println("Сравнение слайсов:", nums.Equals(Numbers[int]{1, 2, 3, 4, 5}))
	index, found := nums.IndexOf(3)
	fmt.Println("Индекс 3:", index, found)
	nums.RemoveValue(3)
	fmt.Println("После удаления элемента (3):", nums)
	nums.RemoveIndex(1)
	fmt.Println("После удаления элемента по индексу(1):", nums)
}
