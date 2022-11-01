package main

import "fmt"

// pass by value
func ubahNama(data string) string {
	fmt.Println("data awal:", data)
	data = "Budi"
	return data
}

// pass by reference
func ubahNamaPointer(data *string) string {
	fmt.Println("alamat data awal:", data)
	fmt.Println("data awal:", *data)
	*data = "Budi"
	return *data
}

func main() {
	var name string = "Adi"
	hasilValue := ubahNama(name)
	fmt.Println("hasilValue:", hasilValue)
	fmt.Println("name asli:", name)
	fmt.Println("==============")

	hasilValuePointer := ubahNamaPointer(&name)
	fmt.Println("hasilValuePointer:", hasilValuePointer)
	fmt.Println("name asli:", name)
}
