package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connection() *gorm.DB {
	dsn := "host=localhost port=5432 user=postgres password=admin dbname=url_shortener sslmode=disable timezone=Europe/Istanbul"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return nil
	}

	DB = connection
	return DB
}
