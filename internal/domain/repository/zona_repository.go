package repository

import "pronaces_back/internal/domain/entity"

type ZonaRepository interface {
	CreateZona(zona *entity.Zona) error
	GetZonaByID(zonaID uint) (*entity.Zona, error)
	GetZonaID(zona entity.Zona) (uint, error)
}
