/*
 5. Перепишите задачу #3 с использованием структуры-дженерик Cache, изученной на семинаре.

Храните в кеше таблицы студентов и предметов.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

type Subject struct {
	Id    int    `json:"id"`
	Title string `json:"name"`
}

type Result struct {
	SubjectID int `json:"object_id"`
	StudentID int `json:"student_id"`
	Result    int `json:"result"`
}

type Data struct {
	Students []Student `json:"students"`
	Subjects []Subject `json:"objects"`
	Results  []Result  `json:"results"`
}

type resultTable struct {
	studentName string
	grade       int
	subjectName string
	result      int
}

// обобщенная структура кэш
type Cache[K comparable, V any] struct {
	m map[K]V
}

var users Cache[int, *Student]
var subjects Cache[int, *Subject]

func (c *Cache[K, V]) Init() {
	c.m = make(map[K]V)
}

func (c Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}
func (c *Cache[K, V]) Get(key K) (V, bool) {
	k, ok := c.m[key]
	return k, ok
}
func main() {
	resTable := []resultTable{}
	headers := []string{"Student", "Grade", "Subject", "Result"}
	//Читаем файл
	byteValue, err := os.ReadFile("homework_2/task_5_6_7/dz3.json")
	if err != nil {
		log.Fatal(err)
	}
	//парсим json в структуру data

	var data Data
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatal(err)
	}
	users.Init()
	subjects.Init()

	//studentMap := make(map[int]Student)
	//objectMap := make(map[int]Subject)
	//перекладываю студентов
	for _, student := range data.Students {
		//studentMap[student.Id] = student
		users.Set(student.Id, &student)
	}
	//fmt.Println(users)
	//перекладываю предметы в мапу
	for _, object := range data.Subjects {
		//objectMap[object.Id] = object
		subjects.Set(object.Id, &object)
	}

	// Печать заголовка таблицы'

	var res resultTable
	// Заполнение и печать содержания таблицы
	for _, result := range data.Results {
		//student := studentMap[result.StudentID]
		student, _ := users.Get(result.StudentID)
		//object := objectMap[result.SubjectID]
		object, _ := subjects.Get(result.SubjectID)
		res.studentName = student.Name
		res.grade = student.Grade
		res.subjectName = object.Title
		res.result = result.Result
		resTable = append(resTable, res)
	}
	// считаем ширину столбцов, чтобы ровно вывести таблицу
	// вычисляем ширину каждой колонки
	//создаем мапу , в которой храним ширину столбцов
	columnWidths := make([]int, len(headers))
	//проходим циклом по заголовку
	for i, headColumn := range headers {
		columnWidths[i] = len(headColumn)
	}
	//проходим циклом по строкам

	for _, row := range resTable {
		if len(row.studentName) > columnWidths[0] {
			columnWidths[0] = len(row.studentName)
		}

		if len(strconv.Itoa(row.grade)) > columnWidths[1] {
			columnWidths[1] = len(strconv.Itoa(row.grade))
		}

		if len(row.subjectName) > columnWidths[2] {
			columnWidths[2] = len(row.subjectName)
		}

		if len(strconv.Itoa(row.result)) > columnWidths[1] {
			columnWidths[3] = len(strconv.Itoa(row.result))
		}

	}
	// печать заголовка таблиц
	divider := strings.Repeat("-", totalWidth(columnWidths)+len(columnWidths)+9)
	fmt.Println(divider)

	// вывод заголовков
	for i, header := range headers {
		fmt.Printf("%s", "| ")
		fmt.Printf("%-*s", columnWidths[i]+1, header)
	}
	fmt.Printf("%s\n", "|")

	fmt.Println(divider)

	//печать таблицы данных
	for _, row := range resTable {
		fmt.Printf("%s", "| ")
		fmt.Printf("%-*s", columnWidths[0]+1, row.studentName)
		fmt.Printf("%s", "| ")
		fmt.Printf("%-*d", columnWidths[1]+1, row.grade)
		fmt.Printf("%s", "| ")
		fmt.Printf("%-*s", columnWidths[2]+1, row.subjectName)
		fmt.Printf("%s", "| ")
		fmt.Printf("%-*d", columnWidths[3]+1, row.result)
		fmt.Printf("%s\n", "|")
	}
	fmt.Println(divider)
	// расчет средних значений
	calcAvarage(resTable)

	// Фильтрация и вывод круглых отличников
	fmt.Println("Best students:")
	honorStudents := Filter(data.Students, func(s Student) bool {
		return isHonorStudent(s.Id, data.Results)
	})
	for _, student := range honorStudents {
		fmt.Println(student.Name)
	}

}

// totalWidth вычисляет суммарную ширину всех колонок
func totalWidth(columnWidths []int) int {
	total := 0
	for _, width := range columnWidths {
		total += width
	}
	return total
}

// Функция для подсчета средницх оценок и вывода сводной таблицы
func calcAvarage(grades []resultTable) {
	//структуры для хранения промежуточных данных
	classSubjectScores := make(map[int]map[string][]float64)
	subjectScores := make(map[string][]float64)
	// Заполняем мапы данными об оценках . из  слайса grades собираем данные в мапу мапу  ,
	//где ключ - класс, а значение - мапа, в которой ключ предмет, а значение - слайс со оценками

	for _, grade := range grades {
		//проверяем, есть ли в мапе уже ключ , если нет, то добавляем
		if _, ok := classSubjectScores[grade.grade]; !ok {
			classSubjectScores[grade.grade] = make(map[string][]float64)
		}
		classSubjectScores[grade.grade][grade.subjectName] = append(classSubjectScores[grade.grade][grade.subjectName], float64(grade.result))
		subjectScores[grade.subjectName] = append(subjectScores[grade.subjectName], float64(grade.result))
	}

	//fmt.Println(classSubjectScores)
	//fmt.Println(subjectScores)

	//сортировка классов и предметов для  вывода
	classes := make([]int, 0, len(classSubjectScores))
	for class := range classSubjectScores {
		classes = append(classes, class)
	}
	sort.Ints(classes)
	//fmt.Println(classes)

	subjects := make([]string, 0, len(subjectScores))
	for subject := range subjectScores {
		subjects = append(subjects, subject)

	}

	sort.Strings(subjects)
	//fmt.Println(subjects)
	fmt.Printf("%s\n", "Avarage results by subjects:")
	// считаем среднее значение по классам, по предметам
	for _, subject := range subjects {
		totalScore := 0.0
		totalCount := 0
		fmt.Printf("%-*s\n", 10, subject)
		for _, class := range classes {

			if scores, ok := classSubjectScores[class][subject]; ok {
				classTotal := 0.0
				for _, score := range scores {
					classTotal += score
				}
				classAverage := classTotal / float64(len(scores))
				//fmt.Printf("%s | %d | %.2f\n", subject, class, classAverage)

				//fmt.Printf("%s", "| ")
				//fmt.Printf("%-*s", 10, subject)
				fmt.Printf("%s", "| ")
				fmt.Printf("%-*d", 3, class)
				fmt.Printf("%s", " grade | ")
				fmt.Printf("%-*.2f\n", 10, classAverage)
				//	fmt.Printf("%-*s", 10, subject)
				//	fmt.Printf("%-*i", 10, class)
				///	fmt.Printf("%.2f\n", 10, classAverage)
				totalScore += classTotal
				totalCount += len(scores)
			} else {
				fmt.Printf("%s | %d | -\n", subject, class)
			}

		}
		subjectAverage := totalScore / float64(totalCount)
		//fmt.Printf("%s | Все классы | %.2f\n", subject, subjectAverage)
		//fmt.Printf("%s", "| ")
		//fmt.Printf("%-*s", 10, subject)
		fmt.Println("------------------------------------")
		fmt.Printf("%s", "| ")
		fmt.Printf("%-*s", 10, "Mean")
		fmt.Printf("%s", "| ")
		fmt.Printf("%-*.2f\n", 10, subjectAverage)
		fmt.Println("------------------------------------")
		fmt.Println()
	}
}

// Функция фильтрации списка студентов
func Filter(students []Student, fn func(Student) bool) []Student {
	var result []Student
	for _, student := range students {
		if fn(student) {
			result = append(result, student)
		}
	}
	return result
}

// Функция проверки, является ли студент круглым отличником
func isHonorStudent(studentID int, results []Result) bool {
	for _, result := range results {
		if result.StudentID == studentID && result.Result != 5 {
			return false
		}
	}
	return true
}

func Reduce[T1 float64, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	r := init
	for _, v := range s {
		r = f(v, r)
	}
	return r
}
