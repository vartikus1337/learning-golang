package main

import "fmt"

func main() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	delete(m, 1)
	fmt.Println(m)
	if value, inMap := m[1]; inMap {
		fmt.Println(value)
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}
