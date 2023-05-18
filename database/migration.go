package database

import (
	"pronaces_back/formularios/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Habitantes{}, &models.Escolaridad{}, &models.Ocupacion{},
		&models.User{}, &models.Vivienda{}, &models.Zona{},
		&models.AguaPotable{}, &models.ProblemasSalud{}, &models.Fallecimientos{},
		&models.Cancer{}, &models.Embarazo{}, &models.PerdidaEmbarazo{},
		&models.Parto{}, &models.Discapacidad{}, &models.Tratamiento{},
		&models.Formacos{}, &models.Ciudad{}, &models.DonadorSangre{},
		&models.RegistrosViviendas{},
	)

	if err != nil {
		panic(err)
	}

}
