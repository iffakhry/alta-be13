package main

import "fmt"

func linierSearch(elements []int, x int) int {
	var count int
	for i := 0; i < len(elements); i++ {
		count++
		fmt.Println(count)
		if elements[i] == x {
			return i
		}

	}
	return -1
}

func main() {
	var data = []int{5, 2, 7, 9, 10, 15, 2, 11, 8, 2}
	var cari = 2

	lokasi := linierSearch(data, cari)
	if lokasi != -1 {
		fmt.Println("lokasi bilangan ditemukan di index ke:", lokasi)
	} else {
		fmt.Println("Bilangan yang anda cari tidak ditemukan")
	}
}
