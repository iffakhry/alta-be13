package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	SERVER_PORT int16
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	// if _, exist := os.LookupEnv("SECRET"); !exist {
	// 	if err := godotenv.Load(".env"); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.SERVER_PORT = int16(cnv)
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	cnv, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.DB_PORT = cnv

	return &defaultConfig
}
