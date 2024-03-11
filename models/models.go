package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Username      string         `gorm:"unique"`
	Relationships []Relationship `gorm:"foreignkey:PersonID"`
}

type Relationship struct {
	gorm.Model
	PersonID         uint
	RelationshipType string
	RelatedPersonID  uint
}

const (
	Father   = "father"
	Son      = "son"
	Daughter = "daughter"
	Mother   = "mother"
	Wife     = "wife"
	Husband  = "husband"
	// Add other types as needed
)
