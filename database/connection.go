package database

import (
	"belajar-go-dasar/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// koneksi ke database
	// dsn := config.Env("CONNECTION")

	username := config.Env("DB_USERNAME")
	password := config.Env("DB_PASSWORD")
	host := config.Env("DB_HOST")
	port := config.Env("DB_PORT")
	databaseName := config.Env("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, databaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal connect database")
	} else {
		log.Println("connect success")
	}

	DB = db

	// migrate table
	Migration()

}
