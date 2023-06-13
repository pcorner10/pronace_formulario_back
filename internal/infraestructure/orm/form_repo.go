package orm

import (
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"

	"gorm.io/gorm"
)

type formRepository struct {
	db *gorm.DB
}

func NewFormRepository(db *gorm.DB) repository.FormRepository {
	return &formRepository{
		db: db,
	}
}

func (r *formRepository) CreateForm(table0 entity.Table0) (*entity.Table0, error) {

	err := r.db.Create(&table0).Error
	if err != nil {
		return nil, err
	}

	return &table0, nil
}
