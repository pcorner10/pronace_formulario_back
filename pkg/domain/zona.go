package domain

type ZonaDB interface {
	FirstOrCreateZona(zona Zona) (*Zona, error)
	CreateZona(zona Zona) (*Zona, error)
	GetZonaByID(zonaID uint) (*Zona, error)
	GetZonaID(zona Zona) (uint, error)
}

type ZonasService interface {
	CreateZona(zona Zona) error
	GetZonaByID(zonaID uint) (*Zona, error)
	GetZonaID(zona Zona) (uint, error)
}
