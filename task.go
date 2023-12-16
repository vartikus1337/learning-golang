// Задача: https://stepik.org/lesson/353243/step/6?unit=337227
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// type JSON struct {
// 	ID       int    `json:"ID"`
// 	Number   string `json:"Number"`
// 	Year     int    `json:"Year"`
// 	Students []struct {
// 		LastName   string `json:"LastName"`
// 		FirstName  string `json:"FirstName"`
// 		MiddleName string `json:"MiddleName"`
// 		Birthday   string `json:"Birthday"`
// 		Address    string `json:"Address"`
// 		Phone      string `json:"Phone"`
// 		Rating     []int  `json:"Rating"`
// 	} `json:"Students"`
// }


// Чтобы демаршалировать JSON в структуру, Unmarshal сопоставляет входящие ключи объекта с ключами, используемыми Marshal (либо имя поля структуры, либо его тег),
// предпочитая точное совпадение, но также принимая совпадение без учета регистра

type JSON struct {
	Students []struct {
		Rating []int
	}
}

type Out struct {
	Average float32 `json:"Average"`
}

func main() {
	file, err := os.Open("task.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	var s JSON

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}
	var countStudents, sumRating float32
	 
	for _, student := range s.Students {
		countStudents++
		for _, rating := range student.Rating {
			sumRating += float32(rating)
		}
	}
	out := Out{Average: sumRating / countStudents}
	data, err = json.Marshal(out)
	data, _ = json.MarshalIndent(out, "", "    ")
	fmt.Printf("%s \n", data)
}
