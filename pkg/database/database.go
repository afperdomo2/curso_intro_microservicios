package database

import (
	"fmt"
	"log"

	"afperdomo2/go/microservicios/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=devuser password=devpassword123 dbname=intro_microservicios port=5433 sslmode=disable"
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
