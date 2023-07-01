package domain

import (
	"time"
)

type FormDB interface {
	CreateForm0(table0 []Table0) error
	CreateForm1(table1 Table1) error
	CreateForm2(table2 Table2) error
	CreateForm3(table3 Table3) error
	CreateForm4(table4 []Table4) error
	CreateForm5(table5 []Table5) error
	CreateForm6(table6 []Table6) error
	CreateForm7(table7 []Table7) error
	CreateForm8(table8 []Table8) error
	CreateForm8_1(table8_1 []Table8_1) error
	CreateForm9(table9 []Table9) error
	CreateForm9_1(table9 []Table9_1) error
	CreateForm10(table10 []Table10) error
	CreateForm10_1(table10 []Table10_1) error
	CreateForm11(table11 []Table11) error
	CreateForm12(table12 Table12) error
	CreateForm13(table13 Table13) error
	CreateForm14(table14 []Table14) error
}

type FormService interface {
	CreateForm0(table0 []Table0) error
	CreateForm1(table1 Table1) error
	CreateForm2(table2 Table2) error
	CreateForm3(table3 Table3) error
	CreateForm4(table4 []Table4) error
	CreateForm5(table5 []Table5) error
	CreateForm6(table6 []Table6) error
	CreateForm7(table7 []Table7) error
	CreateForm8(table8 []Table8) error
	CreateForm8_1(table8_1 []Table8_1) error
	CreateForm9(table9 []Table9) error
	CreateForm9_1(table9 []Table9_1) error
	CreateForm10(table10 []Table10) error
	CreateForm10_1(table10 []Table10_1) error
	CreateForm11(table11 []Table11) error
	CreateForm12(table12 Table12) error
	CreateForm13(table13 Table13) error
	CreateForm14(table14 []Table14) error
}

type Zona struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Municipio string    `json:"municipio"`
	Localidad string    `json:"localidad"`
	Colonia   string    `json:"colonia"`
	Manzana   string    `json:"manzana"`
	Lote      string    `json:"lote"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Table0 struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID uint      `json:"encuestador_id"`
	Encuestador   User      `gorm:"foreignKey:EncuestadorID"`
	ZonaID        uint      `json:"zona_id"`
	Zona          Zona      `gorm:"foreignKey:ZonaID"`
	Sexo          string    `json:"sexo"`
	Edad          int       `json:"edad"`
	Parentesco    string    `json:"parentesco"`
	Escolaridad   string    `json:"escolaridad"`
	Trabaja       string    `json:"trabaja"`
	HorasPorDia   int       `json:"horas_dia"`
	DiasPorSemana int       `json:"dias_semana"`
	Ocupacion     string    `json:"ocupacion"`
	Fumador       string    `json:"fumador"`
	TiempoCiudad  int       `json:"tiempo_ciudad"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Table1 struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID       uint      `json:"encuestador_id"`
	ZonaID              uint      `json:"zona_id"`
	TieneElectricidad   string    `json:"tiene_electricidad"`
	FuentesAgua         string    `json:"fuentes_agua"`
	TieneGas            string    `json:"tiene_gas"`
	EliminacionDesechos string    `json:"eliminacion_desechos"`
	MaterialTecho       string    `json:"material_techo"`
	MaterialPiso        string    `json:"material_piso"`
	MaterialParedes     string    `json:"material_paredes"`
	NumCuartos          int       `json:"num_cuartos"`
	OloresDesagradables string    `json:"olores_desagradables"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type Table2 struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID    uint      `json:"encuestador_id"`
	ZonaID           uint      `json:"zona_id"`
	LlaveInterior    string    `json:"llave_interior"`
	Garrafon         string    `json:"garrafon"`
	LlaveComunitaria string    `json:"llave_comunitaria"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Table3 struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID uint      `json:"encuestador_id"`
	ZonaID        uint      `json:"zona_id"`
	Imss          string    `json:"imss"`
	Issste        string    `json:"issste"`
	SeguroPopular string    `json:"seguro_popular"`
	Privado       string    `json:"privado"`
	Ninguno       string    `json:"ninguno"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Table4 struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID      uint      `json:"encuestador_id"`
	ZonaID             uint      `json:"zona_id"`
	EnfermedadReferida string    `json:"enfermedad_referida"`
	EnfermedadRecabada string    `json:"enfermedad_recabada"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Table5 struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID      uint      `json:"encuestador_id"`
	ZonaID             uint      `json:"zona_id"`
	EnfermedadReferida string    `json:"enfermedad_referida"`
	EnfermedadRecabada string    `json:"enfermedad_recabada"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Table6 struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID     uint      `json:"encuestador_id"`
	ZonaID            uint      `json:"zona_id"`
	Sexo              string    `json:"sexo"`
	EdadFallecimiento int       `json:"edad_fallecimiento"`
	AñoFallecimiento  int       `json:"año_fallecimiento"`
	CausaReferida     string    `json:"causa_referida"`
	CausaRecabada     string    `json:"causa_recabada"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Table7 struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID uint      `json:"encuestador_id"`
	ZonaID        uint      `json:"zona_id"`
	Sexo          string    `json:"sexo"`
	EdadDetection int       `json:"edad_detection"`
	AñoDetection  int       `json:"año_detection"`
	TipoReferido  string    `json:"tipo_referido"`
	TipoRecabado  string    `json:"tipo_recabado"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Table8 struct {
	ID                   uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID        uint      `json:"encuestador_id"`
	ZonaID               uint      `json:"zona_id"`
	AñoNacimientoPerdida int       `json:"año_nacimiento_perdida"`
	EnCurso              string    `json:"en_curso"`
	TipoParto            string    `json:"tipo_parto"`
	TuvoComplicaciones   string    `json:"tuvo_complicaciones"`
	ComplicacionReferida string    `json:"complicacion_referida"`
	ComplicacionRecabada string    `json:"complicacion_recabada"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type Table8_1 struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID uint      `json:"encuestador_id"`
	ZonaID        uint      `json:"zona_id"`
	AñoPerdida    int       `json:"año_perdida"`
	Trimestre     string    `json:"trimestre"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Table9 struct {
	ID                   uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID        uint      `json:"encuestador_id"`
	ZonaID               uint      `json:"zona_id"`
	BajoPeso             string    `json:"bajo_peso"`
	Prematuro            string    `json:"prematuro"`
	Malformacion         string    `json:"malformacion"`
	MalformacionReferida string    `json:"malformacion_referida"`
	MalformacionRecabada string    `json:"malformacion_recabada"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type Table9_1 struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID    uint      `json:"encuestador_id"`
	ZonaID           uint      `json:"zona_id"`
	Año              int       `json:"año"`
	ProblemaReferido string    `json:"problema_referido"`
	ProblemaRecabado string    `json:"problema_recabado"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Table10 struct {
	ID                   uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID        uint      `json:"encuestador_id"`
	ZonaID               uint      `json:"zona_id"`
	DiscapacidadReferida string    `json:"discapacidad_referida"`
	DiscapacidadRecabada string    `json:"discapacidad_recabada"`
	TipoCondicion        string    `json:"tipo_condicion"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type Table10_1 struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID    uint      `json:"encuestador_id"`
	ZonaID           uint      `json:"zona_id"`
	SiDonde          string    `json:"si_donde"`
	NoPorque         string    `json:"no_porque"`
	TieneCertificado string    `json:"tiene_certificado"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Table11 struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID  uint      `json:"encuestador_id"`
	ZonaID         uint      `json:"zona_id"`
	NombreFarmaco  string    `json:"nombre_farmaco"`
	EsPrescrito    string    `json:"es_prescrito"`
	MotivoReferido string    `json:"motivo_referido"`
	MotivoRecabado string    `json:"motivo_recabado"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Table12 struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID uint      `json:"encuestador_id"`
	ZonaID        uint      `json:"zona_id"`
	HayProblema   string    `json:"hay_problema"`
	QueProblema   string    `json:"que_problema"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Table13 struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID    uint      `json:"encuestador_id"`
	ZonaID           uint      `json:"zona_id"`
	HayContaminacion string    `json:"hay_contaminacion"`
	CualEs           string    `json:"cual_es"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Table14 struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	EncuestadorID uint      `json:"encuestador_id"`
	ZonaID        uint      `json:"zona_id"`
	Tiempo        string    `json:"tiempo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
