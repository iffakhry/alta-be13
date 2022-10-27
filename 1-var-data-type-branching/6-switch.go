package main

import "fmt"

func main() {
	var today int = 2
	var yesterday int = 1
	// switch today {
	// case 1:
	// 	fmt.Printf("Today is Monday")
	// case 2:
	// 	fmt.Printf("Today is Tuesday")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	switch {
	case today == 1:
		fmt.Printf("Today is Monday \n")
	case today == 2 && yesterday == 1:
		fmt.Printf("Today is Tuesday \n")
	default:
		fmt.Printf("Invalid Date \n")
	}

	value := 100
	switch value {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

}
