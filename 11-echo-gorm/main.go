package main

// go get -u github.com/labstack/echo/v4
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"type:varchar(15)" json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Books    []Book
}

type Book struct {
	gorm.Model
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
}

// type User struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Name      string         `json:"name" form:"name"`
// 	Email     string         `json:"email" form:"email"`
// 	Password  string         `json:"password" form:"password"`
// }

// database connection
func InitDB() {

	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_be13_gorm?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

//proses database migration
func InitialMigration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Book{})
}

//func ini akan dijalankan sebelum func main
func init() {
	fmt.Println("func init berjalan")
	InitDB()
	InitialMigration()
}

func main() {
	fmt.Println("func main berjalan")
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))

}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user) // menangkap data yg dikirim dari req body dan disimpan ke variabel

	tx := DB.Create(&user) // proses insert data

	if tx.Error != nil {
		return c.JSON(http.StatusBadRequest, FailedResponse("Error Db Create "+tx.Error.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResponse("success create user"))
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	tx := DB.Find(&users)
	if tx.Error != nil {
		return c.JSON(http.StatusBadRequest, FailedResponse("error read data"))
	}

	return c.JSON(http.StatusOK, SuccessWithDataResponse("success read all users", users))
}

func FailedResponse(msg string) map[string]any {
	return map[string]any{
		"status":  "failed",
		"message": msg,
	}
}

func SuccessResponse(msg string) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": msg,
	}
}

func SuccessWithDataResponse(msg string, data any) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}
