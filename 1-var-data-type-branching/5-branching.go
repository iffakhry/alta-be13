package main

import "fmt"

func main() {
	// var myAge = 5

	// if myAge == 5 { //false
	// 	fmt.Println("You too young")
	// }
	// if myAge == 17 {
	// 	fmt.Println("So Sweet")
	// }
	// if myAge > 17 || myAge < 30 {
	// 	fmt.Println("My Age is between 17 and 30")
	// }
	// if dadAge := 9; dadAge < myAge {
	// 	fmt.Println(dadAge)
	// }

	// number := 4
	// if number%2 == 0 {
	// 	fmt.Println("Genap")
	// 	fmt.Println("Yeay")
	// } else {
	// 	fmt.Println("Ganjil")
	// }

	hour := 18
	if hour < 12 {
		if hour == 10 {
			fmt.Println("olahraga")
		}
		fmt.Println("Selamat Pagi")
	} else if hour <= 18 {
		fmt.Println("Selamat Sore")
	} else if hour < 21 {
		fmt.Println("waktunya tidur")
	} else {
		fmt.Println("Selamat Malam")
	}

	if hour == 18 {
		fmt.Println("waktunya pulang ke rumah")
	}

}
