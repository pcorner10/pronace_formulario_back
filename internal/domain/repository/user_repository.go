package repository

import (
	"errors"
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/infraestructure/utility"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(email string) (*entity.User, error)
	ValidateUser(email string, password string) (*entity.User, error)
}

type userRepository struct {
	dbHandler *gorm.DB
}

func NewUserRepository(dbHandler *gorm.DB) UserRepository {
	return &userRepository{
		dbHandler: dbHandler,
	}
}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {

	return nil, nil
}

func (r *userRepository) ValidateUser(email string, password string) (*entity.User, error) {
	user, err := r.FindUserByEmail(email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	if !utility.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
