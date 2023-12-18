// https://stepik.org/lesson/345547/step/13?unit=329291.

package main

import (
	"fmt"
)

// В комментариях предложили для проверки использовать вот это:
func main1() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})
	ch1 <- 3
	// close(stop)
	r := calculator1(ch1, ch2, stop)

	fmt.Println(<-r)
}

// Ответ который работает там, но не работает тут, используя main2 работает и тут 
func calculator1(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	channel := make(chan int)
	go func() {
		defer close(channel)
		select {
		case num := <-firstChan:
			channel <- num * 2
		case num := <-secondChan:
			channel <- num * 3
		case <-stopChan:
		}
	}()
	return channel
}

// Как надо что бы работало
func main2() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})
	go func() {
		close(stop)
	}()
	
	// close(stop)
	r := calculator1(ch1, ch2, stop)

	fmt.Println(<-r)
}

// Ответ который не работает там, но работает тут, если использовать main2
func calculator2(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	channel := make(chan int)
	select {
	case num := <-firstChan:
		go func() {
			channel <- num * 2
			close(channel)
		}()
		return channel
	case num := <-secondChan:
		go func() {
			channel <- num * 3
			close(channel)
		}()
		return channel
	case <-stopChan:
		close(channel)
		return channel
	}
}


// Всё работает
func mainTest() {
	ch1, ch2 := make(chan int), make(chan int)
	// stop := make(chan struct{})
	go func() {
		ch1 <- 3
	}()
	
	// close(stop)
	check(ch1, ch2)
}


func check(ch1, ch2 <-chan int) {
	select {
	case <-ch1:
		fmt.Println("1")
	case <-ch2:
		fmt.Println("2")
	}
}