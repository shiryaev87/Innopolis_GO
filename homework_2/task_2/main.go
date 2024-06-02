package main

import (
	"fmt"
	"sort"
)

/*
 2. Подсчет голосов.

Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
Пример.
Вход: ["Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"]
Вывод: [{Ann, 3}, {Kate, 2}, {Peter, 1}, {Helen, 1}]

	type Candidate struct {
		Name  string
		Votes int
	}
*/

type Candidate struct {
	Name  string
	Votes int
}

func main() {

	allVotes := []string{"Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"}
	result := countVotes(allVotes)
	//сортируем слайс

	for _, Candidate := range result {
		fmt.Printf("{%s, %d}\n", Candidate.Name, Candidate.Votes)
	}
}

// функция подсчета голосов
func countVotes(votes []string) []Candidate {
	// мапа для подсчета голосов
	votesCount := make(map[string]int)

	//считаем голоса, проходим по массиву голосов
	for _, name := range votes {
		votesCount[name]++
	}

	// перекладываем мапу в срез
	var candidates []Candidate
	for name, count := range votesCount {
		candidates = append(candidates, Candidate{Name: name, Votes: count})
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Votes > candidates[j].Votes
	})
	return candidates
}
