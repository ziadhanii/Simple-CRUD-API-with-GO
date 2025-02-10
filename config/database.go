package config

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "sqlserver://localhost:1433?database=HotelDB&trusted_connection=yes"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	DB = database
	log.Println("✅ Connected to SQL Server")
}
