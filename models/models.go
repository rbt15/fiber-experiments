package models

import "gorm.io/gorm"

// Cat model
type Cat struct {
	gorm.Model

	Name  string `json:"name"`
	Breed string `json:"breed"`
}

// Specify table name for appropriate schema
func (Cat) TableName() string {
	return "fiber.cats"
}
