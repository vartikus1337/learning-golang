package main

import "fmt"

// Метод для класса
func (b *B) GetValueB() string {
	return string(b.b)
}

// Функция для класса
func getAFromB(b B) int {
	return b.A.a
}

func main() {
	// Инициализация структур
	// var t Test
	//  t := Test{}
	//  t := new(Test) - pointer

	//  Task struct
	c := new(B)

	c.A.a = 1
	c.b = 2

	fmt.Println(getAFromB(*c))

	a := A{1}

	b := B{a, 2}

	fmt.Println(getAFromB(b))

	fmt.Println(b.GetValueB())

	testStruct := Gun{true, 1, 1}
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.Shoot())

	// Math struct
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
}
