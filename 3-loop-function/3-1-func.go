package main

import "fmt"

/*
buat fungsi untuk melakukan pencarian angka.
jika angka ketemu maka akan mengembalikan true. jika tidak ketemu maka mengembalikan false

Step:
1. baca tiap data nya
2. setiap data yang kita baca, kita lakukan pengecekan, apakah data yg kita baca itu sama dengan data yang dicari.
*/
func CekAngka(angka []int, cari int) bool {
	for i := 0; i < len(angka); i++ {
		// fmt.Println(angka[10])
		if angka[i] == cari {
			return true
		}
	}
	return false
}

func main() {
	var bil1 = []int{1, 2, 3, 4, 6, 7, 20, 90, 5, 40, 15}
	var bil2 = 10
	fmt.Println(CekAngka(bil1, bil2))
}
