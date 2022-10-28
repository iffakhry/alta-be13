package main

import "fmt"

func main() {
	fmt.Println("selamat datang")
	CetakHalo()
	// CetakHalo()
	// CetakHalo()
	// CetakSalam()

	Sapa("budi")
	var data1 = "adi"
	Sapa(data1)
	WithParamInt(30)
	var user1 = "abu"
	var umur1 uint = 20
	WithParam2(user1, umur1)
	// pemanggilan fungsi dg return
	// var hasil = WithParamReturn(user1, umur1)
	// fmt.Println("tampilkan:", hasil)
	fmt.Println(WithParamReturn(user1, umur1))

	var nilai1 = 100
	skor1, status1 := HitungRaport(nilai1)
	fmt.Println("Hasil Raport User1: ", skor1, status1)
	var nilai2 = 75
	skor2, status2 := HitungRaport(nilai2)
	fmt.Println("Hasil Raport User2: ", skor2, status2)
}

func CetakHalo() {
	fmt.Println("hello")
	fmt.Println("halo")
	fmt.Println("konnijiwa")
	CetakSalam()
}

func CetakSalam() {
	fmt.Println("Sampai jumpa")
}

// dengan parameter
func Sapa(nama string) {
	fmt.Println("Haloo ", nama)
}

func WithParamInt(data int) {
	var result = data + 10
	fmt.Println("hasil:", result)
}

func WithParam2(nama string, umur uint) {
	fmt.Printf("Halo %s, umur anda: %d\n", nama, umur)
	fmt.Println()
}

func WithParamReturn(nama string, umur uint) string {
	// fmt.Printf("Dengan Return. Halo %s, umur anda: %d\n", nama, umur)
	// fmt.Println()
	var result = fmt.Sprintf("Dengan Return. Halo %s, umur anda: %d\n", nama, umur)

	return result
}

// return 1 = nilai dalam bentuk Huruf (A, B+ dll)
// return 2 = status kelulusan siswa
func HitungRaport(nilai int) (string, string) {
	var nilaiHuruf, status string
	if nilai > 85 {
		nilaiHuruf = "A"
		if nilai == 100 {
			status = "Cumlaude"
		} else {
			status = "Sangat memuaskan"
		}
	} else if nilai > 70 && nilai <= 85 {
		nilaiHuruf = "B+"
		status = "memuaskan"
	} else {
		nilaiHuruf = "B"
		status = "bagus"
	}

	return nilaiHuruf, status
}
