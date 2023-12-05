package main

import "fmt"

func main() {
	arr := [5]int{0, 1: 12, 23, 34, 45}

	for idx, elem := range arr {
		fmt.Println(idx, ":", elem)
	}
	// 
	// or
	// 
	// for idx := range arr {
	// 	fmt.Println(arr[idx]) // 
	// }
	//
    // or
    //
	// for _, elem := range arr {
	// 	fmt.Println(elem)
	// }
	// 
	// or
    // 
    // for i := 0; i < len(arr); i++ {
    //     fmt.Println(arr[i])
    // }
	// 
}
