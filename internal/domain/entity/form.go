package entity

import (
	"gorm.io/gorm"
)

type Zona struct {
	gorm.Model
	Municipio string `json:"municipio"`
	Localidad string `json:"localidad"`
	Colonia   string `json:"colonia"`
	Manzana   string `json:"manzana"`
	Lote      string `json:"lote"`
}

type Table0 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID           uint   `json:"zona_id"`
	Sexo             string `json:"sexo"`
	Edad             int    `json:"edad"`
	Parentesco       string `json:"parentesco"`
	Escolaridad      string `json:"escolaridad"`
	Trabaja          string `json:"trabaja"`
	HorasPorDia      int    `json:"horas_dia"`
	DiasPorSemana    int    `json:"dias_semana"`
	Ocupacion        string `json:"ocupacion"`
	Fumador          string `json:"fumador"`
	TiempoCiudad     int    `json:"tiempo_ciudad"`
}

type Table1 struct {
	gorm.Model
	EncuestadorEmail    string `json:"encuestador_email"`
	ZonaID              uint   `json:"zona_id"`
	TieneElectricidad   string `json:"tiene_electricidad"`
	FuentesAgua         string `json:"fuentes_agua"`
	TieneGas            string `json:"tiene_gas"`
	EliminacionDesechos string `json:"eliminacion_desechos"`
	MaterialTecho       string `json:"material_techo"`
	MaterialPiso        string `json:"material_piso"`
	MaterialParedes     string `json:"material_paredes"`
	NumCuartos          int    `json:"num_cuartos"`
	OloresDesagradables string `json:"olores_desagradables"`
}

type Table2 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID           uint   `json:"zona_id"`
	LlaveInterior    string `json:"llave_interior"`
	Garrafon         string `json:"garrafon"`
	LlaveComunitaria string `json:"llave_comunitaria"`
}

type Table3 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID        uint   `json:"zona_id"`
	Imss          string `json:"imss"`
	Issste        string `json:"issste"`
	Particular    string `json:"particular"`
}

type Table4 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID             uint   `json:"zona_id"`
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Table5 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID             uint   `json:"zona_id"`
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Table6 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID            uint   `json:"zona_id"`
	Sexo              string `json:"sexo"`
	EdadFallecimiento int    `json:"edad_fallecimiento"`
	AñoFallecimiento  int    `json:"año_fallecimiento"`
	CausaReferida     string `json:"causa_referida"`
	CausaRecabada     string `json:"causa_recabada"`
}

type Table7 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID        uint   `json:"zona_id"`
	Sexo          string `json:"sexo"`
	EdadDetection int    `json:"edad_detection"`
	AñoDetection  int    `json:"año_detection"`
	TipoReferido  string `json:"tipo_referido"`
	TipoRecabado  string `json:"tipo_recabado"`
}

type Table8 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID               uint   `json:"zona_id"`
	AñoNacimientoPerdida int    `json:"año_nacimiento_perdida"`
	EnCurso              string `json:"en_curso"`
	TipoParto            string `json:"tipo_parto"`
	TuvoComplicaciones   string `json:"tuvo_complicaciones"`
	ComplicacionReferida string `json:"complicacion_referida"`
	ComplicacionRecabada string `json:"complicacion_recabada"`
}

type Table8_1 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID        uint   `json:"zona_id"`
	AñoPerdida    int    `json:"año_perdida"`
	Trimestre     string `json:"trimestre"`
}

type Table9 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID               uint   `json:"zona_id"`
	BajoPeso             string `json:"bajo_peso"`
	Prematuro            string `json:"prematuro"`
	Malformacion         string `json:"malformacion"`
	MalformacionReferida string `json:"malformacion_referida"`
	MalformacionRecabada string `json:"malformacion_recabada"`
}

type Table9_1 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID           uint   `json:"zona_id"`
	Año              int    `json:"año"`
	ProblemaReferido string `json:"problema_referido"`
	ProblemaRecabado string `json:"problema_recabado"`
}

type Table10 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID               uint   `json:"zona_id"`
	DiscapacidadReferida string `json:"discapacidad_referida"`
	DiscapacidadRecabada string `json:"discapacidad_recabada"`
	TipoCondicion        string `json:"tipo_condicion"`
}

type Table10_1 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID           uint   `json:"zona_id"`
	SiDonde          string `json:"si_donde"`
	NoPorque         string `json:"no_porque"`
	TieneCertificado string `json:"tiene_certificado"`
}

type Table11 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID         uint   `json:"zona_id"`
	NombreFarmaco  string `json:"nombre_farmaco"`
	EsPrescrito    string `json:"es_prescrito"`
	MotivoReferido string `json:"motivo_referido"`
	MotivoRecabado string `json:"motivo_recabado"`
}

type Table12 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID        uint   `json:"zona_id"`
	HayProblema   string `json:"hay_problema"`
	QueProblema   string `json:"que_problema"`
}

type Table13 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID           uint   `json:"zona_id"`
	HayContaminacion string `json:"hay_contaminacion"`
	CualEs           string `json:"cual_es"`
}

type Table14 struct {
	gorm.Model
	EncuestadorEmail string `json:"encuestador_email"`
	ZonaID        uint   `json:"zona_id"`
	Tiempo        string `json:"tiempo"`
}
