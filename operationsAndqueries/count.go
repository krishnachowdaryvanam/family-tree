package operationsAndqueries

import (
	"family-tree/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

// CountSons returns the number of sons for a given person.
func CountSons(db *gorm.DB, name string) (int, error) {
	var person models.Person
	if err := db.Where("username = ?", name).First(&person).Error; err != nil {
		return 0, err
	}

	var count int
	if err := db.Model(&models.Relationship{}).
		Where("related_person_id = ? AND relationship_type = ?", person.ID, models.Son).
		Count(&count).Error; err != nil {
		// Check if the error is due to relationship type not found
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}

// Function to count daughters of a person
func CountDaughters(db *gorm.DB, name string) (int, error) {
	var count int
	var person models.Person

	if err := db.Where("username = ?", name).First(&person).Error; err != nil {
		return 0, err
	}

	// Count the number of daughters for the person
	if err := db.Model(&models.Relationship{}).Where("related_person_id = ? AND relationship_type = ?", person.ID, models.Daughter).Count(&count).Error; err != nil {
		// Check if the error is due to relationship type not found
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}

// Function to count wives of a person
func CountWives(db *gorm.DB, name string) (int, error) {
	var count int
	var person models.Person

	if err := db.Where("username = ?", name).First(&person).Error; err != nil {
		return 0, err
	}

	if err := db.Model(&models.Relationship{}).Where("person_id = ? AND relationship_type = ?", person.ID, models.Wife).Count(&count).Error; err != nil {
		// Check if the error is due to relationship type not found
		if err == gorm.ErrRecordNotFound {
			return 0, nil // No daughters found, return count as 0
		}
		return 0, err
	}

	return count, nil
}

// FatherOf returns the name of the father for a given person.
func FatherOf(db *gorm.DB, name string) (string, error) {
	var person models.Person

	if err := db.Where("username = ?", name).First(&person).Error; err != nil {
		return "", err
	}

	var fatherRelationship models.Relationship
	if err := db.Where("related_person_id = ? AND relationship_type = ?", person.ID, "father").First(&fatherRelationship).Error; err != nil {
		// If the relationship is not found, return an error
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("father not found for %s", name)
		}
		return "", err
	}

	var father models.Person
	if err := db.Where("id = ?", fatherRelationship.PersonID).First(&father).Error; err != nil {
		return "", err
	}

	return father.Username, nil
}
