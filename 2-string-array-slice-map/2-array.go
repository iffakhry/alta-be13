package main

import (
	"fmt"
	"reflect"
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

	primes := [...]int{2, 3, 3}
	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(len(primes))
	fmt.Println(primes[1])

}
