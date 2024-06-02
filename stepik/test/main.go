package main

import (
	"fmt"
	"sort"
)

// Структура для хранения данных об оценках
type Grade struct {
	Student string
	Class   int
	Subject string
	Score   float64
}

// Функция для вычисления средних оценок и вывода сводной таблицы
func calculateAverages(grades []Grade) {
	// Структура для хранения промежуточных данных
	classSubjectScores := make(map[int]map[string][]float64)
	subjectScores := make(map[string][]float64)

	// Заполняем мапы данными об оценках . из  слайса grades собираем данные в мапу мапу  , где ключ - класс, а значение - мапа, в которой ключ предмет, а значение - слайс со оценками
	for _, grade := range grades {
		// проверяем есть ли в мапе уже ключ, если нет , то добавляем
		if _, ok := classSubjectScores[grade.Class]; !ok {
			classSubjectScores[grade.Class] = make(map[string][]float64)
		}
		classSubjectScores[grade.Class][grade.Subject] = append(classSubjectScores[grade.Class][grade.Subject], grade.Score)
		subjectScores[grade.Subject] = append(subjectScores[grade.Subject], grade.Score)
	}

	// Сортировка классов и предметов для красивого вывода
	classes := make([]int, 0, len(classSubjectScores))
	for class := range classSubjectScores {
		classes = append(classes, class)
	}
	sort.Ints(classes)

	subjects := make([]string, 0, len(subjectScores))
	for subject := range subjectScores {
		subjects = append(subjects, subject)
	}
	sort.Strings(subjects)

	// Вывод сводной таблицы
	fmt.Println("Сводная таблица по предметам:")
	fmt.Println("Предмет | Класс | Средняя оценка")
	fmt.Println("------------------------------------")

	for _, subject := range subjects {
		totalScore := 0.0
		totalCount := 0

		for _, class := range classes {
			if scores, ok := classSubjectScores[class][subject]; ok {
				classTotal := 0.0
				for _, score := range scores {
					classTotal += score
				}
				classAverage := classTotal / float64(len(scores))
				fmt.Printf("%s | %d | %.2f\n", subject, class, classAverage)

				totalScore += classTotal
				totalCount += len(scores)
			} else {
				fmt.Printf("%s | %d | -\n", subject, class)
			}
		}

		subjectAverage := totalScore / float64(totalCount)
		fmt.Printf("%s | Все классы | %.2f\n", subject, subjectAverage)
		fmt.Println("------------------------------------")
	}
}

func main() {
	grades := []Grade{
		{"Alice", 8, "Math", 4},
		{"Bob", 8, "Math", 4},
		{"Charlie", 9, "Math", 5},
		{"Alice", 8, "Science", 5},
		{"Bob", 8, "Science", 3},
		{"Charlie", 9, "Science", 3},
	}

	calculateAverages(grades)
}
