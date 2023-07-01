package domain

type ZonaDB interface {
	CreateZona(zona Zona) error
	GetZonaByID(zonaID uint) (*Zona, error)
	GetZonaID(zona Zona) (uint, error)
}

type ZonasService interface {
	CreateZona(zona Zona) error
	GetZonaByID(zonaID uint) (*Zona, error)
	GetZonaID(zona Zona) (uint, error)
}
