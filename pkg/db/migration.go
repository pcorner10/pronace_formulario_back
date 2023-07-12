package db

import (
	"pronaces_back/pkg/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{}, &domain.Zona{}, &domain.Table0{}, &domain.Table1{},
		&domain.Table2{}, &domain.Table3{}, &domain.Table4{}, &domain.Table5{},
		&domain.Table6{}, &domain.Table7{}, &domain.Table8{}, &domain.Table8_1{},
		&domain.Table9{}, &domain.Table9_1{}, &domain.Table10{}, &domain.Table10_1{},
		&domain.Table11{}, &domain.Table12{}, &domain.Table13{},
		&domain.Table14{}, &domain.Logs{},
	)

	if err != nil {
		panic(err)
	}

}
