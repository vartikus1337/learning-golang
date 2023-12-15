// Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект,
// удовлетворяющий интерфейсу fmt.Stringer. Назовем его "Батарейка".

// Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.

// Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]:
// где пробелы - "опустошенная" емкость батареи, а X - "заряженная".

// Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр:
// 0 или 1 (порядок 0/1 случайный). Ваша задача считать эту строку любым возможным способом и создать на основе этой строки объект объявленного вами на первом этапе типа:
// надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".

// В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).

//  1000010011 -> [      XXXX]

package main

import (
	"fmt"
    "strings"
)

type Buttery struct {
    charge string        
}

func (b Buttery) String() string {
	var chargeFormatted string
	chargeFormatted += strings.Repeat(" ", strings.Count(b.charge, "0"))
	chargeFormatted += strings.Repeat("X", strings.Count(b.charge, "1"))
    return fmt.Sprintf("[%v]", chargeFormatted)
}

func main() {
	var charge string
	fmt.Scan(&charge)
	batteryForTest := Buttery{charge}
	fmt.Println(batteryForTest)
}

//  Another solution:

type battery string

func (b battery) String() string {
    return fmt.Sprintf("[%10s]", strings.Repeat("X", strings.Count(string(b), "1")))
}

func main1() {
    var batteryForTest battery
    fmt.Scan(&batteryForTest)
}