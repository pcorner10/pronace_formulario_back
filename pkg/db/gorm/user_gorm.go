package db

import (
	"fmt"
	"pronaces_back/pkg/domain"
)

// TODO
func (r *gormStore) LoginUser(email, password string) (*domain.ReponseLogin, error) {
	return nil, nil
}

// TODO
func (r *gormStore) RegisterUser(user *domain.User) error {
	return nil
}

func (r *gormStore) CreateUser(user domain.User) error {
	// implementar la lógica de crear un usuario con GORM
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) GetUserByEmail(email string) (*domain.User, error) {
	// implementar la lógica de obtener un usuario por ID con GORM
	fmt.Println("Estamos en GetUserByEmail de db")
	user := domain.User{}

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *gormStore) UpdateUser(user domain.User) error {
	// implementar la lógica de actualizar un usuario con GORM
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) DeleteUser(userID uint64) error {
	// implementar la lógica de eliminar un usuario con GORM
	err := r.db.Delete(&domain.User{}, userID).Error
	if err != nil {
		return err
	}

	return nil
}
