package main

import "fmt"

func main() {

	var name string = "any name"
	var age = 18 // or age := 18

	//  or

	// var (
	// 	name string = "any name"
	// 	age  int    = 18
	// )

	var symbol int32 = 'g'

	fmt.Println("Hello " + name)
	fmt.Println("Your age:", age, "\nsymbol (char)-<int>:", symbol, "\nsymbol (char)-<string>:", string(symbol))

	a := 10

	f := 2.0

	fmt.Println("rounded division:", a/3)
	fmt.Println("surplus division:", a%3)
	a++
	fmt.Println("++", a)
	a--
	fmt.Println("--", a)
	fmt.Println("float division", f/3)
}
