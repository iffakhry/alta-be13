package main

/*
go get -u github.com/go-sql-driver/mysql
*/
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       string
	Nama     string
	Email    string
	Password string
	Phone    string
	Domisili string
}

func main() {
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	var connectionString = "root:qwerty123@tcp(127.0.0.1:3306)/db_be13"
	db, err := sql.Open("mysql", connectionString) // membuka koneksi ke daatabase
	if err != nil {                                // pengecekan error yang terjadi ketika proses open connection
		log.Fatal("error open connection", err.Error())
	}

	errPing := db.Ping() // mengecek apakah apliasi masih terkoneksi ke database
	if errPing != nil {  //handling error ketika gagal konek ke db
		log.Fatal("error connect to db ", errPing.Error())
	} else {
		fmt.Println("koneksi berhasil")
	}

	defer db.Close() // menutup koneksi

	result, errSelect := db.Query("SELECT id, nama, email, password, phone, domisili FROM users") // proses menjalankana query SQL
	if errSelect != nil {                                                                         //handling error saat proses menjalankan query
		log.Fatal("error select ", errSelect.Error())
	}

	// var dataUser []User
	for result.Next() { // membaca tiap baris/row dari hasil query
		var userrow User                                                                                                         // membuat variabel penampung
		errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Password, &userrow.Phone, &userrow.Domisili) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
		if errScan != nil {                                                                                                      // handling ketika ada error pada saat proses scannign
			log.Fatal("error scan ", errScan.Error())
		}
		fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
	}

}
