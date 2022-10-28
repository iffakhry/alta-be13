package main

import (
	"fmt"
)

func main() {
	var nilai [5]int
	nilai[0] = 100
	nilai[1] = 90

	fmt.Println(nilai[4])

	odd_numbers := [5]int{1, 3, 5, 7, 9}      // Intialized with values
	var even_numbers [5]int = [5]int{0, 2, 4} // Partial assignment

	fmt.Println(odd_numbers)
	fmt.Println(even_numbers)

	// primes := [...]int{2, 3, 3}
	// fmt.Println(reflect.ValueOf(primes).Kind())
	// fmt.Println(len(primes))
	// fmt.Println(primes[1])

	data1 := [5]int{1: 2, 4: 4, 2: 9}
	fmt.Println(data1)

	primes := [5]int{2, 3, 5, 8, 9}
	// technique 1
	fmt.Println(len(primes))
	// index < 5 --> index <= 4
	for index := 0; index < len(primes); index++ {
		// fmt.Println(index)
		fmt.Printf("Index: %d --> isi: %v \n", index, primes[index])
	}
	// // technique 2
	// for index, element := range primes {
	// 	fmt.Println(index, "=>", element)
	// }
	// for _, value := range primes {
	// 	fmt.Println(value)
	// }
	// // technique 3
	// index := 0
	// for range primes {
	// 	fmt.Println(primes[index])
	// 	index++
	// }

}
