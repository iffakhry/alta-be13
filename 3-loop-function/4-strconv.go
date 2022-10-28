package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data1 int = 20
	var result1 = strconv.Itoa(data1)
	fmt.Println(result1)
	fmt.Println(string(result1[1]))

	// var resultInt, err = strconv.Atoi(string(result1[0]))
	var resultInt, err = strconv.Atoi("5000")
	fmt.Println(resultInt, err)
}
