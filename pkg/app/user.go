package app

import (
	"pronaces_back/pkg/domain"
)

func (s *surveyService) CreateUser(user domain.User) error {
	return s.DB.CreateUser(user)
}

func (s *surveyService) GetUserByEmail(email string) (*domain.User, error) {

	return s.DB.GetUserByEmail(email)
}

func (s *surveyService) UpdateUser(user domain.User) error {
	return s.DB.UpdateUser(user)
}

func (s *surveyService) DeleteUser(userID uint64) error {
	return s.DB.DeleteUser(userID)
}
