package domain

type FormService interface {
	CreateForm0(req Form0Request) error
	CreateForm123(req Form123Request) error
	CreateForm1(req Form1, encuestador User, zona Zona) error
	CreateForm2(req Form2, encuestador User, zona Zona) error
	CreateForm3(req Form3, encuestador User, zona Zona) error
	CreateForm4(req Form4Request) error
	CreateForm5(req Form5Request) error
	CreateForm6(req Form6Request) error
	CreateForm7(req Form7Request) error
	CreateForm8(req Form8Request) error
	CreateForm8_1(req Form8_1Request) error
	CreateForm9(req Form9Request) error
	CreateForm9_1(req Form9_1Request) error
	CreateForm10(req Form10Request) error
	CreateForm10_1(req Form10_1Request) error
	CreateForm11(req Form11Request) error
	CreateForm1213(forms Form1213Request) error
	CreateForm12(req Form12, encuestador User, zona Zona) error
	CreateForm13(req Form13, encuestador User, zona Zona) error
	CreateForm14(req Form14Request) error
}
type Form0Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Integrante       []Form0 `json:"integrante"`
}

type Form0 struct {
	Sexo          string `json:"sexo"`
	Edad          string `json:"edad"`
	Parentesco    string `json:"parentesco"`
	Escolaridad   string `json:"escolaridad"`
	Trabaja       string `json:"trabaja"`
	HorasPorDia   string `json:"horas_dia"`
	DiasPorSemana string `json:"dias_semana"`
	Ocupacion     string `json:"ocupacion"`
	Fumador       string `json:"fumador"`
	TiempoCiudad  string `json:"tiempo_ciudad"`
}

type Form1 struct {
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
	LlaveInterior    string `json:"llave_interior"`
	Garrafon         string `json:"garrafon"`
	LlaveComunitaria string `json:"llave_comunitaria"`
}

type Form3 struct {
	Imss          string `json:"imss"`
	Issste        string `json:"issste"`
	SeguroPopular string `json:"seguro_popular"`
	Privado       string `json:"privado"`
	Ninguno       string `json:"ninguno"`
}

type Form123Request struct {
	EncuestadorEmail string `json:"email"`
	Zona             Zona   `json:"zona"`
	Form1            Form1  `json:"form1"`
	Form2            Form2  `json:"form2"`
	Form3            Form3  `json:"form3"`
}

type Form4 struct {
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Form4Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Integrante       []Form4 `json:"integrante"`
}

type Form5 struct {
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Form5Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Integrante       []Form5 `json:"integrante"`
}

type Form6 struct {
	Sexo              string `json:"sexo"`
	EdadFallecimiento string `json:"edad_fallecimiento"`
	AñoFallecimiento  string `json:"año_fallecimiento"`
	CausaReferida     string `json:"causa_referida"`
	CausaRecabada     string `json:"causa_recabada"`
}

type Form6Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Integrante       []Form6 `json:"integrante"`
}

type Form7 struct {
	Sexo          string `json:"sexo"`
	EdadDetection string `json:"edad_deteccion"`
	AñoDetection  string `json:"año_deteccion"`
	TipoReferido  string `json:"tipo_referida"`
	TipoRecabado  string `json:"tipo_recabada"`
}

type Form7Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Form7            []Form7 `json:"form7"`
}

type Form8 struct {
	AñoNacimientoPerdida string `json:"año_nacimiento_perdida"`
	EnCurso              string `json:"en_curso"`
	TipoParto            string `json:"tipo_parto"`
	TuvoComplicaciones   string `json:"tuvo_complicaciones"`
	ComplicacionReferida string `json:"complicacion_referida"`
	ComplicacionRecabada string `json:"complicacion_recabada"`
}

type Form8Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Form8            []Form8 `json:"form8"`
}

type Form8_1 struct {
	AñoPerdida string `json:"año_perdida"`
	Trimestre  string `json:"trimestre"`
}

type Form8_1Request struct {
	EncuestadorEmail string    `json:"email"`
	Zona             Zona      `json:"zona"`
	Form8_1          []Form8_1 `json:"form8_1"`
}

type Form9 struct {
	BajoPeso             string `json:"bajo_peso"`
	Prematuro            string `json:"prematuro"`
	Malformacion         string `json:"malformacion"`
	MalformacionReferida string `json:"malformacion_referida"`
	MalformacionRecabada string `json:"malformacion_recabada"`
}

type Form9Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Integrante       []Form9 `json:"integrante"`
}

type Form9_1 struct {
	Año              string `json:"año"`
	ProblemaReferido string `json:"problema_referido"`
	ProblemaRecabado string `json:"problema_recabado"`
}

type Form9_1Request struct {
	EncuestadorEmail string    `json:"email"`
	Zona             Zona      `json:"zona"`
	Integrante       []Form9_1 `json:"integrante"`
}

type Form10 struct {
	DiscapacidadReferida string `json:"discapacidad_referida"`
	DiscapacidadRecabada string `json:"discapacidad_recabada"`
	TipoCondicion        string `json:"tipo_condicion"`
}

type Form10Request struct {
	EncuestadorEmail string   `json:"email"`
	Zona             Zona     `json:"zona"`
	Integrante       []Form10 `json:"integrante"`
}

type Form10_1 struct {
	SiDonde          string `json:"si_donde"`
	NoPorque         string `json:"no_porque"`
	TieneCertificado string `json:"tiene_certificado"`
}

type Form10_1Request struct {
	EncuestadorEmail string     `json:"email"`
	Zona             Zona       `json:"zona"`
	Integrante       []Form10_1 `json:"integrante"`
}

type Form11 struct {
	NombreFarmaco  string `json:"nombre_farmaco"`
	EsPrescrito    string `json:"es_prescrito"`
	MotivoReferido string `json:"motivo_referido"`
	MotivoRecabado string `json:"motivo_recabado"`
}

type Form11Request struct {
	EncuestadorEmail string   `json:"email"`
	Zona             Zona     `json:"zona"`
	Integrante       []Form11 `json:"integrante"`
}

type Form12 struct {
	HayProblema string `json:"hay_problema"`
	QueProblema string `json:"que_problema"`
}

type Form13 struct {
	HayContaminacion string `json:"hay_contaminacion"`
	CualEs           string `json:"cual_es"`
}

type Form1213Request struct {
	EncuestadorEmail string `json:"email"`
	Zona             Zona   `json:"zona"`
	Form12           Form12 `json:"form12"`
	Form13           Form13 `json:"form13"`
}

type Form14 struct {
	Tiempo string `json:"tiempo"`
}

type Form14Request struct {
	EncuestadorEmail string   `json:"email"`
	Zona             Zona     `json:"zona"`
	Integrante       []Form14 `json:"integrante"`
}
