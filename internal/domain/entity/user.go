package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
}
