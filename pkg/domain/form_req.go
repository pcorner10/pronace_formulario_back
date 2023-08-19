package domain

type FormService interface {
	LoginUser(email, password string) (*ReponseLogin, error)
	RegisterUser(user *User) error

	CreateUser(user User) error
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user User) error
	DeleteUser(userID uint64) error

	FirstOrCreateZona(zona Zona) (*Zona, error)
	CreateZona(zona Zona) (*Zona, error)
	GetZonaByID(zonaID uint) (*Zona, error)
	GetZonaID(zona Zona) (uint, error)

	CreateForm0(table0 interface{}) error
	CreateForm123(forms interface{}) error
	CreateForm1(table1 interface{}) error
	CreateForm2(table2 interface{}) error
	CreateForm3(table3 interface{}) error
	CreateForm4(table4 interface{}) error
	CreateForm5(table5 interface{}) error
	CreateForm6(table6 interface{}) error
	CreateForm7(table7 interface{}) error
	CreateForm8(table8 interface{}) error
	CreateForm8_1(table8_1 interface{}) error
	CreateForm9(table9 interface{}) error
	CreateForm9_1(table9 interface{}) error
	CreateForm10(table10 interface{}) error
	CreateForm10_1(table10 interface{}) error
	CreateForm11(table11 interface{}) error
	CreateForm1213(forms interface{}) error
	CreateForm12(table12 interface{}) error
	CreateForm13(table13 interface{}) error
	CreateForm14(table14 interface{}) error
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
	EncuestadorID       uint   `json:"encuestador_id"`
	ZonaID              uint   `json:"zona_id"`
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
	EncuestadorID    uint   `json:"encuestador_id"`
	ZonaID           uint   `json:"zona_id"`
	LlaveInterior    string `json:"llave_interior"`
	Garrafon         string `json:"garrafon"`
	LlaveComunitaria string `json:"llave_comunitaria"`
}

type Form3 struct {
	EncuestadorID uint     `json:"encuestador_id"`
	ZonaID        uint     `json:"zona_id"`
	Salud         []string `json:"salud"`
}

type Form123Request struct {
	EncuestadorEmail string `json:"email"`
	Zona             Zona   `json:"zona"`
	Form1            Form1  `json:"form1"`
	Form2            Form2  `json:"form2"`
	Form3            Form3  `json:"form3"`
}

type Form4 struct {
	Parentesco         string `json:"parentesco"`
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Form4Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Integrante       []Form4 `json:"integrante"`
}

type Form5 struct {
	Parentesco         string `json:"parentesco"`
	EnfermedadReferida string `json:"enfermedad_referida"`
	EnfermedadRecabada string `json:"enfermedad_recabada"`
}

type Form5Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Integrante       []Form5 `json:"integrante"`
}

type Form6 struct {
	Parentesco        string `json:"parentesco"`
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
	Parentesco    string `json:"parentesco"`
	EdadDetection string `json:"edad_deteccion"`
	AñoDetection  string `json:"año_deteccion"`
	TipoReferido  string `json:"tipo_referida"`
	TipoRecabado  string `json:"tipo_recabada"`
}

type Form7Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Form7            []Form7 `json:"integrante"`
}

type Form8 struct {
	AñoNacimientoPerdida string `json:"año_nacimiento_perdida"`
	EnCurso              string `json:"en_curso"`
	TipoParto            string `json:"tipo_parto"`
	Parentesco           string `json:"parentesco"`
	TuvoComplicaciones   string `json:"tuvo_complicacion"`
	ComplicacionReferida string `json:"complicacion_referida"`
	ComplicacionRecabada string `json:"complicacion_recabada"`
}

type Form8Request struct {
	EncuestadorEmail string  `json:"email"`
	Zona             Zona    `json:"zona"`
	Form8            []Form8 `json:"integrante"`
}

type Form8_1 struct {
	Parentesco string `json:"parentesco"`
	AñoPerdida string `json:"año_perdida"`
	Trimestre  string `json:"trimestre"`
}

type Form8_1Request struct {
	EncuestadorEmail string    `json:"email"`
	Zona             Zona      `json:"zona"`
	Form8_1          []Form8_1 `json:"integrante"`
}

type Form9 struct {
	Parentesco           string `json:"parentesco"`
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
	Parentesco       string `json:"parentesco"`
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
	Parentesco           string `json:"parentesco"`
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
	Parentesco       string `json:"parentesco"`
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
	Parentesco     string `json:"parentesco"`
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
	EncuestadorID uint   `json:"encuestador_id"`
	ZonaID        uint   `json:"zona_id"`
	HayProblema   string `json:"hay_problema"`
	QueProblema   string `json:"que_problema"`
}

type Form13 struct {
	EncuestadorID    uint   `json:"encuestador_id"`
	ZonaID           uint   `json:"zona_id"`
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
	Parentesco string `json:"parentesco"`
	Tiempo     string `json:"tiempo"`
}

type Form14Request struct {
	EncuestadorEmail string   `json:"email"`
	Zona             Zona     `json:"zona"`
	Integrante       []Form14 `json:"integrante"`
}
