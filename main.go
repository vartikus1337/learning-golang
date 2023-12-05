package main

import (
	"fmt"
)

func main() {
	// First: 10 numbers for array
	// Second: 3 pair for change index numbers in array
	var (
		to        uint8
		from      uint8
		workArray [10]uint8
	)
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&to)
		fmt.Scan(&from)
		workArray[to], workArray[from] = workArray[from], workArray[to]
	}
	for _, value := range workArray {
		fmt.Print(value, " ")
	}
}
