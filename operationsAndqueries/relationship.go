package operationsAndqueries

import (
	"family-tree/models"

	"github.com/jinzhu/gorm"
)

// CreateRelationship creates a new relationship in the database.
func AddRelationship(db *gorm.DB, relationship *models.Relationship) error {
	err := db.Create(relationship).Error
	return err
}

// RelationshipExists checks if a relationship with the given name already exists in the database
func RelationshipExists(db *gorm.DB, name string) bool {
	var count int
	db.Model(&models.Relationship{}).Where("relationship_type = ?", name).Count(&count)
	return count > 0
}
