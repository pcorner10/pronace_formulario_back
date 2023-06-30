package apprepository

import "pronaces_back/internal/domain/entity"

type Form0Request struct {
	Integrante []entity.Table0 `json:"integrante"`
}

type Form0 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	Sexo             string `json:"sexo"`
	Edad             string `json:"edad"`
	Parentesco       string `json:"parentesco"`
	Escolaridad      string `json:"escolaridad"`
	Trabaja          string `json:"trabaja"`
	HorasPorDia      string `json:"horas_dia"`
	DiasPorSemana    string `json:"dias_semana"`
	Ocupacion        string `json:"ocupacion"`
	Fumador          string `json:"fumador"`
	TiempoCiudad     string `json:"tiempo_ciudad"`
}

type Form1 struct {
	EncuestadorEmail    string `json:"encuestador_email"`
	Zona                string `json:"zona"`
	TieneElectricidad   string `json:"tiene_electricidad"`
	FuentesAgua         string `json:"fuentes_agua"`
	TieneGas            string `json:"tiene_gas"`
	EliminacionDesechos string `json:"eliminacion_desechos"`
	MaterialTecho       string `json:"material_techo"`
	MaterialPiso        string `json:"material_piso"`
	MaterialParedes     string `json:"material_paredes"`
	NumCuartos          string `json:"num_cuartos"`
	OloresDesagradables string `json:"olores_desagradables"`
}

type Form2 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	LlaveInterior    string `json:"llave_interior"`
	Garrafon         string `json:"garrafon"`
	LlaveComunitaria string `json:"llave_comunitaria"`
}

type Form3 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	Imss             string `json:"imss"`
	Issste           string `json:"issste"`
	SeguroPopular    string `json:"seguro_popular"`
	Privado          string `json:"privado"`
	Ninguno          string `json:"ninguno"`
}

type Form4 struct {
	EncuestadorEmail   string `json:"encuestador_email"`
	Zona               string `json:"zona"`
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Form4Request struct {
	Integrante   []Form4 `json:"integrante"`
}

type Form5 struct {
	EncuestadorEmail   string `json:"encuestador_email"`
	Zona               string `json:"zona"`
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Form5Request struct {
	Integrante   []Form5 `json:"integrante"`
}

type Form6 struct {
	EncuestadorEmail  string `json:"encuestador_email"`
	Zona              string `json:"zona"`
	Sexo              string `json:"sexo"`
	EdadFallecimiento string `json:"edad_fallecimiento"`
	AñoFallecimiento  string `json:"año_fallecimiento"`
	CausaReferida     string `json:"causa_referida"`
	CausaRecabada     string `json:"causa_recabada"`
}

type Form6Request struct {
	Integrante   []Form6 `json:"integrante"`
}

type Form7 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	Sexo             string `json:"sexo"`
	EdadDetection    string `json:"edad_detection"`
	AñoDetection     string `json:"año_detection"`
	TipoReferido     string `json:"tipo_referido"`
	TipoRecabado     string `json:"tipo_recabado"`
}

type Form7Request struct {
	Integrante   []Form7 `json:"integrante"`
}

type Form8 struct {
	EncuestadorEmail     string `json:"encuestador_email"`
	Zona                 string `json:"zona"`
	AñoNacimientoPerdida string `json:"año_nacimiento_perdida"`
	EnCurso              string `json:"en_curso"`
	TipoParto            string `json:"tipo_parto"`
	TuvoComplicaciones   string `json:"tuvo_complicaciones"`
	ComplicacionReferida string `json:"complicacion_referida"`
	ComplicacionRecabada string `json:"complicacion_recabada"`
}

type Form8Request struct {
	Integrante   []Form8 `json:"integrante"`
}

type Form8_1 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	AñoPerdida       string `json:"año_perdida"`
	Trimestre        string `json:"trimestre"`
}

type Form8_1Request struct {
	Integrante   []Form8_1 `json:"integrante"`
}

type Form9 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
}

type Form9Request struct {
	Integrante   []Form9 `json:"integrante"`
}

type Form9_1 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	Año              string `json:"año"`
	ProblemaReferido string `json:"problema_referido"`
	ProblemaRecabado string `json:"problema_recabado"`
}

type Form9_1Request struct {
	Integrante   []Form9_1 `json:"integrante"`
}

type Form10 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	DiscapacidadReferida string `json:"discapacidad_referida"`
	DiscapacidadRecabada string `json:"discapacidad_recabada"`
	TipoCondicion        string `json:"tipo_condicion"`
}

type Form10Request struct {
	Integrante   []Form10 `json:"integrante"`
}

type Form10_1 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	SiDonde          string `json:"si_donde"`
	NoPorque         string `json:"no_porque"`
	TieneCertificado string `json:"tiene_certificado"`
}

type Form10_1Request struct {
	Integrante   []Form10_1 `json:"integrante"`
}

type Form11 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	NombreFarmaco    string `json:"nombre_farmaco"`
	EsPrescrito      string `json:"es_prescrito"`
	MotivoReferido   string `json:"motivo_referido"`
	MotivoRecabado   string `json:"motivo_recabado"`
}

type Form11Request struct {
	Integrante   []Form11 `json:"integrante"`
}

type Form12 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	HayProblema      string `json:"hay_problema"`
	QueProblema      string `json:"que_problema"`
}

type Form13 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	HayContaminacion string `json:"hay_contaminacion"`
	CualEs           string `json:"cual_es"`
}

type Form14 struct {
	EncuestadorEmail string `json:"encuestador_email"`
	Zona             string `json:"zona"`
	Tiempo           string `json:"tiempo"`
}

type Form14Request struct {
	Integrante   []Form14 `json:"integrante"`
}