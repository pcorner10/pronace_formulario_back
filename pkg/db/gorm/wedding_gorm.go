package db

import "pronaces_back/pkg/domain"

func (r *gormStore) FirstOrCreateUserWedding(userWedding domain.UserWedding) (*domain.UserWedding, error) {
	err := r.db.FirstOrCreate(&userWedding, userWedding).Error
	if err != nil {
		return nil, err
	}

	return &userWedding, nil
}

func (r *gormStore) ConfirmarInvitado(userWedding domain.InvitadosConfirmados) (*domain.InvitadosConfirmados, error) {
	err := r.db.Create(&userWedding).Error
	if err != nil {
		return nil, err
	}

	return &userWedding, nil
}
