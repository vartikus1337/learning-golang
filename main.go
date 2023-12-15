package main

import "fmt"


type Person struct {
	name string
	age uint
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%d)", p.name, p.age)
}

type PersonErrorAge int

func (e PersonErrorAge) Error() string {
    return fmt.Sprintf("Ошибка в указании возраста: %d", e)
}

func SetAge(age int, p *Person) error {
	if age < 0 {
		return PersonErrorAge(age)
	}
	return nil
}

func main() {
	p := &Person{
		name: "John",
		age: 30,
	}
	fmt.Println(p)
    err := SetAge(-10, p)
	if err != nil {
		fmt.Printf("Ошибка обработана: %v\n", err)
	}
	if cError, ok := err.(PersonErrorAge); ok {
		fmt.Printf("Контекст: %d\n", cError)
	}
}