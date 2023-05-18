package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
}
