package main

import "fmt"

type Akun struct {
	name string // Both non exported fields.
	age  int
}

func (P Akun) GetName() string {
	return P.name + " amazing!"
}

func (P *Akun) IncreaseAge() {
	P.age = P.age + 1
}

func main() {
	akunA := Akun{"John", 50}

	fmt.Printf("%v\n", akunA)
	fmt.Println(akunA.GetName())

	fmt.Println(akunA.age)
	akunA.IncreaseAge()
	fmt.Println(akunA.age)
}
