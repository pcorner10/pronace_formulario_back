package database

import (
	"pronaces_back/internal/domain/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{}, &entity.Vivienda{}, &entity.AguaPotable{},
		&entity.Fallecimientos{}, &entity.Cancer{}, &entity.Embarazo{},
		&entity.PerdidaEmbarazo{}, &entity.Parto{}, &entity.Discapacidad{},
		&entity.Tratamiento{}, &entity.Farmacos{}, &entity.Ciudad{},
		&entity.DonadorSangre{}, &entity.RegistrosViviendas{}, &entity.Zona{},
		&entity.ProblemasSalud{}, &entity.Habitantes{}, &entity.Escolaridad{},
		&entity.Ocupacion{},
	)

	if err != nil {
		panic(err)
	}

}
