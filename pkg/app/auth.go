package app

import (
	"errors"
	"pronaces_back/pkg/domain"
	"pronaces_back/pkg/ihttp"
)

type authService struct {
	DBUser domain.UserDB
}

func NewAuthService(dbUser domain.UserDB) domain.AuthService {
	return &authService{
		DBUser: dbUser,
	}
}

func (u *authService) Login(email, password string) (*domain.ReponseLogin, error) {
	// Buscar usuario por nombre de usuario
	user, err := u.DBUser.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// Verificar contraseña
	if !ihttp.ComparePasswords(user.Password, password) {
		return nil, errors.New("invalid password")
	}

	// Generar token de autenticación
	token, err := ihttp.GenerateToken(email, user.Role)
	if err != nil {
		return nil, err
	}

	ReponseLogin := domain.ReponseLogin{
		Token:    token,
		UserName: user.Email,
		Role:     user.Role,
		Id:       user.ID,
	}

	return &ReponseLogin, nil
}

func (u *authService) Register(user *domain.User) error {
	// Verificar si el usuario ya existe
	_, err := u.DBUser.GetByEmail(user.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Encriptar contraseña
	Password, err := ihttp.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = Password

	return u.DBUser.Create(*user)
}
