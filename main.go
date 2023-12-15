// Используем анонимные функции на практике.

// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
// которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0.
// Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

// 727178 -> 28



package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(n uint) uint {
		var strFormatted string
		str := strconv.FormatUint(uint64(n), 10)
		for index, c := range str {
			res, _ := strconv.Atoi(string(c))
			if str[index] == '0' {
				continue
			}
			if res%2 == 0 {
				strFormatted += string(str[index])
			}
		}
		if len(strFormatted) == 0 {
			strFormatted = "100"
		}
		res, _ := strconv.ParseUint(strFormatted, 10, 32)
		return uint(res)
	}
	fmt.Println(fn(1203))
}

//  Как использование может даже зайти:
//  defer func() {fmt.Println("defer в чём то")}(); panic()

// Другие решения:
func anotherSolution() {
	fn := func(X uint) uint {
		var x uint
		s := strconv.FormatUint(uint64(X), 10)
		for i := range s {
			if s[i]%2 == 0 && s[i] != '0' {
				x = x*10 + uint(s[i]-'0')
			}
		}
		if x == 0 {
			x = 100
		}
		return x
	}
	fmt.Println(fn(1203))
	fn2 := func(x uint) (y uint) {
		for k := uint(1); x > 0; x /= 10 {
			if d := x % 10; d%2 == 0 && d != 0 {
				y += k * d
				k *= 10
			}
		}
		if y == 0 {
			y = 100
		}
		return y
	}
	fmt.Println(fn2(1203))
}
