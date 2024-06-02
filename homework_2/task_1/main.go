package main

/* 1. Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
Примеры:
1. Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
2. Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
3. Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].
*/

import (
	"fmt"
	"sort"
)

func main() {
	/*
		общий порядок решения
		1. Удаляем дубликаты из каждого слайса
		2. Делаем мапу- где ключ- число, значение - количество вхождений во все слайсы
		3. Те ключи , у которых количество вхождений = количеству слайсов , будут решением. добавляем их в слайс результат
		4. Сортируем слайс и выводим
	*/

	//  массив слайс слайсов,, который приходит на вход
	sliceOfSlices := [][]int{
		{4, 4, 1, 2, 3, 2, 4},
		{3, 2},
		{1, 2, 3, 2},
		//{},
	}
	//слайс с уникальными значениям, в него будем добавлять слайсы после удаления дубликатов
	sliceUnique := [][]int{}

	for _, sliceOfSlice := range sliceOfSlices {
		//проходим по слайсу, передаем слайс в функцию удаления дубликатов
		sliceUnique = append(sliceUnique, removeDuplicates(sliceOfSlice))

	}

	//fmt.Println(sliceUnique)
	//мапа для подсчета элементов
	mapElementCount := make(map[int]int)
	// количество слайсов
	numSlices := len(sliceUnique)

	for _, slice := range sliceUnique {
		for _, element := range slice {
			mapElementCount[element]++
		}
	}

	//результирующий слайс -
	result := []int{}
	//проходим по мапе "количество элевентов" , если количество слайсов равно количеству элементов , тогда добавляем в слайс  result
	for key, count := range mapElementCount {
		if count == numSlices {
			result = append(result, key)
		}
	}
	//сортируем
	sort.Ints(result)
	fmt.Println(result)

}

func removeDuplicates(slice []int) []int {
	//создание мапы для отслеживания уникальных элементов ключ -  int число, значение  - есть или нет уже этот элемент в мапе
	uniqueMap := make(map[int]bool)
	// создания слайса для хранения уникальных элементов
	uniqueSlice := []int{}
	// проход по оригинальному слайсу
	for _, element := range slice {
		//проверка
		if !uniqueMap[element] {
			uniqueMap[element] = true
			uniqueSlice = append(uniqueSlice, element)
		}

	}
	//	fmt.Println(uniqueMap)
	return uniqueSlice
}
