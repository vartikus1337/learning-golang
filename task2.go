// На вход подается строка. Нужно определить, является ли она палиндромом.
// Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
// Все входные строчки в нижнем регистре.

// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях
// (например, "топот", "око", "заказ").

package main

import (
	"fmt"
)

func task2() {
	var (
		word    string
		reverse string
	)
	fmt.Scan(&word)
	for _, letter := range word {
		reverse = string(letter) + reverse 
	}
	if reverse == word {
	    fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
	fmt.Println(reverse)
}
