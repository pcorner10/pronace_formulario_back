package app

import (
	"pronaces_back/pkg/db/gorm"
	"pronaces_back/pkg/ihttp"

	"gorm.io/gorm"
)

func BindAuth(dbReq *gorm.DB) *ihttp.AuthHandler {
	userDB, err := db.NewUserGorm(dbReq)
	if err != nil {
		return nil
	}

	service := NewAuthService(userDB)
	handler := ihttp.NewAuthHandler(service)

	return handler
}

func BindUser(dbReq *gorm.DB) *ihttp.UserHandler {
	userDB, err := db.NewUserGorm(dbReq)
	if err != nil {
		return nil
	}

	service := NewUserService(userDB)
	handler := ihttp.NewUserHandler(service)

	return handler
}

func BindForm(dbReq *gorm.DB) *ihttp.FormHandler {
	formDB, err := db.NewFormGorm(dbReq)
	if err != nil {
		return nil
	}

	userDB, err := db.NewUserGorm(dbReq)
	if err != nil {
		return nil
	}

	zonaDB, err := db.NewZonaGorm(dbReq)
	if err != nil {
		return nil
	}

	service := NewFormService(formDB, userDB, zonaDB)
	handler := ihttp.NewFormHandler(service)

	return handler
}
