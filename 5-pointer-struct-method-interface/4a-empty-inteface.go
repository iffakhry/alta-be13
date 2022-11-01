package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func cetak(data any) {
	fmt.Println("cetak", data)
}

func main() {
	// var i interface{}
	// describe(i)

	data1 := 42
	describe(data1)

	data2 := "hello"
	describe(data2)

	data3 := true
	describe(data3)
	cetak(data3)

}
