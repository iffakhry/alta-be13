package main

import "fmt"

type Product struct {
	Nama  string
	Tipe  string
	Harga int
	Stock int
}

func main() {
	var products []Product
	products = append(products, Product{
		Nama:  "Baju anak",
		Tipe:  "Pakaian",
		Harga: 50000,
		Stock: 100,
	})

	var temp = Product{"Kemeja Polos", "Pakaian", 90000, 50}
	products = append(products, temp)

	// fmt.Println("Products", products)

	for i, value := range products {
		fmt.Println("Nama Product ke ", i, ":", value.Nama)
		fmt.Println("Harga: ", value.Harga)
	}

	for i := 0; i < len(products); i++ {
		fmt.Println("Nama Product ke ", i, ":", products[i].Nama)
		fmt.Println("Harga: ", products[i].Harga)
	}

	var mapproducts = map[string]Product{}
	mapproducts["budi"] = Product{
		Nama: "Celana Jeans",
	}
	mapproducts["adi"] = Product{
		Nama: "Jacket Bomber",
	}

	for key, v := range mapproducts {
		fmt.Println("key", key, "data:", v.Nama)
	}

}
