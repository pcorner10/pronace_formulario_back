package usecase

import (
	"errors"
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"
	"pronaces_back/internal/domain/service"
)

type AuthUseCase struct {
	userRepository repository.UserRepository
}

func NewAuthUseCase(userRepository repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepository: userRepository,
	}
}

func (u *AuthUseCase) Login(email, password string) (string, error) {
	// Buscar usuario por nombre de usuario
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	// Verificar contraseña
	if !service.ComparePasswords(user.Password, password) {
		return "", errors.New("invalid password")
	}

	// Generar token de autenticación
	token, err := service.GenerateToken(email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthUseCase) CreateUser(user *entity.User) error {
	return u.userRepository.Create(user)
}
