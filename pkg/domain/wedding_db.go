package domain

import "gorm.io/gorm"

type WeddingDB interface {
	FirstOrCreateUserWedding(userWedding UserWedding) (*UserWedding, error)
	ConfirmarInvitado(userWedding InvitadosConfirmados) (*InvitadosConfirmados, error)
}

type UserWedding struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserName  string `gorm:"uniqueIndex;not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Invitados int64  `gorm:"not null" json:"invitados"`
}

type InvitadosConfirmados struct {
	gorm.Model
	Invitados int64  `gorm:"not null" json:"invitados"`
	Nombre    string `gorm:"not null" json:"nombre"`
	Correo    string `gorm:"not null" json:"correo"`
	Telefono  string `gorm:"not null" json:"telefono"`
}
