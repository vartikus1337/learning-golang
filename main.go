// Задание https://stepik.org/lesson/350788/step/10?unit=334666

package main

import "fmt"


func readTask() (interface{}, interface{}, interface{}) {
	return interface{}(1.123123), interface{}(2.7343), interface{}("*")
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
                                            // все полученные значения имеют тип пустого интерфейса
    if _, ok := value1.(float64); !ok {
        fmt.Printf("value=%v: %T", value1, value1)
        return 
    }
    if _, ok := value2.(float64); !ok {
        fmt.Printf("value=%v: %T", value2, value2)
        return 
    }
    if _, ok := operation.(string); !ok {
        fmt.Println("неизвестная операция")
        return
    }
    switch o := operation.(string); o {
        case "+":
            fmt.Printf("%.4f", value1.(float64) + value2.(float64))
        case "-":
            fmt.Printf("%.4f", value1.(float64) - value2.(float64))
        case "*":
            fmt.Printf("%.4f", value1.(float64) * value2.(float64))
        case "/":
            fmt.Printf("%.4f", value1.(float64) / value2.(float64))
        default:
            fmt.Println("неизвестная операция")
    }
}

func anotherSolution() {
	value1, value2, operation := readTask()

	v1, ok := value1.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}

	switch operation.(string) {
	case "+":
		fmt.Printf("%.4f", v1+v2)
	case "-":
		fmt.Printf("%.4f", v1-v2)
	case "*":
		fmt.Printf("%.4f", v1*v2)
	case "/":
		fmt.Printf("%.4f", v1/v2)
	default:
		fmt.Println("неизвестная операция")
	}
}
func test() {
	var i1 interface{} = true
	var i2 interface{} = 2.0120012010210
	var operation interface{} = "/"

	

	if _, ok := i1.(float64); !ok {
        fmt.Printf("value=%v: %T", i1, i1)
        return 
    }
    if f2, ok := i2.(float64); !ok {
        fmt.Printf("value=%v;%T", f2, f2)
        return 
    }
    if _, ok := operation.(string); !ok {
        fmt.Println("неизвестная операция")
        return
    }

	switch o := operation.(string); o {
		case "/":
			fmt.Printf("%.4f", i1.(float64) / i2.(float64))
		default:
			fmt.Println("неизвестная операция")
	}

}

