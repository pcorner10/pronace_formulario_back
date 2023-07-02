package app

import (
	"pronaces_back/pkg/db"
	"pronaces_back/pkg/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateForm0(t *testing.T) {
	dbReq := db.GetDB()

	formDB, err := db.NewFormGorm(dbReq)
	assert.NoError(t, err, "error creating formDB")

	userDB, err := db.NewUserGorm(dbReq)
	assert.NoError(t, err, "error creating userDB")

	zonaDB, err := db.NewZonaGorm(dbReq)
	assert.NoError(t, err, "error creating zonaDB")

	service := NewFormService(formDB, userDB, zonaDB)

	tt := []struct {
		name string
		form domain.Form0Request
	}{
		{
			name: "User not found",
			form: domain.Form0Request{
				EncuestadorEmail: "princon0@ucol.com",
				Zona: domain.Zona{
					Municipio: "Colima",
				},
				Integrante: []domain.Form0{},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := service.CreateForm0(tc.form)
			assert.Error(t, err, "error should be returned")
		})
	}
}
