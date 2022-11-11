package controllers

import (
	"be13/sql/structured/entities"
	"database/sql"
)

func GetAllUser(db *sql.DB) ([]entities.User, error) {
	result, errSelect := db.Query("SELECT id, nama, email, password, phone, domisili FROM users") // proses menjalankana query SQL
	if errSelect != nil {                                                                         //handling error saat proses menjalankan query
		// log.Fatal("error select ", errSelect.Error())
		return nil, errSelect
	}

	var dataUser []entities.User
	for result.Next() { // membaca tiap baris/row dari hasil query
		var userrow entities.User                                                                                                // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
		errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Password, &userrow.Phone, &userrow.Domisili) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
		if errScan != nil {                                                                                                      // handling ketika ada error pada saat proses scannign
			// log.Fatal("error scan ", errScan.Error())
			return nil, errScan
		}
		// fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
		dataUser = append(dataUser, userrow)
	}
	return dataUser, nil
}

func InsertData(db *sql.DB, newData entities.User) {

}
