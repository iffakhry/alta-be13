package main

import "fmt"

type User struct {
	Nama    string
	Email   string
	Telepon []string
	Alamats []Alamat
}

type Alamat struct {
	Jalan   string
	NoRumah string
	Kota    string
}

func main() {
	var namaInput string
	var data1 User
	fmt.Println("Masukkan Nama User:")
	fmt.Scanln(&namaInput)
	data1.Nama = namaInput
	data1.Email = "budi@mail.com"
	data1.Telepon = append(data1.Telepon, "081234")
	data1.Telepon = append(data1.Telepon, "085123")
	data1.Telepon = append(data1.Telepon, "08545667")

	data1.Alamats = append(data1.Alamats, Alamat{
		Jalan:   "Jalan Lurus",
		NoRumah: "A1",
		Kota:    "Jakarta",
	})

	data1.Alamats = append(data1.Alamats, Alamat{
		Jalan:   "Jalan Lurus Kanan",
		NoRumah: "B1",
		Kota:    "Aceh",
	})

	fmt.Println("data1", data1)
	fmt.Println("data1 nama", data1.Nama)
	fmt.Println("data1", data1.Telepon[0])
	for _, value := range data1.Telepon {
		fmt.Println("telepon: ", value)
	}
	for _, value := range data1.Alamats {
		fmt.Println("No Rumah: ", value.NoRumah)
	}

}
