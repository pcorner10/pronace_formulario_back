package db

import (
	"pronaces_back/pkg/domain"

	"gorm.io/gorm"
)

type LogsGorm struct {
	db *gorm.DB
}

func NewLogsGorm(db *gorm.DB) (*LogsGorm, error) {
	return &LogsGorm{
		db: db,
	}, nil
}

func (r *LogsGorm) Create(log *domain.Logs) error {
	// implementar la lógica de crear un log con GORM
	err := r.db.Create(&log).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *LogsGorm) FindAll() (*[]domain.Logs, error) {
	// implementar la lógica de obtener todos los logs con GORM
	logs := []domain.Logs{}

	err := r.db.Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return &logs, nil
}

func (r *LogsGorm) FindByID(id uint) (*domain.Logs, error) {
	// implementar la lógica de obtener un log por ID con GORM
	log := domain.Logs{}

	err := r.db.Where("id = ?", id).First(&log).Error
	if err != nil {
		return nil, err
	}

	return &log, nil
}

func (r *LogsGorm) FindByUserID(userID uint) (*[]domain.Logs, error) {
	// implementar la lógica de obtener un log por ID con GORM
	logs := []domain.Logs{}

	err := r.db.Where("user_id = ?", userID).Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return &logs, nil
}

func (r *LogsGorm) Update(log *domain.Logs) error {
	// implementar la lógica de actualizar un log con GORM
	err := r.db.Save(log).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *LogsGorm) Delete(log *domain.Logs) error {
	// implementar la lógica de eliminar un log con GORM
	err := r.db.Delete(&log).Error
	if err != nil {
		return err
	}

	return nil
}
