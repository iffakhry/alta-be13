package main

import (
	"fmt"
)

func main() {
	kalimat := "hello"

	for _, karakter := range kalimat {
		// fmt.Println(karakter)
		fmt.Println(string(karakter))
	}

	// slicedata := []string{"h", "e", "l", "l", "o"}
	// for _, karakter := range slicedata {
	// 	fmt.Println(karakter)
	// }

}
