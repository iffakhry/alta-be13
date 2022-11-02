package main

import (
	"fmt"
	"math"
)

func checkPrimeBiasa(number int) bool {
	if number <= 1 {
		return false
	}
	// var count int
	// for i := 1; i <= number; i++ {
	// 	if number%i == 0 {
	// 		count++
	// 	}
	// }
	// if count > 2 {
	// 	return false
	// }
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkPrimeBiasa(13))
	fmt.Println(checkPrimeBiasa(9))
	// fmt.Println(checkPrimeBiasa(87178291199))
	fmt.Println(checkPrime(87178291199))
}
