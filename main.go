package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	duration := time.Duration(rand.Intn(4000)) * time.Millisecond
	fmt.Printf("gopher %d sleep for %v\n", id, duration)
	time.Sleep(duration)
	c <- id
}

func main() {
	timeout := time.After(2 * time.Second)

	c := make(chan int, 5)

	/**
	Горутины для гоферов нужно создать заранее. Если делать это в for вместе с select, то select будет
	блокировать дальнейшее исполнение и создание cледующей горутины
	**/

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		select { // Оператор select
		case gopherID := <-c: // Ждет, когда проснется гофер
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout: // Ждет окончания времени
			fmt.Println("my patience ran out")

			return // Сдается и возвращается
		}
	}
	// Select default
	tick1 := time.After(time.Second)
	tick2 := time.After(time.Second * 2)
	select {
	case <-tick1:
		fmt.Println("Получено значение из первого канала")
	case <-tick2:
		fmt.Println("Получено значение из второго канала")
	// Блок default выполнится раньше блока case - 1 секунда слишком много для Go // Если убрать default, то выполнится case 2
	default: 
		fmt.Println("Действие по умолчанию")
	}
}
