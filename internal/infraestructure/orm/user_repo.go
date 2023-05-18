package orm

import (
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *entity.User) error {
	// implementar la lógica de crear un usuario con GORM
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	// implementar la lógica de obtener un usuario por ID con GORM
	user := &entity.User{}

	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	// implementar la lógica de actualizar un usuario con GORM
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(userID uint64) error {
	// implementar la lógica de eliminar un usuario con GORM
	err := r.db.Delete(&entity.User{}, userID).Error
	if err != nil {
		return err
	}

	return nil
}
