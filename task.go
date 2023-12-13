// На вход подается строка. Нужно определить, является ли она правильной или нет.
// Правильная строка начинается с заглавной буквы и заканчивается точкой.
// Если строка правильная - вывести Right иначе - вывести Wrong

package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"unicode"
)

func task1() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textRune := []rune(text) 
	if strings.HasSuffix(text, ".") && unicode.IsUpper(textRune[0]) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}