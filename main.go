package main

import (
	"fmt"
	"sync"
	"time"
)

func howWork() {
	// Timer
	t := time.NewTimer(time.Second) // создаем новый таймер, который сработает через 1 секунду
	go func() {
		<-t.C // C - канал, который должен вернуть значение через заданное время
	}()
	t.Stop() // но мы можем остановить таймер и раньше установленного времени

	t.Reset(time.Second * 2) // пока таймер не сработал, мы можем сбросить его, установив новый срок выполнения
	<-t.C
	// Ticker
	// func NewTicker(d Duration) *Ticker // создаем новый Ticker
	// func (t *Ticker) Stop() // останавливаем Ticker
}

func main() {
	<-work()
	/*
	 * тик-так
	 * тик-так
	 * тик-так
	 * тик-так
	 */

	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	wg := new(sync.WaitGroup)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, tick.C, wg)
	}

	wg.Wait()

	/*
	 * worker 1 выполнил работу
	 * worker 5 выполнил работу
	 * worker 3 выполнил работу
	 * worker 4 выполнил работу
	 * worker 2 выполнил работу
	 */
}

func worker(id int, limit <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	<-limit
	fmt.Printf("worker %d выполнил работу\n", id)
}

func work() <-chan struct{} {
	done := make(chan struct{}) // канал для синхронизации горутин

	go func() {
		defer close(done) // синхронизирующий канал будет закрыт, когда функция завершит свою работу

		stop := time.NewTimer(time.Second)

		tick := time.NewTicker(time.Millisecond * 200)
		defer tick.Stop() // освободим ресурсы, при завершении работы функции

		for {
			select {
			case <-stop.C:
				// stop - Timer, который через 1 секунду даст сигнал завершить работу
				fmt.Println("Опа")
				return
			case <-tick.C:
				// tick - Ticker, посылающий сигнал выполнить работу каждый 200 миллисекунд
				fmt.Println("тик-так")
			}
		}
	}()

	return done
}
