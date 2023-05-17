package seed

import (
	"pronaces_back/pkg/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}


}