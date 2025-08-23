package models

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateModelsSlice returns a slice of all models
func GetModels() []any {
	return []any{
		&Account{},
		&Customer{},
		// Add other models here
	}
}
