package app

import "pronaces_back/pkg/domain"

type userService struct {
	userDB domain.UserDB
}

func NewUserService(userDB domain.UserDB) *userService {
	return &userService{
		userDB: userDB,
	}
}

