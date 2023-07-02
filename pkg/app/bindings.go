package app

import (
	"pronaces_back/pkg/db"
	"pronaces_back/pkg/ihttp"

	"gorm.io/gorm"
)

func BindAuth(dbReq *gorm.DB) *ihttp.AuthHandler {
	userDB, err := db.NewUserGorm(dbReq)
	if err != nil {
		return nil, err
	}

	service := NewAuthService(userDB)
	handler := ihttp.NewAuthHandler(service)

	return handler, nil
}
