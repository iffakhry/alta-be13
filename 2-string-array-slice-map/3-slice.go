package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Obtaining a slice from an array `array`
	// array[low:high]
	var primes = [10]int{2, 3, 5, 7, 11}

	// Creating a slice from the array
	var part_primes []int = primes[1:4]

	// part_primes = append(part_primes, 10000)
	// menambah data ke slice akan menambah data ke array juga

	fmt.Println(reflect.ValueOf(part_primes).Kind())
	fmt.Println(part_primes)

	var data2 []int
	if data2 == nil {
		fmt.Println("data2 kosong")
	}
	// var data1 = []string{"ani", "budi", "cindi", "dodi"}
	data1 := []string{"ani", "budi", "cindi", "dodi"}
	fmt.Println(data1)
	fmt.Println(len(data1))
	data1 = append(data1, "edi")
	data1 = append(data1, "frody")
	fmt.Println(data1)
	fmt.Println(len(data1))

}
