package orm

import (
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"

	"gorm.io/gorm"
)

type zonaRepository struct {
	db *gorm.DB
}

func NewZonaRepository(db *gorm.DB) repository.ZonaRepository {
	return &zonaRepository{
		db: db,
	}
}

func (r *zonaRepository) CreateZona(zona *entity.Zona) error {

	err := r.db.Create(&zona).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *zonaRepository) GetZonaByID(zonaID uint64) (*entity.Zona, error) {
	
	var zona entity.Zona
	err := r.db.Where("id = ?", zonaID).First(&zona).Error
	if err != nil {
		return nil, err
	}

	return &zona, nil
}

func (r *zonaRepository) GetZonaID(zona entity.Zona) (uint64, error) {
	
	var zonaID uint64

	query := `
		municipio = ? AND 
		localidad = ? AND 
		colonia = ? AND
		manzana = ? AND
		lote = ?`

	err := r.db.Where(query, 
		zona.Municipio, zona.Localidad, zona.Colonia, zona.Manzana, zona.Lote,
		).First(&zonaID).Error
	// if record not found, return record not found error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}
	
	if err != nil {
		return 0, err
	}

	return zonaID, nil
}
