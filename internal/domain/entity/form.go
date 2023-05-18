package entity

import (
	"time"

	"gorm.io/gorm"
)

type Habitantes struct {
	gorm.Model
	EncuestadorID    uint
	Encuestado       bool        `gorm:"not null"`
	ViviendaID       uint        `gorm:"not null"`
	Sexo             string      `gorm:"not null"`
	Edad             int         `gorm:"not null"`
	Fumador          bool        `gorm:"not null"`
	TiempoResidencia int         `gorm:"not null"`
	EscolaridadId    uint        `gorm:"not null"`
	OcupacionId      uint        `gorm:"not null"`
	Escolaridad      Escolaridad `gorm:"foreignKey:ID"`
	Ocupacion        Ocupacion   `gorm:"foreignKey:ID"`
}

type Escolaridad struct {
	gorm.Model
	Escolaridad    string `gorm:"not null"`
	Trunca         bool   `gorm:"not null"`
	NoEscolorizado bool   `gorm:"not null"`
	NoAplica       bool   `gorm:"not null"`
}

type Ocupacion struct {
	gorm.Model
	Ocupacion     string `gorm:"not null"`
	Remunerado    bool   `gorm:"not null"`
	HorasPorDia   int    `gorm:"not null"`
	DiasPorSemana int    `gorm:"not null"`
	NoAplica      bool   `gorm:"not null"`
}

type Vivienda struct {
	gorm.Model
	Electricidad          bool              `gorm:"not null"`
	FuenteAgua            string            `gorm:"not null"`
	Gas                   bool              `gorm:"not null"`
	BanoInterno           bool              `gorm:"not null"`
	ExtraccionAguasNegras string            `gorm:"not null"`
	MaterialTecho         string            `gorm:"not null"`
	MaterialParades       string            `gorm:"not null"`
	MaterialPiso          string            `gorm:"not null"`
	NoHabitaciones        int               `gorm:"not null"`
	OloresDesagradables   bool              `gorm:"not null"`
	DescripcionOlores     string            `gorm:"not null"`
	ZonaID                uint              `gorm:"not null"`
	Zona                  Zona              `gorm:"foreignKey:ViviendaID"`
	Habitantes            []Habitantes      `gorm:"foreignKey:ViviendaID" json:"-"`
	AguaPotableID         uint              `gorm:"not null"`
	AguaPotable           AguaPotable       `gorm:"foreignKey:AguaPotableID"`
	ProblemasSalud        []ProblemasSalud  `gorm:"foreignKey:ViviendaID" json:"-"`
	Fallecimientos        []Fallecimientos  `gorm:"foreignKey:ViviendaID" json:"-"`
	Cancer                []Cancer          `gorm:"foreignKey:ViviendaID" json:"-"`
	Embarazo              []Embarazo        `gorm:"foreignKey:ViviendaID" json:"-"`
	PerdidaEmbarazo       []PerdidaEmbarazo `gorm:"foreignKey:ViviendaID" json:"-"`
	Parto                 []Parto           `gorm:"foreignKey:ViviendaID" json:"-"`
	Discapacidad          []Discapacidad    `gorm:"foreignKey:ViviendaID" json:"-"`
	Tratamiento           []Tratamiento     `gorm:"foreignKey:ViviendaID" json:"-"`
	Farmacos              []Farmacos        `gorm:"foreignKey:ViviendaID" json:"-"`
	CiudadID              uint              `gorm:"not null"`
	Ciudad                Ciudad            `gorm:"foreignKey:CiudadID" json:"-"`
	DonadorSangre         []DonadorSangre   `gorm:"foreignKey:ViviendaID" json:"-"`
}

type Zona struct {
	gorm.Model
	ViviendaID uint   `gorm:"not null"`
	Calle      string `gorm:"not null"`
	Localidad  string `gorm:"not null"`
	Manzana    string `gorm:"not null"`
	Municipio  string `gorm:"not null"`
	Estado     string `gorm:"not null"`
	Ageb       string `gorm:"not null"`
}

type AguaPotable struct {
	gorm.Model
	ViviendaID    uint   `gorm:"not null"`
	FuenteAgua    string `gorm:"not null"`
	TiempoConsumo int    `gorm:"not null"`
}

type ProblemasSalud struct {
	gorm.Model
	ViviendaID       uint   `gorm:"not null"`
	Ultimos12Meses   bool   `gorm:"not null"`
	ProblemaReferido string `gorm:"not null"`
	ProblemaRecabado string `gorm:"not null"`
	CIE10            string `gorm:"not null"`
}

type Fallecimientos struct {
	gorm.Model
	ViviendaID         uint      `gorm:"not null"`
	Sexo               string    `gorm:"not null"`
	EdadFallecimienti  int       `gorm:"not null"`
	FechaFallecimiento time.Time `gorm:"not null"`
	CausaReferida      string    `gorm:"not null"`
	CausaRecabada      string    `gorm:"not null"`
	CIE10              string    `gorm:"not null"`
}

type Cancer struct {
	gorm.Model
	ViviendaID       uint      `gorm:"not null"`
	Sexo             string    `gorm:"not null"`
	EdadDiagnostico  int       `gorm:"not null"`
	FechaDiagnostico time.Time `gorm:"not null"`
	CancerReferido   string    `gorm:"not null"`
	CancerRecabado   string    `gorm:"not null"`
	CIE10            string    `gorm:"not null"`
}

type Embarazo struct {
	gorm.Model
	ViviendaID           uint   `gorm:"not null"`
	EdadEmbarazo         int    `gorm:"not null"`
	EnCurso              bool   `gorm:"not null"`
	PartoNatural         bool   `gorm:"not null"`
	Complicacion         bool   `gorm:"not null"`
	ComplicacionReferida string `gorm:"not null"`
	ComplicacionRecabada string `gorm:"not null"`
	CIE10                string `gorm:"not null"`
}

type PerdidaEmbarazo struct {
	gorm.Model
	ViviendaID   uint      `gorm:"not null"`
	FechaPerdida time.Time `gorm:"not null"`
	Trimestre    int       `gorm:"not null"`
	NoAplica     bool      `gorm:"not null"`
}

type Parto struct {
	gorm.Model
	ViviendaID           uint      `gorm:"not null"`
	FechaParto           time.Time `gorm:"not null"`
	BajoPeso             bool      `gorm:"not null"`
	Prematura            bool      `gorm:"not null"`
	Malformacion         bool      `gorm:"not null"`
	MalformacionReferida string    `gorm:"not null"`
	MalformacionRecabada string    `gorm:"not null"`
	CIE10                string    `gorm:"not null"`
}

type Discapacidad struct {
	gorm.Model
	ViviendaID uint   `gorm:"not null"`
	Referida   string `gorm:"not null"`
	Recabada   string `gorm:"not null"`
	Congenita  bool   `gorm:"not null"`
	CIE10      string `gorm:"not null"`
}

type Tratamiento struct {
	gorm.Model
	ViviendaID  uint   `gorm:"not null"`
	Atencion    string `gorm:"not null"`
	Ciudad      string `gorm:"not null"`
	Razon       string `gorm:"not null"`
	Certificado bool   `gorm:"not null"`
	NoAplica    bool   `gorm:"not null"`
}

type Farmacos struct {
	gorm.Model
	ViviendaID     uint   `gorm:"not null"`
	Formaco        string `gorm:"not null"`
	Prescrito      bool   `gorm:"not null"`
	MotivoReferido string `gorm:"not null"`
	MotivoRecabado string `gorm:"not null"`
	Meses          int    `gorm:"not null"`
	CIE10          string `gorm:"not null"`
}

type Ciudad struct {
	gorm.Model
	ViviendaID    uint   `gorm:"not null"`
	Problema      bool   `gorm:"not null"`
	Descripcion   string `gorm:"not null"`
	Contaminacion bool   `gorm:"not null"`
	Tipo          string `gorm:"not null"`
}

type DonadorSangre struct {
	gorm.Model
	ViviendaID uint `gorm:"not null"`
	Donador    bool `gorm:"not null"`
	Antiguedad int  `gorm:"not null"`
	NoAplica   bool `gorm:"not null"`
}

type RegistrosViviendas struct {
	gorm.Model
	ViviendaID      uint `gorm:"not null"`
	Manzana         int  `gorm:"not null"`
	Atendio         bool `gorm:"not null"`
	Deshabitada     bool `gorm:"not null"`
	Baldio          bool `gorm:"not null"`
	Local           bool `gorm:"not null"`
	EnConstrucccion bool `gorm:"not null"`
	Otro            bool `gorm:"not null"`
}
