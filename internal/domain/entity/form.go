package entity

import (
	"gorm.io/gorm"
)

type Zona struct {
	gorm.Model
	Municipio string `json:"municipio" gorm:"not null"`
	Localidad string `json:"localidad" gorm:"not null"`
	Colonia   string `json:"colonia" gorm:"not null"`
	Manzana   string `json:"manzana" gorm:"not null"`
	Lote      string `json:"lote" gorm:"not null"`
}

type Table0 struct {
	gorm.Model
	EncuestadorID uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID        uint   `json:"zona_id" gorm:"not null"`
	Sexo          string `json:"sexo" gorm:"not null"`
	Edad          int    `json:"edad" gorm:"not null"`
	Parentesco    string `json:"parentesco" gorm:"not null"`
	Escolaridad   string `json:"escolaridad" gorm:"not null"`
	Trabaja       string `json:"trabaja" gorm:"not null"`
	HorasPorDia   int    `json:"horas_dia" gorm:"not null"`
	DiasPorSemana int    `json:"dias_semana" gorm:"not null"`
	Ocupacion     string `json:"ocupacion" gorm:"not null"`
	Fumador       string `json:"fumador" gorm:"not null"`
	TiempoCiudad  int    `json:"tiempo_ciudad" gorm:"not null"`
}

type Table1 struct {
	gorm.Model
	EncuestadorID       uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID              uint   `json:"zona_id" gorm:"not null"`
	TieneElectricidad   string `json:"tiene_electricidad" gorm:"not null"`
	FuentesAgua         string `json:"fuentes_agua" gorm:"not null"`
	TieneGas            string `json:"tiene_gas" gorm:"not null"`
	EliminacionDesechos string `json:"eliminacion_desechos" gorm:"not null"`
	MaterialTecho       string `json:"material_techo" gorm:"not null"`
	MaterialPiso        string `json:"material_piso" gorm:"not null"`
	MaterialParedes     string `json:"material_paredes" gorm:"not null"`
	NumCuartos          int    `json:"num_cuartos" gorm:"not null"`
	OloresDesagradables string `json:"olores_desagradables" gorm:"not null"`
}

type Table2 struct {
	gorm.Model
	EncuestadorID    uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID           uint   `json:"zona_id" gorm:"not null"`
	LlaveInterior    string `json:"llave_interior" gorm:"not null"`
	Garrafon         string `json:"garrafon" gorm:"not null"`
	LlaveComunitaria string `json:"llave_comunitaria" gorm:"not null"`
}

type Table3 struct {
	gorm.Model
	EncuestadorID uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID        uint   `json:"zona_id" gorm:"not null"`
	Imss          string `json:"imss" gorm:"not null"`
	Issste        string `json:"issste" gorm:"not null"`
	Particular    string `json:"particular" gorm:"not null"`
}

type Table4 struct {
	gorm.Model
	EncuestadorID      uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID             uint   `json:"zona_id" gorm:"not null"`
	EnfermedadReferida string `json:"enfermedad_referida" gorm:"not null"`
	EnfermedadRecabada string `json:"enfermedad_recabada" gorm:"not null"`
}

type Table5 struct {
	gorm.Model
	EncuestadorID      uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID             uint   `json:"zona_id" gorm:"not null"`
	EnfermedadReferida string `json:"enfermedad_referida" gorm:"not null"`
	EnfermedadRecabada string `json:"enfermedad_recabada" gorm:"not null"`
}

type Table6 struct {
	gorm.Model
	EncuestadorID     uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID            uint   `json:"zona_id" gorm:"not null"`
	Sexo              string `json:"sexo" gorm:"not null"`
	EdadFallecimiento int    `json:"edad_fallecimiento" gorm:"not null"`
	AñoFallecimiento  int    `json:"año_fallecimiento" gorm:"not null"`
	CausaReferida     string `json:"causa_referida" gorm:"not null"`
	CausaRecabada     string `json:"causa_recabada" gorm:"not null"`
}

type Table7 struct {
	gorm.Model
	EncuestadorID uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID        uint   `json:"zona_id" gorm:"not null"`
	Sexo          string `json:"sexo" gorm:"not null"`
	EdadDetection int    `json:"edad_detection" gorm:"not null"`
	AñoDetection  int    `json:"año_detection" gorm:"not null"`
	TipoReferido  string `json:"tipo_referido" gorm:"not null"`
	TipoRecabado  string `json:"tipo_recabado" gorm:"not null"`
}

type Table8 struct {
	gorm.Model
	EncuestadorID        uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID               uint   `json:"zona_id" gorm:"not null"`
	AñoNacimientoPerdida int    `json:"año_nacimiento_perdida" gorm:"not null"`
	EnCurso              string `json:"en_curso" gorm:"not null"`
	TipoParto            string `json:"tipo_parto" gorm:"not null"`
	TuvoComplicaciones   string `json:"tuvo_complicaciones" gorm:"not null"`
	ComplicacionReferida string `json:"complicacion_referida" gorm:"not null"`
	ComplicacionRecabada string `json:"complicacion_recabada" gorm:"not null"`
}

type Table8_1 struct {
	gorm.Model
	EncuestadorID uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID        uint   `json:"zona_id" gorm:"not null"`
	AñoPerdida    int    `json:"año_perdida" gorm:"not null"`
	Trimestre     string `json:"trimestre" gorm:"not null"`
}

type Table9 struct {
	gorm.Model
	EncuestadorID        uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID               uint   `json:"zona_id" gorm:"not null"`
	BajoPeso             string `json:"bajo_peso" gorm:"not null"`
	Prematuro            string `json:"prematuro" gorm:"not null"`
	Malformacion         string `json:"malformacion" gorm:"not null"`
	MalformacionReferida string `json:"malformacion_referida" gorm:"not null"`
	MalformacionRecabada string `json:"malformacion_recabada" gorm:"not null"`
}

type Table9_1 struct {
	gorm.Model
	EncuestadorID    uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID           uint   `json:"zona_id" gorm:"not null"`
	Año              int    `json:"año" gorm:"not null"`
	ProblemaReferido string `json:"problema_referido" gorm:"not null"`
	ProblemaRecabado string `json:"problema_recabado" gorm:"not null"`
}

type Table10 struct {
	gorm.Model
	EncuestadorID        uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID               uint   `json:"zona_id" gorm:"not null"`
	DiscapacidadReferida string `json:"discapacidad_referida" gorm:"not null"`
	DiscapacidadRecabada string `json:"discapacidad_recabada" gorm:"not null"`
	TipoCondicion        string `json:"tipo_condicion" gorm:"not null"`
}

type Table10_1 struct {
	gorm.Model
	EncuestadorID    uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID           uint   `json:"zona_id" gorm:"not null"`
	SiDonde          string `json:"si_donde" gorm:"not null"`
	NoPorque         string `json:"no_porque" gorm:"not null"`
	TieneCertificado string `json:"tiene_certificado" gorm:"not null"`
}

type Table11 struct {
	gorm.Model
	EncuestadorID  uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID         uint   `json:"zona_id" gorm:"not null"`
	NombreFarmaco  string `json:"nombre_farmaco" gorm:"not null"`
	EsPrescrito    string `json:"es_prescrito" gorm:"not null"`
	MotivoReferido string `json:"motivo_referido" gorm:"not null"`
	MotivoRecabado string `json:"motivo_recabado" gorm:"not null"`
}

type Table12 struct {
	gorm.Model
	EncuestadorID uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID        uint   `json:"zona_id" gorm:"not null"`
	HayProblema   string `json:"hay_problema" gorm:"not null"`
	QueProblema   string `json:"que_problema" gorm:"not null"`
}

type Table13 struct {
	gorm.Model
	EncuestadorID    uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID           uint   `json:"zona_id" gorm:"not null"`
	HayContaminacion string `json:"hay_contaminacion" gorm:"not null"`
	CualEs           string `json:"cual_es" gorm:"not null"`
}

type Table14 struct {
	gorm.Model
	EncuestadorID uint   `json:"encuestador_id" gorm:"not null"`
	ZonaID        uint   `json:"zona_id" gorm:"not null"`
	Tiempo        string `json:"tiempo" gorm:"not null"`
}
