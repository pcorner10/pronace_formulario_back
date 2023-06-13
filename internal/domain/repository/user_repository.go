// internal\domain\repository\user_repository.go

package repository

import "pronaces_back/internal/domain/entity"

// UserRepository define las operaciones para acceder y manipular los datos de los usuarios
type UserRepository interface {
	Create(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(userID uint64) error
}
