package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	var count_positive_answers int = 0

	// Определение флага для имени файла
	fileFlag := flag.String("file", "homework_1/task_1/problems.csv", "Укажите путь к файлу с вопросами")

	//Определение флага перемешивания вопросов
	shuffleFlag := flag.Bool("shuffle", false, "Флаг перемешивания вопросов")
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Создаем новый csv.Reader
	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = 2

	// Инициализируем слайс
	var problems []Problem

	// Считываем  строки из файла

	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println("Ошибка чтения CSV:", err)
			break
		}

		p := Problem{
			Question: record[0],
			Answer:   record[1],
		}
		problems = append(problems, p)
	}
	//Перемешиваем вопросы в зависимости от флага
	if *shuffleFlag {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	var userAnswer string
	for index, p := range problems {
		fmt.Println("Вопрос №", index+1, ":", p.Question)
		fmt.Print("Введите ваш ответ: ")
		_, _ = fmt.Scanln(&userAnswer)

		if strings.ToLower(userAnswer) == strings.ToLower(p.Answer) {
			count_positive_answers++
		}
	}

	fmt.Println("\nКоличество неправильных ответов:", len(problems)-count_positive_answers)
	fmt.Println("Количество правильных ответов:", count_positive_answers)

}
