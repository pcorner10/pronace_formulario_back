package domain

import "gorm.io/gorm"

type UserDB interface {
	Create(user User) error
	GetByEmail(email string) (*User, error)
	Update(user User) error
	Delete(userID uint64) error
}

type UserService interface {
	Create(user User) error
	GetByEmail(email string) (*User, error)
	Update(user User) error
	Delete(userID uint64) error
}

type User struct {
	gorm.Model
	UserName string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
}
