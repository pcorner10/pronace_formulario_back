package db

import (
	"pronaces_back/pkg/domain"

	"gorm.io/gorm"
)

type UserGorm struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) (*UserGorm, error) {
	return &UserGorm{
		db: db,
	}, nil
}

func (r *UserGorm) Create(user domain.User) error {
	// implementar la lógica de crear un usuario con GORM
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserGorm) GetByEmail(email string) (*domain.User, error) {
	// implementar la lógica de obtener un usuario por ID con GORM
	user := &domain.User{}

	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserGorm) Update(user domain.User) error {
	// implementar la lógica de actualizar un usuario con GORM
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserGorm) Delete(userID uint64) error {
	// implementar la lógica de eliminar un usuario con GORM
	err := r.db.Delete(&domain.User{}, userID).Error
	if err != nil {
		return err
	}

	return nil
}
