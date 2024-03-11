package operationsAndqueries

import (
	"family-tree/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func AddPerson(db *gorm.DB, person *models.Person) error {
	err := db.Create(&person).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

// GetPersonByName retrieves a person by their name from the database
func GetPersonByName(db *gorm.DB, name string) (*models.Person, error) {
	var person models.Person
	err := db.Where("name = ?", name).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// PersonExists checks if a person with the given name already exists in the database
func PersonExists(db *gorm.DB, name string) bool {
	var count int
	db.Model(&models.Person{}).Where("username = ?", name).Count(&count)
	return count > 0
}
