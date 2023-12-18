package main

import "fmt"

func main1() {
	done := make(chan struct{}) // Канал используется для синхронизации
	go myFunc1(done)
	<-done
}

func myFunc1(done chan struct{}) {
	fmt.Println("hello")
	close(done)
}

// Для синхронизации выведен шаблон:

func main() {
	<-myFunc()
}

func myFunc() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()
	return done
}

//  канал для синхронизации создается самой функцией myFunc(), полезная работа выполняется в отдельной горутине
// <-chan struct{}. Стрелка слева от ключевого слова chan означает, что возвращаемый канал предназначен только для чтения из него
//  вернуть из функции или передать в нее в качестве аргумента канал, предназначенный только для записи - chan<- struct{}

// Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.
// Функция work() ничего не принимает и не возвращает. https://stepik.org/lesson/345547/step/3?thread=solutions&unit=329291

func answer() {
	done := make(chan struct{})
	go func() {
		// work()
		close(done) // Лучше делать defer close(done) (на будущее)
	}()
	<-done
}

//  как синхронизировать не синхронизированное)
func explanation() {
	done := make(chan struct{}) // создаем канал
	go func(ch chan struct{}) { // создаем анонимную горутину
		// work()
		close(ch) // канал закроется после выполнения work()
	}(done) // передаем канал в анонимную горутину и запускаем ее
	<-done
}

func explanationReboot() {
	done := make(chan struct{})
	go func() {
		// work()
        close(done)
    }()
	<-done
}
