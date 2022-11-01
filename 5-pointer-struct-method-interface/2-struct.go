package main

import "fmt"

type Person struct {
	Name   string
	Age    uint
	Height float32
	Weight float32
}

type Student struct {
	Orang   Person
	Kelas   string
	NIM     string
	Nilai   int
	JK      string
	isAktif bool
}

func main() {
	var orang1 Person
	orang1.Name = "Adi"
	orang1.Age = 20
	orang1.Height = 170.0
	orang1.Weight = 60.5

	fmt.Println(orang1)

	var orang2 Person
	orang2.Name = "Budi"
	orang2.Age = 30
	orang2.Height = 165
	orang2.Weight = 55
	fmt.Println(orang2)

	var siswa1 Student
	siswa1.Orang.Name = "Joko"
	siswa1.Orang.Age = 17
	siswa1.Kelas = "10 IPA 1"
	siswa1.NIM = "1235678"
	fmt.Println(siswa1)
	fmt.Println(siswa1.Orang.Name)
}
