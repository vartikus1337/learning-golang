package main

import (
	"fmt"
	// "unicode/utf8"
)

func stringTest() {
	var s string = "Это строка"
	fmt.Printf("Длина строки: %d байт\n", len(s))
	s = s + " Новая строка"
	fmt.Printf("%v\n", s)
	for _, b := range s {
		fmt.Printf("%v ", b)
	}
	fmt.Println()
	for _, b := range s {
		fmt.Printf("%v ", string(b))
	}
}

func stringIsBytes() {
	bs := []byte("Это байтовый срез")

	fmt.Printf("Байтовый срез внутри: %v\n", bs)

	for i := range bs {
		if bs[i]%2 == 0 {
			bs[i] = bs[i] + 1
			continue
		}
		bs[i] = bs[i] - 1
	}

	fmt.Printf("Измененный байтовый срез в виде строки: %s", bs)
}

func runeString() {
	rs := []rune("Срез рун")
	fmt.Printf("Срез рун внутри: %v\n", string(rs))
	for i := range rs {
		fmt.Println(rs[i])
		if rs[i] == 'р' {
			rs[i] = '*'
		}
	}
	fmt.Println(string(rs))
}

// func main() {
// 	// len(возвращает количество байтов в строке)
// 	var en = "english"
// 	var ru = "русский"
// 	fmt.Println(len(en), len(ru))
// 	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))
// }

//  https://pkg.go.dev/strings
