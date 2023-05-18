package database

import (
	"pronaces_back/internal/domain/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.Habitantes{}, &entity.Escolaridad{}, &entity.Ocupacion{},
		&entity.User{}, &entity.Vivienda{}, &entity.Zona{},
		&entity.AguaPotable{}, &entity.ProblemasSalud{}, &entity.Fallecimientos{},
		&entity.Cancer{}, &entity.Embarazo{}, &entity.PerdidaEmbarazo{},
		&entity.Parto{}, &entity.Discapacidad{}, &entity.Tratamiento{},
		&entity.Formacos{}, &entity.Ciudad{}, &entity.DonadorSangre{},
		&entity.RegistrosViviendas{},
	)

	if err != nil {
		panic(err)
	}

}
