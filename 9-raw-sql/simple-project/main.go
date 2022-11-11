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
	var connectionString = "root:qwerty123@tcp(127.0.0.1:3306)/db_be13?parseTime=true"
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

	//buat mekanisme menu
	fmt.Println("MENU:\n1. Baca data\n2. Tambah data")
	fmt.Println("Masukkan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			result, errSelect := db.Query("SELECT id, nama, email, password, phone, domisili FROM users") // proses menjalankana query SQL
			if errSelect != nil {                                                                         //handling error saat proses menjalankan query
				log.Fatal("error select ", errSelect.Error())
			}

			var dataUser []User
			for result.Next() { // membaca tiap baris/row dari hasil query
				var userrow User                                                                                                         // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
				errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Password, &userrow.Phone, &userrow.Domisili) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
				if errScan != nil {                                                                                                      // handling ketika ada error pada saat proses scannign
					log.Fatal("error scan ", errScan.Error())
				}
				// fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
				dataUser = append(dataUser, userrow)
			}
			// fmt.Println("data all", dataUser)
			for _, value := range dataUser { // membaca seluruh data user yang telah ditampung di variable slice
				fmt.Printf("id: %s, nama: %s, email: %s, phone: %s\n", value.Id, value.Nama, value.Email, value.Phone)
			}
		}
	case 2:
		{
			newUser := User{}
			fmt.Println("masukkan id user")
			fmt.Scanln(&newUser.Id)
			fmt.Println("masukkan nama user")
			fmt.Scanln(&newUser.Nama)
			fmt.Println("masukkan Email user")
			fmt.Scanln(&newUser.Email)
			fmt.Println("masukkan Password user")
			fmt.Scanln(&newUser.Password)
			fmt.Println("masukkan Phone user")
			fmt.Scanln(&newUser.Phone)
			fmt.Println("masukkan DOmisili user")
			fmt.Scanln(&newUser.Domisili)

			var query = "INSERT INTO users(id, nama, email, password, phone, domisili) VALUES (?,?,?,?,?,?)"
			statement, errPrepare := db.Prepare(query)
			if errPrepare != nil {
				log.Fatal("error prepare insert", errPrepare.Error())
			}

			result, errExec := statement.Exec(newUser.Id, newUser.Nama, newUser.Email, newUser.Password, newUser.Phone, newUser.Domisili)
			if errExec != nil {
				log.Fatal("error exec insert", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("Insert berhasil")
				} else {
					fmt.Println("Insert gagal")
				}
			}
		}
	case 3:
		{
			fmt.Println("update")
		}

	case 4:
		{
			fmt.Println("delete")
		}

	case 5:
		{
			fmt.Println("baca data by id")
		}

	}

}
