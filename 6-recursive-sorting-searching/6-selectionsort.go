package main

import "fmt"

func selectionSort(elements []int) []int {
	n := len(elements)
	for k := 0; k < n; k++ {
		minimal := k
		for j := k + 1; j < n; j++ {
			if elements[j] < elements[minimal] {
				minimal = j
			}
		}
		elements[k], elements[minimal] = elements[minimal], elements[k]
	}
	return elements
}

func main() {
	var data = []int{5, 2, 7, 9, 10, 15, 2, 11, 8, 2}
	fmt.Println("sebelum:", data)
	result := selectionSort(data)
	fmt.Println("sesudah:", result)
}
