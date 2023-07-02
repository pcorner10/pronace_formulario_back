package db

import (
	"pronaces_back/pkg/domain"

	"gorm.io/gorm"
)

type zonaGorm struct {
	db *gorm.DB
}

func NewZonaGorm(db *gorm.DB) (domain.ZonaDB, error) {
	return &zonaGorm{
		db: db,
	}, nil
}

func (r *zonaGorm) FirstOrCreateZona(zona domain.Zona) (*domain.Zona, error) {
	err := r.db.FirstOrCreate(&zona, zona).Error
	if err != nil {
		return nil, err
	}

	return &zona, nil
}

func (r *zonaGorm) CreateZona(zona domain.Zona) (*domain.Zona, error) {
	err := r.db.Create(&zona).Error
	if err != nil {
		return nil, err
	}

	return &zona, nil
}

func (r *zonaGorm) GetZonaByID(zonaID uint) (*domain.Zona, error) {

	var zona domain.Zona
	err := r.db.Where("id = ?", zonaID).First(&zona).Error
	if err != nil {
		return nil, err
	}

	return &zona, nil
}

func (r *zonaGorm) GetZonaID(zona domain.Zona) (uint, error) {

	var zonaID uint

	query := `
		municipio = ? AND 
		localidad = ? AND 
		colonia = ? AND
		manzana = ? AND
		lote = ?`

	err := r.db.Where(query,
		zona.Municipio, zona.Localidad, zona.Colonia, zona.Manzana, zona.Lote,
	).First(&zona).Error
	// if record not found, return record not found error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	zonaID = zona.ID

	return zonaID, nil
}
