package formularios

import (
	"pronaces_back/formularios/models"

	"gorm.io/gorm"
)

type FormularioStorage interface {
	getUserByEmail(email string) (*models.User, error)
	createUser(user *models.User) (*models.User, error)
}

type FormularioService struct {
	db *gorm.DB
}

func NewFormularioService(db *gorm.DB) *FormularioService {
	return &FormularioService{db}
}

func (s *FormularioService) createUser(user *models.User) (*models.User, error) {
	err := s.db.Create(user).Error
	return user, err
}

func (s *FormularioService) getUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
