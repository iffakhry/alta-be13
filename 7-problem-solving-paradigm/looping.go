package main

import "fmt"

func main() {
	// var count = 10
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	/*
		0 tidak sama dengan 5 = true
			1 tidak sama dengan 5 =true
			2 tidak sama dengan 5 =true
			3 tidak sama dengan 5 =true
			4 tidak sama dengan 5 =true
			5 tidak sama dengan 5 =false
	*/
	// var res bool = true
	// temp := !res
	// i := 0
	// for i != count {
	// 	fmt.Println(i)
	// 	i++
	// }

	// if i != count {
	// 	fmt.Println("tidak sama")
	// }
	numbers := []int{1, 2, 4, 7, 8, 9, 10, 20, 23, 45, 50}
	for i := 1; i < len(numbers); i += 2 {
		fmt.Println("index", i, "value", numbers[i])
	}

}
