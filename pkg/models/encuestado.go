package models

import "gorm.io/gorm"

type Habitantes struct {
	gorm.Model
	EncuestadorID    uint
	Encuestado       bool   `gorm:"not null"`
	ViviendaId       uint   `gorm:"not null"`
	Sexo             string `gorm:"not null"`
	Edad             int    `gorm:"not null"`
	Fumador          bool   `gorm:"not null"`
	TiempoResidencia int    `gorm:"not null"`
	EscolaridadId    uint   `gorm:"not null"`
	OcupacionId      uint   `gorm:"not null"`
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

