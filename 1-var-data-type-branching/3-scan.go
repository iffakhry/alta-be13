package main

import "fmt"

// var (
// 	firstName, lastName, s string
// 	i                      int
// 	f                      float32
// 	input                  = "56.12 / 5212 / Go"
// 	format                 = "%f / %d / %s"
// )

func main() {
	var firstName string
	var amount int
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName)
	fmt.Println("halo ", firstName)
	fmt.Println("Please enter your bill: ")
	fmt.Scanln(&amount)
	fmt.Println("your amount ", amount)
	fmt.Println("total pay ", amount+1000)

	// fmt.Scanf("%s %s", &firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	// fmt.Sscanf(input, format, &f, &i, &s)
	// fmt.Println("From the string we read: ", f, i, s)
}
