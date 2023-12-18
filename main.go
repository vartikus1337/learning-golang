//  Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

package main

import (
	"fmt"
)

func task(c chan int, N int) {
	c <- N + 1
}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}

}

func task2AnotherWay(c chan string, s string) {
	for _, res := range "     " {
		c <- s + string(res)
	}
}

func main() {
	channel := make(chan string, 5)
	task2(channel, "hi")
	fmt.Println(<-channel)
}
