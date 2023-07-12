package app

import "pronaces_back/pkg/domain"

type userService struct {
	userDB domain.UserDB
}

func NewUserService(userDB domain.UserDB) domain.UserService {
	return &userService{
		userDB: userDB,
	}
}

func (s *userService) Create(user domain.User) error {
	return s.userDB.Create(user)
}

func (s *userService) GetByEmail(email string) (*domain.User, error) {
	return s.userDB.GetByEmail(email)
}

func (s *userService) Update(user domain.User) error {
	return s.userDB.Update(user)
}

func (s *userService) Delete(userID uint64) error {
	return s.userDB.Delete(userID)
}

