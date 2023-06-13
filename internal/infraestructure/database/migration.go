package database

import (
	"pronaces_back/internal/domain/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{}, &entity.Zona{}, &entity.Table0{}, &entity.Table1{},
		&entity.Table2{}, &entity.Table3{}, &entity.Table4{}, &entity.Table5{},
		&entity.Table6{}, &entity.Table7{}, &entity.Table8{}, &entity.Table8_1{}, 
		&entity.Table9{}, &entity.Table9_1{}, &entity.Table10{}, &entity.Table10_1{}, 
		&entity.Table11{}, &entity.Table12{}, &entity.Table13{},
		&entity.Table14{},
	)

	if err != nil {
		panic(err)
	}

}
