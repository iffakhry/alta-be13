package main

import "fmt"

func main() {
	// long declaration
	var salary = map[string]int{}
	fmt.Println(salary)

	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)
	salary_a["budi"] = 5000
	salary_a["umam"] = 10000
	salary_a["edi"] = 0

	fmt.Println(salary_a["umam"])

	// isExist akan bernilai true jika key "umam" ada di map
	// isExist akan bernilai false jika key "umam" tidak ada di map
	value, isExist := salary_a["edi"] // check key is exist
	fmt.Println("cek", value, isExist)

	// short declaration
	salary_b := map[string]int{}
	fmt.Println(salary_b)

	// using make
	var salary_c = make(map[string]int)
	salary_c["doe"] = 7000 // assign value
	fmt.Println(salary_c)

	var data1 = map[int]string{}
	data1[10] = "sepuluh"
	data1[50] = "lima puluh"
	data1[15] = "lima belas"
	fmt.Println(data1)

	for key, value := range salary_a { // loop over maps
		fmt.Println("->", key, value)
	}

}
