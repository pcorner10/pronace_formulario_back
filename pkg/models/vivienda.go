package models

import (
	"time"

	"gorm.io/gorm"
)

type Vivienda struct {
	gorm.Model
	Electricidad          bool         `gorm:"not null"`
	FuenteAgua            string       `gorm:"not null"`
	Gas                   bool         `gorm:"not null"`
	BanoInterno           bool         `gorm:"not null"`
	ExtraccionAguasNegras string       `gorm:"not null"`
	MaterialTecho         string       `gorm:"not null"`
	MaterialParades       string       `gorm:"not null"`
	MaterialPiso          string       `gorm:"not null"`
	NoHabitaciones        int          `gorm:"not null"`
	OloresDesagradables   bool         `gorm:"not null"`
	DescripcionOlores     string       `gorm:"not null"`
	Zona                  []Zona       `gorm:"foreignKey:ViviendaId;references:ID"`
	Habitantes            []Habitantes `gorm:"foreignKey:ViviendaId;references:ID"`
}

type Zona struct {
	gorm.Model
	Calle     string `gorm:"not null"`
	Localidad string `gorm:"not null"`
	Manzana   string `gorm:"not null"`
	Municipio string `gorm:"not null"`
	Estado    string `gorm:"not null"`
	Ageb      string `gorm:"not null"`
}

type AguaPotable struct {
	gorm.Model
	ViviendaId    uint   `gorm:"not null"`
	FuenteAgua    string `gorm:"not null"`
	TiempoConsumo int    `gorm:"not null"`
}

type ProblemasSalud struct {
	gorm.Model
	ViviendaId       uint   `gorm:"not null"`
	Ultimos12Meses   bool   `gorm:"not null"`
	ProblemaReferido string `gorm:"not null"`
	ProblemaRecabado string `gorm:"not null"`
	CIE10            string `gorm:"not null"`
}

type Fallecimientos struct {
	gorm.Model
	ViviendaId         uint      `gorm:"not null"`
	Sexo               string    `gorm:"not null"`
	EdadFallecimienti  int       `gorm:"not null"`
	FechaFallecimiento time.Time `gorm:"not null"`
	CausaReferida      string    `gorm:"not null"`
	CausaRecabada      string    `gorm:"not null"`
	CIE10              string    `gorm:"not null"`
}

type Cancer struct {
	gorm.Model
	ViviendaId       uint      `gorm:"not null"`
	Sexo             string    `gorm:"not null"`
	EdadDiagnostico  int       `gorm:"not null"`
	FechaDiagnostico time.Time `gorm:"not null"`
	CancerReferido   string    `gorm:"not null"`
	CancerRecabado   string    `gorm:"not null"`
	CIE10            string    `gorm:"not null"`
}

type Embarazo struct {
	gorm.Model
	ViviendaId           uint   `gorm:"not null"`
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
	ViviendaId   uint      `gorm:"not null"`
	FechaPerdida time.Time `gorm:"not null"`
	Trimestre    int       `gorm:"not null"`
	NoAplica     bool      `gorm:"not null"`
}

type Parto struct {
	gorm.Model
	ViviendaId           uint      `gorm:"not null"`
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
	ViviendaId uint   `gorm:"not null"`
	Referida   string `gorm:"not null"`
	Recabada   string `gorm:"not null"`
	Congenita  bool   `gorm:"not null"`
	CIE10      string `gorm:"not null"`
}

type Tratamiento struct {
	gorm.Model
	ViviendaId  uint   `gorm:"not null"`
	Atencion    string `gorm:"not null"`
	Ciudad      string `gorm:"not null"`
	Razon       string `gorm:"not null"`
	Certificado bool   `gorm:"not null"`
	NoAplica    bool   `gorm:"not null"`
}

type Formacos struct {
	gorm.Model
	ViviendaId     uint   `gorm:"not null"`
	Formaco        string `gorm:"not null"`
	Prescrito      bool   `gorm:"not null"`
	MotivoReferido string `gorm:"not null"`
	MotivoRecabado string `gorm:"not null"`
	Meses          int    `gorm:"not null"`
	CIE10          string `gorm:"not null"`
}
