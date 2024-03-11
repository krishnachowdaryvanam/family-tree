package operationsAndqueries

import (
	"family-tree/models"
	"fmt"

	"errors"

	"github.com/jinzhu/gorm"
)

func ConnectPersons(db *gorm.DB, person1Name, relationship, person2Name string) error {
	if person1Name == "" || person2Name == "" || relationship == "" {
		return errors.New("invalid input, please provide person1Name, relationship, and person2Name")
	}

	// Find IDs of person1Name and person2Name
	var person1, person2 models.Person
	if err := db.Where("username = ?", person1Name).First(&person1).Error; err != nil {
		return fmt.Errorf("failed to find %s: %v", person1Name, err)
	}
	if err := db.Where("username = ?", person2Name).First(&person2).Error; err != nil {
		return fmt.Errorf("failed to find %s: %v", person2Name, err)
	}

	// Check if the relationship already exists
	var existingRelationship models.Relationship
	if err := db.Where("person_id = ? AND relationship_type = ? AND related_person_id = ?", person1.ID, relationship, person2.ID).First(&existingRelationship).Error; err == nil {
		return errors.New("relationship already exists")
	}

	// Create the relationship
	relationshipModel := models.Relationship{
		PersonID:         person1.ID,
		RelationshipType: relationship,
		RelatedPersonID:  person2.ID,
	}
	if err := db.Create(&relationshipModel).Error; err != nil {
		return err
	}

	return nil
}
