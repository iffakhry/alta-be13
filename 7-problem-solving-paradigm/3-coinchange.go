package main

import (
	"fmt"
	"sort"
)

func coinChange(uang int, coins []int) []int {
	sort.SliceStable(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	// 25 10 5 1
	result := []int{}
	for _, pecahan := range coins {
		if uang >= pecahan { //pengecekan bahwa uang harus lebih besar dari pecahan
			for uang >= pecahan {
				fmt.Println("uang yg didapat", pecahan)
				result = append(result, pecahan)
				uang = uang - pecahan
			}
		} else {
			continue
		}
	}
	return result
}

func main() {
	coinValue := []int{10, 25, 5, 1} // --> desc
	result := coinChange(100, coinValue)
	fmt.Println(result)
}

/*
coinValue := 10, 25, 5, 1
uang = 42

42 = 1,1,1,1,1,1,1 .... (42x) = 42 pecahan uang
42 = 5,5,5,5,5,5,5,5,1,1 = 10 pecahan uang
42 = 10,10,10,10,1,1 = 6 pecahan
42 = 25, 5,5,5,1,1 = 6 pecahan
...
42 = 25,10,5,1,1 = 5 pecahan

*/
