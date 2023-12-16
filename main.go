package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main1() {
	// JSON encoding работает с только заглавной буквой
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.

	data, err := json.Marshal(s) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)

	data, err = json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", data)

	data = []byte(`{"Name":"John 123","Age":123,"Status":false,"Values":[1,2,3]}`)
	var sNew myStruct
	if err := json.Unmarshal(data, &sNew); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", sNew) // {John 123 123 false [1 2 3]}

	data = bytes.Trim(data, "{")

	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}
	fmt.Printf("%s \n", data)

	// Annotating Structures

	type myStructAnn struct {
		// при кодировании / декодировании будет использовано имя name, а не Name
		Name string `json:"name"`

		// при кодировании / декодировании будет использовано то же имя (Age),
		// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
		// то при кодировании оно будет опущено
		Age int `json:",omitempty"`

		// при кодировании / декодировании поле всегда игнорируется
		Status bool `json:"-"`
	}

	m := myStructAnn{Name: "John Connor", Age: 0, Status: true}

	data, err = json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", data) // {"name":"John Connor"}

}
