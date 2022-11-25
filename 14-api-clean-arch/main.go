package main

import (
	"be13/clean/config"
	"be13/clean/utils/database/mysql"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
}
