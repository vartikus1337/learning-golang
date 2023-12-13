// Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке.
// В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции.
// Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error)

package main

import (
	"errors"
	"fmt"
)

func task() {
	var (
		a, b int
	)
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		_, err := divide(a, b)
		if err != nil {
			fmt.Println("ошибка")
		}
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	}
	return a / b, nil
}
