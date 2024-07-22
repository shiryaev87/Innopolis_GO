package main

import (
	"fmt"
	"strconv"
)

func EvalSequence(matrix [][]int, userAnswer []int) int {

	// validation
	maxGrade := calMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	percent := userGrade * 100 / maxGrade

	return percent
}

func calMaxGrade(matrix [][]int) int {
	maxLength := 0
	visited := make([]bool, len(matrix))

	for node := 0; node < len(matrix); node++ {
		dfs(matrix, node, visited, 0, &maxLength)
	}

	return maxLength
}

func calcUserGrade(matrix [][]int, userAnswer []int) int {
	totalLength := 0

	for i := 0; i < len(userAnswer)-1; i++ {
		from := userAnswer[i]
		to := userAnswer[i+1]

		//Проверка, чтомжду узлами нет ребра.

		/*	if matrix[from][to] == 0 {
				log.Fatalf("Нет пути между вершинами %d и %d", from, to)
			}
		*/
		totalLength += matrix[from][to]
	}

	return totalLength
}

// isZeroDiagonalSquareMatrix проверяет, является ли матрица квадратной и содержат ли все элементы на главной диагонали нули.
func checkValidMatrix(matrix [][]int) (bool, string) {
	n := len(matrix)
	err := ""
	// Проверяем, что матрица не пустая
	if n == 0 {
		err = "Ошибка. Матрица пустая"
		return false, err
	}
	// Проверяем, что матрица квадратная
	for _, row := range matrix {
		if len(row) != n {
			err = "Ошибка. Матрица не квадратная"
			return false, err
		}
	}
	// Проверяем, что все элементы на главной диагонали равны нулю
	for i := 0; i < n; i++ {
		if matrix[i][i] != 0 {
			err = "Ошибка. В графе не должно быть петель. Элементы на главной диагонали не равны нулю. Строка = " + strconv.Itoa(i) + ". Столбец = " + strconv.Itoa(i) + "."
			return false, err
		}
	}
	return true, ""
}

// areElementsUniqueAndValid проверяет, что элементы массива уникальны и не превышают размер матрицы n.
func checkValidUserAnswer(arr []int, n int) (bool, string) {
	elementSet := make(map[int]bool)
	for _, value := range arr {
		if value >= n {
			return false, "Ошибка. В ответе пользователя есть элемент, который не является вершиной графа. Элемент из ответа пользователя " + strconv.Itoa(value) + ". "
		}
		if _, exists := elementSet[value]; exists {
			return false, "Ошибка. Элементы в ответах пользователя должны быть уникальными. Повторяющейся ответ  " + strconv.Itoa(value) + ". "
		}
		elementSet[value] = true
	}
	return true, ""
}

// dfs выполняет глубинный поиск и находит максимальную длину пути
func dfs(graph [][]int, node int, visited []bool, currentLength int, maxLength *int) {
	visited[node] = true
	isEnd := true

	for neighbor := 0; neighbor < len(graph); neighbor++ {
		if graph[node][neighbor] > 0 && !visited[neighbor] {
			isEnd = false
			dfs(graph, neighbor, visited, currentLength+graph[node][neighbor], maxLength)
		}
	}

	visited[node] = false // откат изменений
	if isEnd {
		if currentLength > *maxLength {
			*maxLength = currentLength
		}
	}
}

func main() {
	// Пример графа (в виде взвешенной матрицы смежности)
	graph := [][]int{
		{0, 0, 0, 0, 2, 0, 0, 0}, // Вершина 0
		{1, 0, 1, 0, 4, 3, 0, 0}, // Вершина 1
		{0, 0, 0, 0, 0, 1, 0, 0}, // Вершина 2
		{0, 0, 0, 0, 0, 0, 0, 0}, // Вершина 3
		{0, 0, 0, 3, 0, 0, 0, 1}, // Вершина 4
		{0, 0, 0, 5, 0, 0, 2, 0}, // Вершина 5
		{0, 0, 0, 2, 0, 0, 0, 0}, // Вершина 6
		{0, 0, 0, 1, 0, 0, 0, 0}, // Вершина 7
	}

	//fmt.Println(graph)
	//fmt.Println("Самый длинный путь в графе:", findLongestPath(graph))

	// Пример пути (в порядке прохождения вершин)

	userAnswer := []int{1, 2, 6, 3}

	isValidMatrix, errMatrix := checkValidMatrix(graph)
	isValidUserAnswer, errUserAnswer := checkValidUserAnswer(userAnswer, len(graph))
	if isValidMatrix && isValidUserAnswer {
		fmt.Println("Оценка за ответ %", EvalSequence(graph, userAnswer))
	} else {
		if errMatrix != "" {
			fmt.Println(errMatrix)
		}

		if errUserAnswer != "" {
			fmt.Println(errUserAnswer)
		}
	}

	//length := calculatePathLength(graph, path)
	//fmt.Printf("Длина пути: %d\n", length)

}
