package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := gorm.Open(postgres.Open(postgresInfo), &gorm.Config{})
	if err != nil {
		log.Println("Error Connecting to Database: ", err)
	}
	return db

}
