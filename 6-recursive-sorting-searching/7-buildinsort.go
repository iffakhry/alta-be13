package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"bc", "ab", "ba", "z", "u"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	var data = []int{5, 2, 7, 9, 10, 15, 2, 11, 8, 2}
	// sort.Slice(data, func(i, j int) bool {
	// 	return data[i] > data[j]
	// })
	sort.SliceStable(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	fmt.Println("data", data)

	/*
		data[i] < data[j] --> ascending
		data[i] > data[j] --> descending
	*/
	datastr := []string{"bc", "ab", "ba", "z", "u"}
	sort.SliceStable(datastr, func(i, j int) bool {
		return datastr[i] > datastr[j]
	})
	fmt.Println("data", datastr)

}
