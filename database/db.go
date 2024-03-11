package database

import (
	"family-tree/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() (*gorm.DB, error) {
	connStr := "postgres://postgres:postgres@localhost:5432/familytree?sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return nil, err
	}

	db.AutoMigrate(&models.Person{}, &models.Relationship{})

	return db, nil
}
