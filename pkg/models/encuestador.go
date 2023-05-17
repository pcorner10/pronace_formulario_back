package models

import "gorm.io/gorm"

type Encuestador struct {
	gorm.Model
	UserName string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
}

