package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	return s
}

func task2Deep() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	word := parseInput(input)
	fmt.Println(input)
	var reverse string
	for _, letter := range word {
		reverse = string(letter) + reverse 
	}
	if reverse == word {
        fmt.Println("Палиндром")
    } else {
        fmt.Println("Нет")
    }
}
