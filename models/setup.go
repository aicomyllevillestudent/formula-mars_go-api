package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	SSLMode := os.Getenv("SSL_MODE")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", DbHost, DbPort, DbUser, DbName, SSLMode, DbPassword)

	DB, err = gorm.Open(DbDriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database: ", DbDriver)
		log.Fatal("Connection error: ", err)
	} else {
		fmt.Println("Connection success: ", DbDriver)
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Race{})
	DB.AutoMigrate(&Championship{})
}
