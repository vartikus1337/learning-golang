package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)
	/*
	*	Как работает?
	*	make(тип, на сколько заполнен, длина заполненности)
	*	как только дойдёт до 10 (в нашем случае), произойдёт пересоздание слайса с увеличением его ёмкости (на 10)
	 */
	fmt.Println(slice)
	users := []string{"Name 0", "Name 1", "Name 2", "Name 3"}
	fmt.Println(users[:3])
	fmt.Println(users[0:2])
	fmt.Println(users[2:])
	users = append(users, "Name 4")
	fmt.Println(users[4])
	fmt.Println("slice on practice")

	// Преобразование массива в слайс:
	array := [3]int{1, 2, 3}
	slice = array[:]

	// Практика
	fmt.Println("Практика")

	basicArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr:", basicArray)

	baseSlice := basicArray[5:8]
	fmt.Printf("Срез, основанный на базовом массиве длиной %d и емкостью %d: %v\n",
		len(baseSlice),
		cap(baseSlice), // Ёмкость он получил из длины от 5 индекса до конца массива
		baseSlice,      // если сделать [11]arr и добавить 11 элемент, то cap покажет 6
	)

	// Добавил поинтер чтоб ссылаться на наш изначальный массив
	pointer := fmt.Sprintf("%p", baseSlice)

	// Добавляем новый элемент
	fmt.Println("********************************")
	fmt.Println("baseSlice = append(baseSlice, 10)")
	baseSlice = append(baseSlice, 10)
	fmt.Println("arr:", basicArray) // Получилась подмена
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println("Наш ли это массив?", pointer == fmt.Sprintf("%p", baseSlice))

	// Добавим больше элементов, так что бы не хватило ёмкости
	fmt.Println("********************************")
	fmt.Println("baseSlice = append(baseSlice, 10, 11, 12)")
	baseSlice = append(baseSlice, 10, 11, 12)
	fmt.Println("arr:", basicArray) // Вместо 8: 10  потому что мы его подменили
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println("Наш ли это массив?", pointer == fmt.Sprintf("%p", baseSlice)) // Получили новый массив

	// Удаление элементов из slice
	fmt.Println("********************************")
	arrForChange := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr before:", arrForChange)
	arrForChange = append(arrForChange[0:2], arrForChange[3:]...)
	fmt.Println("arr after:", arrForChange)

	/*
	*	// Реализации:
	*
	*   // https://goforum.info/thread/kak-udalit-pervyi-element-massiva-v-golang
	*	// func removeByIndex(array []string, index int) []string {
	*	// 	return append(array[:index], array[index+1:]...)
	*	// }
	*
	*   // https://goforum.info/thread/kak-udalit-element-iz-sreza-golang
	*	// Если не важен порядок:
	*
	*	// func remove(slice []int, s int) []int {
	*	// 	return append(slice[:s], slice[s+1:]...)
	*	// }
	*
	*	// Важен порядок:
	*
	*	// func remove(s []int, i int) []int {
	*	// 	s[i] = s[len(s)-1]
	*	// 	return s[:len(s)-1]
	*	// }
	 */

	//  Функция копирования
	fmt.Println("********************************")
	a := []int{1, 2, 3}
	b := make([]int, 3, 3)
	n := copy(b, a) // Возвращает число копированных элементов

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("Скопировано %d элемента\n", n)
	
	//  Сам по себе срез это указатель при передаче в функцию изменится
}
