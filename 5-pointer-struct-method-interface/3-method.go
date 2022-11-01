package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       uint
}

type account struct {
	Nama string
}

//function
func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

//method
func (emp Employee) fullName() string {
	return emp.FirstName + " " + emp.LastName
}

func (emp Employee) tambahUmur(digit int) uint {
	hasil := emp.Age + uint(digit)
	return hasil
}

func main() {
	data := Employee{
		FirstName: "Ross",
		LastName:  "Geller",
		Age:       20,
	}

	fmt.Println(fullName(data.FirstName, data.LastName))
	//method
	fmt.Println("method", data.fullName())
	result := data.tambahUmur(5)
	fmt.Println("umur sekarang:", result)

	// dataAccount := account{
	// 	Nama: "budi",
	// }
	// dataAccount.tambahUmur(5)
}
