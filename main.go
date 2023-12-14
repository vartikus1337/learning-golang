package main

import (
	"fmt"
	"strconv"
)

func ok() {
	age := 1337
	fmt.Println("Привет мне " + strconv.Itoa(age) + " лет")
	// или res, err := strconv.ParseInt(age, 10, 0)

	var f float64 = 1.3223
	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - точность (кол-во знаков после запятой)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	fmt.Println(strconv.FormatFloat(f, 'f', 2, 64)) // 1.01

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	var fb float64 = 2222 * 1023 * 245 * 2 * 52
	fmt.Println(strconv.FormatFloat(fb, 'e', -1, 64)) // 5.791874088e+10

	//  https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/

	var b bool = true
	fmt.Println(strconv.FormatBool(b))

	str1 := "10"
	str2 := "20"

	str1Int, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println(err)
	}

	str2Int, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str1Int - str2Int)

	float := "1.0000000012345678"
	result32, _ := strconv.ParseFloat(float, 32)
	result64, _ := strconv.ParseFloat(float, 64)
	fmt.Println("bitSize32:", result32)  // вывод 1 (не уместились)
	fmt.Println("bitSize64:", result64)  // вывод  1.0000000012345678

	str3 := "true"
	res, _ := strconv.ParseBool(str3)
	fmt.Printf("%T \n", res)
}
