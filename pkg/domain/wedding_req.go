package domain

type WeddingService interface {
	FirstOrCreateUserWedding(userWedding UserWedding) (*UserWedding, error)
	ConfirmarInvitado(userWedding InvitadosConfirmados) (*InvitadosConfirmados, error)
}
