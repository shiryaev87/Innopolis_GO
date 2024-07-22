package main

import (
	"fmt"
)

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

// findLongestPath находит самый длинный путь в графе
func findLongestPath(graph [][]int) int {
	maxLength := 0
	visited := make([]bool, len(graph))

	for node := 0; node < len(graph); node++ {
		dfs(graph, node, visited, 0, &maxLength)
	}

	return maxLength
}

func main() {
	// Пример графа (в виде взвешенной матрицы смежности)
	graph := [][]int{
		{0, 0, 0, 0, 2, 0, 0, 0}, // Вершина
		{1, 0, 1, 0, 4, 3, 0, 0}, // Вершина 1
		{0, 0, 0, 0, 0, 1, 0, 0}, // Вершина 2
		{0, 0, 0, 0, 0, 0, 0, 0}, // Вершина 3
		{0, 0, 0, 3, 0, 0, 0, 1}, // Вершина 4
		{0, 0, 0, 5, 0, 0, 2, 0}, // Вершина 5
		{0, 0, 0, 2, 0, 0, 0, 0}, // Вершина 5
		{0, 0, 0, 1, 0, 0, 0, 0}, // Вершина 5
	}

	fmt.Println("Самый длинный путь в графе:", findLongestPath(graph))
}
