package main

import (
	"fmt"
)

type calculate interface {
	large() int
	keliling() int
	Sapa(nama string) string
}

type square struct {
	side int
}

type lingkaran struct {
	radius int
}

// func (s square) hitung() float64 {
// 	hasil := s.side * s.side
// 	return float64(hasil)
// }

func (s square) large() int {
	return s.side * s.side
}

func (s square) keliling() int {
	hasil := 4 * s.side
	return hasil
}

func (s square) Sapa(data string) string {
	return "halo " + data
}

func (l lingkaran) large() int {
	hasil := 3.14 * float64(l.radius) * float64(l.radius)
	return int(hasil)
}

func (l lingkaran) keliling() int {
	hasil := 4 * l.radius
	return hasil
}

func main() {
	var dimResult calculate
	dimResult = square{
		side: 10,
	}
	fmt.Println("large square :", dimResult.large())
	fmt.Println("keliling square :", dimResult.keliling())
	var dimResultLingkaran calculate
	dimResultLingkaran = lingkaran{
		radius: 10,
	}
	fmt.Println("large lingkaran :", dimResultLingkaran.large())

}
