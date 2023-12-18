//  Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал. 

package main

import (
	"fmt"
)

func task(c chan int, N int) {
	c <- N+1
} 

func main() {
	channel := make(chan int, 1)
	task(channel, 3)
	fmt.Println(<-channel)
}
