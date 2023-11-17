package app

import "pronaces_back/pkg/domain"

type weddingService struct {
	db domain.WeddingDB
}

func NewWeddingService(db domain.WeddingDB) domain.WeddingService {
	return &weddingService{
		db: db,
	}
}

func (s *weddingService) FirstOrCreateUserWedding(userWedding domain.UserWedding) (*domain.UserWedding, error) {
	return s.db.FirstOrCreateUserWedding(userWedding)
}

func (s *weddingService) ConfirmarInvitado(userWedding domain.InvitadosConfirmados) (*domain.InvitadosConfirmados, error) {
	return s.db.ConfirmarInvitado(userWedding)
}