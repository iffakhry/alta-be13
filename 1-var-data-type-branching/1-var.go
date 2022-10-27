package main

import "fmt"

func main() {
	var firstname string = "Fakhry"
	// fmt.Println(firstname)
	firstname = "budi"
	// fmt.Println(firstname)
	firstname = "adi"
	fmt.Println(firstname)

	var umur uint = 10
	fmt.Println(umur)

	var data1, data2, data3 string = "hello 1", "hello 2", "hello 3"
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)

	fmt.Println(data1, data2, data3)

	// hello
	lastname := "Nusantara"
	fmt.Println(lastname)

	angka := 100
	fmt.Println(angka)

	var isAktif bool = true
	fmt.Println(isAktif)

	var kodepos int
	fmt.Println(kodepos)

	// CONST
	const Pi float64 = 3.14
	fmt.Println(Pi)

	fmt.Printf("%d\n", angka)
	fmt.Printf("Nilai anda %d , Nama Anda : %s %s\n", angka, lastname, firstname)
	fmt.Println("Nama anda", angka, "nama anda", lastname, firstname)
	temp := fmt.Sprintf("Nilai anda: %d", angka)
	fmt.Println(temp)
}
