package main

import "fmt"

func main() {
	count := 3
	var result int
	for i := 1; i <= count; i++ {
		for j := 1; j <= count; j++ {
			for k := 1; k <= count; k++ {
				// fmt.Printf("i: %d --> j: %d --> k: %d\n", i, j, k)
				result += 1
			}
		}
	}
	fmt.Println(result)
}
