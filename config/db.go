package config

import (
	"goapi/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=admin1234 dbname=goapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // 👈 this fixes SASL auth issues
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.User{}, &models.BlacklistedToken{})

	log.Println("Database connected successfully! ")
	DB = db
}
