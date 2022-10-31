package main

import (
	calc "fakhry/unittest/calculator/library"
	salary "fakhry/unittest/salary/library"
	"fmt"
)

/*
pertama, jalankan go mod init, untuk generate file go.mod
go mod init namaproject

NB nama project bebas

untuk menjalankan file test
go test

go test ./calculator/library/... -coverprofile=cover.out && go tool cover -html=cover.out
*/

func main() {
	result := calc.Penambahan(20, 50)
	fmt.Println("hasil : ", result)

	resultKurang := calc.Pengurangan(100, 40)
	fmt.Println("hasil pengurangan: ", resultKurang)

	resBagi := calc.Pembagian(10, 2)
	fmt.Println("hasil pembagian: ", resBagi)

	resGaji := salary.HitungGaji(10)
	fmt.Println("hasil gaji: ", resGaji)
}
