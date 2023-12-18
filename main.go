// Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера
//  только если оно отличается от того, что пришло ранее.

// Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять
//  значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)

// Функция должна называться removeDuplicates()

package main

import (
	"fmt"
)

func fillStream(c chan string, s ...string) {
	for _, s := range s {
		c <- s
	}
	close(c)
}

func anotherFillInMainFunc() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
}

func seeInChan(c chan string) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go fillStream(inputStream, "a", "b", "d", "b", "b", "d")
	go removeDuplicates(inputStream, outputStream)
	seeInChan(outputStream)
}

// Не увидел что при вводе они идут отсортированными (у других под это решение)
// Поэтому проверял на дупликацию элементов которые уже прошли
// Получается перевыполнил)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var wall []string
	for value := range inputStream {
		if contains(wall, value) == false {
			wall = append(wall, value)
		}
	}
	for _, value := range wall {
		outputStream <- value
	}
	close(outputStream)
}

func contains(wall []string, value string) bool {
	for _, v := range wall {
		if v == value {
			return true
		}
	}
	return false
}

// Как другие работают с потоками:

func removeDuplicates2(in chan string, out chan string) {
	var buf string
	for v := range in {
		if buf != v {
			out <- v
		}
		buf = v
	}
	close(out)
}

func removeDuplicates3(inputStream, outputStream chan string) {
	var prev string
	for v := range inputStream {
		if v != prev {
			outputStream <- v
			prev = v
		}
	}
	close(outputStream)
}
