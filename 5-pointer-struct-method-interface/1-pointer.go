package main

import "fmt"

func main() {
	var a = 10

	fmt.Println("value a:", a)           // 10
	fmt.Println("memory address a:", &a) // 0xc00012a008

	var p *int = &a
	fmt.Println("value p:", p)             // 0xc00012a008
	fmt.Println("value from address:", *p) // 10

	var name string = "Adi"
	var namePointer *string = &name
	fmt.Println("value name", name)
	fmt.Println("address name", &name)
	fmt.Println("value pointer", namePointer)
	fmt.Println("value from address", *namePointer)
	*namePointer = "Budi"           // mutate
	fmt.Println("value name", name) //Budi
}
