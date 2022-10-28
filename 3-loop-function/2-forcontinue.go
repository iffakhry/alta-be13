package main

import "fmt"

func main() {
	/*
		i=0 --> cetak 0
		i=1 --> continue
		i=2 --> cetak 2
		i=3 --> cetak 3
		i=4 --> break
	*/
	// for i := 0; i < 10; i++ {
	// 	if i == 1 {
	// 		continue
	// 	}
	// 	if i > 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	/*
		baris=0 kolom=0
		baris=0 kolom=1
		baris=0 kolom=2
		baris=0 kolom=3
		baris=0 kolom=4
		baris=1 kolom=0
		baris=1 kolom=1
		baris=1 kolom=2
		baris=1 kolom=3
		baris=1 kolom=4
		baris=2 kolom=0
	*/
	// N := 5
	// for baris := 0; baris < N; baris++ {
	// 	for kolom := 0; kolom < N; kolom++ {
	// 		fmt.Println("baris:", baris, "kolom:", kolom)
	// 	}
	// 	fmt.Println("selesai")
	// }

	/*
		i=0 j=0
		i=1 j=0
		i=1 j=1
		i=2 j=0
		i=2 j=1
		i=2 j=2
		i=3 j=0

	*/
	N := 5
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
