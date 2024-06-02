package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Структуры данных для JSON
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

type Object struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	ObjectID  int `json:"object_id"`
	StudentID int `json:"student_id"`
	Result    int `json:"result"`
}

type Data struct {
	Students []Student `json:"students"`
	Objects  []Object  `json:"objects"`
	Results  []Result  `json:"results"`
}

func main() {
	jsonData := `
    {
        "students": [
            {"id": 1, "name": "Ann", "grade": 9},
            {"id": 2, "name": "Kate", "grade": 9},
            {"id": 3, "name": "Peter", "grade": 9},
            {"id": 4, "name": "John", "grade": 10},
            {"id": 5, "name": "Alex", "grade": 10},
            {"id": 9, "name": "Selena", "grade": 10},
            {"id": 10, "name": "Angela", "grade": 10},
            {"id": 11, "name": "Tim", "grade": 11},
            {"id": 12, "name": "Aaron", "grade": 11},
            {"id": 13, "name": "Jeremy", "grade": 11}
        ],
        "objects": [
            {"id": 1, "name":  "Math"},
            {"id": 2, "name":  "Biology"},
            {"id": 3, "name":  "Geography"}
        ],
        "results": [
            {"object_id": 1, "student_id": 1, "result":  4},
            {"object_id": 2, "student_id": 1, "result":  5},
            {"object_id": 3, "student_id": 1, "result":  3},
            {"object_id": 1, "student_id": 2, "result":  2},
            {"object_id": 2, "student_id": 2, "result":  4},
            {"object_id": 3, "student_id": 2, "result":  5},
            {"object_id": 1, "student_id": 3, "result":  5},
            {"object_id": 2, "student_id": 3, "result":  5},
            {"object_id": 3, "student_id": 3, "result":  5},
            {"object_id": 1, "student_id": 4, "result":  3},
            {"object_id": 2, "student_id": 4, "result":  3},
            {"object_id": 3, "student_id": 4, "result":  2}
        ]
    }`

	// Декодирование JSON
	var data Data
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error decoding JSON: %s", err)
	}

	// Создание мапу для быстрого поиска
	studentMap := make(map[int]Student)
	objectMap := make(map[int]Object)
	//перекладываю студентов в мапу
	for _, student := range data.Students {
		studentMap[student.ID] = student
	}
	//перекладываю предметы в мапу
	for _, object := range data.Objects {
		objectMap[object.ID] = object
	}

	// Печать заголовка таблицы
	fmt.Printf("%-10s %-10s %-10s %-10s\n", "StudentStudent", "Grade", "Object", "Result")

	// Заполнение и печать содержани	я таблицы
	for _, result := range data.Results {
		student := studentMap[result.StudentID]
		object := objectMap[result.ObjectID]
		fmt.Printf("%-10s %-10d %-10s %-10d\n", student.Name, student.Grade, object.Name, result.Result)
	}
}
