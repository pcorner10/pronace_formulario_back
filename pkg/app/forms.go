package app

import (
	"fmt"
	"pronaces_back/pkg/domain"
	"strconv"
)

type surveyService struct {
	DB domain.FormDB
}

func NewSurveyService(db domain.FormDB) domain.FormService {
	return &surveyService{
		DB: db,
	}
}

func (u *surveyService) CreateForm0(req interface{}) error {

	form := []domain.Table0{}

	// Cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := req.(domain.Form0Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}
	encuestador, err = u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		edad, err := strconv.Atoi(integrante.Edad)
		if err != nil {
			return err
		}
		horasPorDia, err := strconv.Atoi(integrante.HorasPorDia)
		if err != nil {
			return err
		}
		diasPorSemana, err := strconv.Atoi(integrante.DiasPorSemana)
		if err != nil {
			return err
		}
		tiempoCiudad, err := strconv.Atoi(integrante.TiempoCiudad)
		if err != nil {
			return err
		}

		form = append(form, domain.Table0{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona.ID,
			Sexo:          integrante.Sexo,
			Edad:          edad,
			Parentesco:    integrante.Parentesco,
			Escolaridad:   integrante.Escolaridad,
			Trabaja:       integrante.Trabaja,
			HorasPorDia:   horasPorDia,
			DiasPorSemana: diasPorSemana,
			Ocupacion:     integrante.Ocupacion,
			Fumador:       integrante.Fumador,
			TiempoCiudad:  tiempoCiudad,
		})
	}

	err = u.DB.CreateForm0(form)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm123(forms interface{}) error {

	formReq, ok := forms.(domain.Form123Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	formReq.Form1.EncuestadorID = encuestador.ID
	formReq.Form1.ZonaID = zona.ID
	err = u.CreateForm1(formReq.Form1)
	if err != nil {
		return err
	}

	formReq.Form2.EncuestadorID = encuestador.ID
	formReq.Form2.ZonaID = zona.ID
	err = u.CreateForm2(formReq.Form2)
	if err != nil {
		return err
	}

	formReq.Form3.EncuestadorID = encuestador.ID
	formReq.Form3.ZonaID = zona.ID
	err = u.CreateForm3(formReq.Form3)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm1(form interface{}) error {

	res := domain.Table1{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form1)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	NumCuartos, err := strconv.Atoi(formReq.NumCuartos)
	if err != nil {
		return err
	}

	res.EncuestadorID = formReq.EncuestadorID
	res.ZonaID = formReq.ZonaID
	res.TieneElectricidad = formReq.TieneElectricidad
	res.FuentesAgua = formReq.FuentesAgua
	res.TieneGas = formReq.TieneGas
	res.EliminacionDesechos = formReq.EliminacionDesechos
	res.MaterialParedes = formReq.MaterialParedes
	res.MaterialPiso = formReq.MaterialPiso
	res.MaterialTecho = formReq.MaterialTecho
	res.NumCuartos = NumCuartos
	res.OloresDesagradables = formReq.OloresDesagradables

	err = u.DB.CreateForm1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm2(form interface{}) error {

	res := domain.Table2{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form2)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	res.EncuestadorID = formReq.EncuestadorID
	res.ZonaID = formReq.ZonaID
	res.LlaveInterior = formReq.LlaveInterior
	res.Garrafon = formReq.Garrafon
	res.LlaveComunitaria = formReq.LlaveComunitaria

	err := u.DB.CreateForm2(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm3(form interface{}) error {

	res := domain.Table3{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form3)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	res.EncuestadorID = formReq.EncuestadorID
	res.ZonaID = formReq.ZonaID
	res.Salud = formReq.Salud

	err := u.DB.CreateForm3(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm4(form interface{}) error {

	res := []domain.Table4{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form4Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {
		res = append(res, domain.Table4{
			EncuestadorID:      encuestador.ID,
			ZonaID:             zona.ID,
			Parentesco:         integrante.Parentesco,
			EnfermedadReferida: integrante.EnfermedadReferida,
			EnfermedadRecabada: integrante.EnfermedadRecabada,
		})
	}

	err = u.DB.CreateForm4(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm5(form interface{}) error {

	res := []domain.Table5{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form5Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {
		res = append(res, domain.Table5{
			EncuestadorID:      encuestador.ID,
			ZonaID:             zona.ID,
			EnfermedadReferida: integrante.EnfermedadReferida,
			EnfermedadRecabada: integrante.EnfermedadRecabada,
		})
	}

	err = u.DB.CreateForm5(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm6(form interface{}) error {

	res := []domain.Table6{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form6Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		EdadFallecimiento, err := strconv.Atoi(integrante.EdadFallecimiento)
		if err != nil {
			return err
		}

		AñoFallecimiento, err := strconv.Atoi(integrante.AñoFallecimiento)
		if err != nil {
			return err
		}

		res = append(res, domain.Table6{
			EncuestadorID:     encuestador.ID,
			ZonaID:            zona.ID,
			Sexo:              integrante.Sexo,
			AñoFallecimiento:  AñoFallecimiento,
			EdadFallecimiento: EdadFallecimiento,
			CausaReferida:     integrante.CausaReferida,
			CausaRecabada:     integrante.CausaRecabada,
		})
	}

	err = u.DB.CreateForm6(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm7(form interface{}) error {

	res := []domain.Table7{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form7Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Form7 {

		EdadDetection, err := strconv.Atoi(integrante.EdadDetection)
		if err != nil {

			return fmt.Errorf("error al convertir edad de deteccion: %v", err)
		}

		AñoDetection, err := strconv.Atoi(integrante.AñoDetection)
		if err != nil {
			return fmt.Errorf("error al convertir año de deteccion: %v", err)
		}

		res = append(res, domain.Table7{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona.ID,
			Sexo:          integrante.Sexo,
			EdadDetection: EdadDetection,
			AñoDetection:  AñoDetection,
			TipoReferido:  integrante.TipoReferido,
			TipoRecabado:  integrante.TipoRecabado,
		})

	}

	err = u.DB.CreateForm7(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm8(form interface{}) error {

	res := []domain.Table8{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form8Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Form8 {

		AñoNacimientoPerdida, err := strconv.Atoi(integrante.AñoNacimientoPerdida)
		if err != nil {
			return err
		}

		res = append(res, domain.Table8{
			EncuestadorID:        encuestador.ID,
			ZonaID:               zona.ID,
			AñoNacimientoPerdida: AñoNacimientoPerdida,
			EnCurso:              integrante.EnCurso,
			TipoParto:            integrante.TipoParto,
			TuvoComplicaciones:   integrante.TuvoComplicaciones,
			ComplicacionReferida: integrante.ComplicacionReferida,
			ComplicacionRecabada: integrante.ComplicacionRecabada,
		})

	}

	err = u.DB.CreateForm8(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm8_1(form interface{}) error {

	res := []domain.Table8_1{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form8_1Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Form8_1 {

		AñoPerdida, err := strconv.Atoi(integrante.AñoPerdida)
		if err != nil {
			return err
		}

		res = append(res, domain.Table8_1{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona.ID,
			AñoPerdida:    AñoPerdida,
			Trimestre:     integrante.Trimestre,
		})
	}

	err = u.DB.CreateForm8_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm9(form interface{}) error {

	res := []domain.Table9{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form9Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		res = append(res, domain.Table9{
			EncuestadorID:        encuestador.ID,
			ZonaID:               zona.ID,
			BajoPeso:             integrante.BajoPeso,
			Prematuro:            integrante.Prematuro,
			Malformacion:         integrante.Malformacion,
			MalformacionReferida: integrante.MalformacionReferida,
			MalformacionRecabada: integrante.MalformacionRecabada,
		})
	}

	err = u.DB.CreateForm9(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm9_1(form interface{}) error {

	res := []domain.Table9_1{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form9_1Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		año, err := strconv.Atoi(integrante.Año)
		if err != nil {
			return err
		}

		res = append(res, domain.Table9_1{
			EncuestadorID:    encuestador.ID,
			ZonaID:           zona.ID,
			Año:              año,
			ProblemaReferido: integrante.ProblemaReferido,
			ProblemaRecabado: integrante.ProblemaRecabado,
		})
	}

	err = u.DB.CreateForm9_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm10(form interface{}) error {

	res := []domain.Table10{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form10Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		res = append(res, domain.Table10{
			EncuestadorID:        encuestador.ID,
			ZonaID:               zona.ID,
			DiscapacidadReferida: integrante.DiscapacidadReferida,
			DiscapacidadRecabada: integrante.DiscapacidadRecabada,
			TipoCondicion:        integrante.TipoCondicion,
		})
	}

	err = u.DB.CreateForm10(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm10_1(form interface{}) error {

	res := []domain.Table10_1{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form10_1Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		res = append(res, domain.Table10_1{
			EncuestadorID:    encuestador.ID,
			ZonaID:           zona.ID,
			SiDonde:          integrante.SiDonde,
			NoPorque:         integrante.NoPorque,
			TieneCertificado: integrante.TieneCertificado,
		})
	}

	err = u.DB.CreateForm10_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm11(form interface{}) error {

	res := []domain.Table11{}
	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form11Request)

	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		res = append(res, domain.Table11{
			EncuestadorID:  encuestador.ID,
			ZonaID:         zona.ID,
			NombreFarmaco:  integrante.NombreFarmaco,
			EsPrescrito:    integrante.EsPrescrito,
			MotivoReferido: integrante.MotivoReferido,
			MotivoRecabado: integrante.MotivoRecabado,
		})
	}

	err = u.DB.CreateForm11(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm1213(forms interface{}) error {

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := forms.(domain.Form1213Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	formReq.Form12.EncuestadorID = encuestador.ID
	formReq.Form12.ZonaID = zona.ID
	err = u.CreateForm12(formReq.Form12)
	if err != nil {
		return err
	}

	formReq.Form13.EncuestadorID = encuestador.ID
	formReq.Form13.ZonaID = zona.ID
	err = u.CreateForm13(formReq.Form13)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm12(form interface{}) error {

	res := domain.Table12{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form12)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	res.EncuestadorID = formReq.EncuestadorID
	res.ZonaID = formReq.ZonaID
	res.HayProblema = formReq.HayProblema
	res.QueProblema = formReq.QueProblema

	err := u.DB.CreateForm12(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm13(form interface{}) error {

	res := domain.Table13{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form13)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	res.EncuestadorID = formReq.EncuestadorID
	res.ZonaID = formReq.ZonaID
	res.HayContaminacion = formReq.HayContaminacion
	res.CualEs = formReq.CualEs

	err := u.DB.CreateForm13(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *surveyService) CreateForm14(form interface{}) error {

	res := []domain.Table14{}

	// cast the req interface to the type that has the EncuestadorEmail field
	formReq, ok := form.(domain.Form14Request)
	if !ok {
		return fmt.Errorf("invalid request type")
	}

	encuestador, err := u.DB.GetUserByEmail(formReq.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DB.FirstOrCreateZona(formReq.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range formReq.Integrante {

		res = append(res, domain.Table14{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona.ID,
			Tiempo:        integrante.Tiempo,
		})
	}

	err = u.DB.CreateForm14(res)
	if err != nil {
		return err
	}

	return nil
}
