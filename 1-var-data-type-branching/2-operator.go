package main

import "fmt"

func main() {
	var bilangan1 int
	var bilangan2 int
	bilangan1 = 15
	// fmt.Println(bilangan1)
	bilangan2 = 10

	// operasi penjumlahan
	var bilangan3 = bilangan1 + bilangan2
	bilangan1 = 30
	fmt.Println("isi bilangan3 : ", bilangan3)

	var nama1 string = "Budi"
	var nama2 string = "Adi"
	var nama3 = nama1 + "-" + nama2
	fmt.Println("nama3: ", nama3)

	a := 10
	b := 3
	// c := a + b
	// a++
	// c = c + a
	var c float64
	c = float64(a) / float64(b) // 10.0 / 3.0 = 3.3333
	c = float64(a / b)          // 3.0
	fmt.Println(c)

	var nilai1 int = -10
	fmt.Println(nilai1)

}
