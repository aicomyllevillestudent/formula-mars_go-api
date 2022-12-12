package models

import (
	"fmt"
	"log"
	"os"
	"time"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database: ", DbDriver)
		log.Fatal("Connection error: ", err)
	} else {
		fmt.Println("Connection success: ", DbDriver)
	}

	postgresDB, _ := DB.DB()
	postgresDB.SetMaxIdleConns(25)
	postgresDB.SetMaxOpenConns(100)
	postgresDB.SetConnMaxLifetime(time.Minute)

	if err := DB.SetupJoinTable(&Race{}, "Drivers", &RaceDriver{}); err != nil {
		fmt.Println(err)
	}

	DB.AutoMigrate(&User{}, &Race{}, &Live{}, &Championship{}, &Driver{}, &Bet{})

}
