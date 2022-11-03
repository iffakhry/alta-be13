package main

import "fmt"

func findMaxMin(data []int) (max int, min int) {
	max = data[0]
	min = data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
		if data[i] < min {
			min = data[i]
		}
	}
	return
}

func main() {
	numbers := []int{10, 4, 7, 9, 8, 2, 1, 20, 50, 23, 45}
	nilaiMax, nilaiMin := findMaxMin(numbers)
	fmt.Println("max:", nilaiMax)
	fmt.Println("min:", nilaiMin)
}
