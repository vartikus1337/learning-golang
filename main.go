package main

import "fmt"

func fullInfo(name, surname string, age int) (string, bool) {
	fullName := name + " " + surname + " "
	var isAdult bool
	if age >= 18 {
		isAdult = true
	} else {
		isAdult = false
	}
	return fullName, isAdult
}

func unnecessaryValues() (string, int, int) {
	return "string", 1, 2
}

func main() {
	fullName, isAdult := fullInfo("name", "surname", 13)
	fmt.Println(fullName, isAdult)

	unnecessaryValues()
	_, value, _ := unnecessaryValues() // _ используется для пропуска значений

	fmt.Println(value)

}
