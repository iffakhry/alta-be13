package main

import "fmt"

func logarithmic(n int) int {
	var result int = 0
	for n > 1 {
		n /= 2 // --> n = n/2
		result += 1
	}
	return result
}

/*
32
16
8
4
2
1
*/

func main() {
	fmt.Println(logarithmic(32))
}
