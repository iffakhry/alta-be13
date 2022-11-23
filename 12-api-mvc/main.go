package main

import (
	"be13/mvc/config"
	"be13/mvc/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))

}
