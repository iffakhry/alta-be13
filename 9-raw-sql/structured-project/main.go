package main

/*
go get -u github.com/go-sql-driver/mysql
*/
import (
	"be13/sql/structured/config"
	"be13/sql/structured/controllers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnection := config.ConnectToDB()

	defer dbConnection.Close() // menutup koneksi

	//buat mekanisme menu
	fmt.Println("MENU:\n1. Baca data\n2. Tambah data")
	fmt.Println("Masukkan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			dataUser, errGetAll := controllers.GetAllUser(dbConnection) // memanggil fungsi di controllers
			if errGetAll != nil {                                       // cek apakah ada error dari getAllUser
				fmt.Println("error get all users")
			}
			// fmt.Println("data all", dataUser)
			for _, value := range dataUser { // membaca/menampilkan  seluruh data user yang telah ditampung di variable slice
				fmt.Printf("id: %s, nama: %s, email: %s, phone: %s\n", value.Id, value.Nama, value.Email, value.Phone)
			}
		}
	case 2:
		{
			// newUser := User{}
			// fmt.Println("masukkan id user")
			// fmt.Scanln(&newUser.Id)
			// fmt.Println("masukkan nama user")
			// fmt.Scanln(&newUser.Nama)
			// fmt.Println("masukkan Email user")
			// fmt.Scanln(&newUser.Email)
			// fmt.Println("masukkan Password user")
			// fmt.Scanln(&newUser.Password)
			// fmt.Println("masukkan Phone user")
			// fmt.Scanln(&newUser.Phone)
			// fmt.Println("masukkan DOmisili user")
			// fmt.Scanln(&newUser.Domisili)

			// var query = "INSERT INTO users(id, nama, email, password, phone, domisili) VALUES (?,?,?,?,?,?)"
			// statement, errPrepare := db.Prepare(query)
			// if errPrepare != nil {
			// 	log.Fatal("error prepare insert", errPrepare.Error())
			// }

			// result, errExec := statement.Exec(newUser.Id, newUser.Nama, newUser.Email, newUser.Password, newUser.Phone, newUser.Domisili)
			// if errExec != nil {
			// 	log.Fatal("error exec insert", errExec.Error())
			// } else {
			// 	row, _ := result.RowsAffected()
			// 	if row > 0 {
			// 		fmt.Println("Insert berhasil")
			// 	} else {
			// 		fmt.Println("Insert gagal")
			// 	}
			// }
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
