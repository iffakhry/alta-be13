package main

import (
	"fmt"
	"sort"
)

func linierSearch(elements []int, x int) int {
	var count int
	for i := 0; i < len(elements); i++ {
		count++
		fmt.Println("count", count)
		if elements[i] == x {
			return i
		}
	}
	return -1
}

func binarySearch(elements []int, x int) int {
	var kiri = 0
	var kanan = len(elements) - 1
	var count int
	for kiri <= kanan {
		count++
		fmt.Println("count", count)
		var tengah = (kiri + kanan) / 2
		if x < elements[tengah] {
			kanan = tengah - 1
		} else if x > elements[tengah] {
			kiri = tengah + 1
		} else if x == elements[tengah] {
			return tengah
		}
	}

	return -1
}

func main() {
	numbers := []int{1, 2, 4, 7, 8, 9, 10, 20, 23, 45, 50}
	sort.SliceStable(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println("sorted:", numbers)
	// fmt.Println(linierSearch(numbers, 2))
	fmt.Println(binarySearch(numbers, 2))
}
