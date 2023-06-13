package usecase

import (
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"
)

type FormUseCase struct {
	formRepository repository.FormRepository
}

func NewFormUseCase(formRepository repository.FormRepository) *FormUseCase {
	return &FormUseCase{
		formRepository: formRepository,
	}
}

func (u *FormUseCase) CreateForm(form entity.Table0) (*entity.Table0, error) {

	res := &entity.Table0{}
	// Buscar usuario por nombre de usuario
	res, err := u.formRepository.CreateForm(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}
