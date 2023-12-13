package main

import (
	"errors"
	"fmt"
)

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err!= nil {
        fmt.Println("Проверьте типы входных параметров")
    } else {
		fmt.Println(input * 3)
	}
	// test_error()

	err_test := error_test()
	if err_test != nil {
		fmt.Println(err_test.Error())
	}
}

func test_error() {
	err := errors.New("my error")
    fmt.Println("", err)
}

func error_test() error {
	i := 0
	fmt.Println("Выполнение функции error_test() 1")
	i++ // 1
	defer deferTest(i) // Попало 1
	i++ // 2
	fmt.Println("Выполнение функции error_test() 2")
	return errors.New("ошибка в error_test()")
}

func deferTest(i int) {
	fmt.Println("Финал error_test()")
	fmt.Println("i:", i)
}
