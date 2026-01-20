package database

import (
	"fmt"
	"log"
	"os"

	"afperdomo2/go/microservicios/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5433"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "devuser"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "devpassword123"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "intro_microservicios"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Adult{}, &models.Child{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database connected and migrated successfully")
}
