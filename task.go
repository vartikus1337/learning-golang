// Представьте что вы работаете в большой компании где используется модульная архитектура.
// Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
// Вы же пишете функцию которая считывает две переменных типа string, а возвращает число типа int64 которое нужно получить сложением этих строк.

// Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
// Поэтому он решил подшутить над вами и подсунул вам свинью.Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

// Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа.
// Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes.

// %^80 hhhhh20&&&&nd
// 100

package main

import (
	"strconv"
	"unicode"
	//  Для друг решен
	"strings"
)

func main() {
	adding("%^80", "hhhhh20&&&&nd")
}

func adding(s1, s2 string) int64 {
	var (
		num1,
		num2 string
	)
	 
	for _, symbol1 := range s1 {
		if unicode.IsDigit(symbol1) {
			num1 += string(symbol1)
		}
	}

    for _, symbol2 := range s2 {
		if unicode.IsDigit(symbol2) {
			num2 += string(symbol2)
		}
	}

	res1 , _ := strconv.ParseInt(num1, 0, 64)
	res2 , _ := strconv.ParseInt(num2, 0, 64)
	return res1 + res2
}

//  Другое решение:

func adding2(s1, s2 string) int64 {

	
	s1 = strings.TrimFunc(s1, func(r rune) bool { return !unicode.IsDigit(r) })
	s2 = strings.TrimFunc(s2, func(r rune) bool { return !unicode.IsDigit(r) })

	i1, _ := strconv.ParseInt(s1, 10, 64)
	i2, _ := strconv.ParseInt(s2, 10, 64)

	return i1 + i2
}

