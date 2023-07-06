package app

import (
	"fmt"
	"pronaces_back/pkg/domain"
	"strconv"
)

type formService struct {
	DBForm domain.FormDB
	DBUser domain.UserDB
	DBZona domain.ZonaDB
}

func NewFormService(dbForm domain.FormDB, dbUser domain.UserDB, dbZona domain.ZonaDB) domain.FormService {
	return &formService{
		DBForm: dbForm,
		DBUser: dbUser,
		DBZona: dbZona,
	}
}

func (u *formService) CreateForm0(req domain.Form0Request) error {

	form := []domain.Table0{}

	encuestador, err := u.DBUser.GetByEmail(req.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(req.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range req.Integrante {

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

	err = u.DBForm.CreateForm0(form)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm123(forms domain.Form123Request) error {
	encuestador, err := u.DBUser.GetByEmail(forms.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(forms.Zona)
	if err != nil {
		return err
	}

	err = u.CreateForm1(forms.Form1, *encuestador, *zona)
	if err != nil {
		return err
	}

	err = u.CreateForm2(forms.Form2, *encuestador, *zona)
	if err != nil {
		return err
	}

	err = u.CreateForm3(forms.Form3, *encuestador, *zona)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm1(form domain.Form1, encuestador domain.User, zona domain.Zona) error {

	res := domain.Table1{}
	fmt.Println(form)
	NumCuartos, err := strconv.Atoi(form.NumCuartos)
	if err != nil {
		return err
	}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona.ID
	res.TieneElectricidad = form.TieneElectricidad
	res.FuentesAgua = form.FuentesAgua
	res.TieneGas = form.TieneGas
	res.EliminacionDesechos = form.EliminacionDesechos
	res.MaterialParedes = form.MaterialParedes
	res.MaterialPiso = form.MaterialPiso
	res.MaterialTecho = form.MaterialTecho
	res.NumCuartos = NumCuartos
	res.OloresDesagradables = form.OloresDesagradables

	err = u.DBForm.CreateForm1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm2(form domain.Form2, encuestador domain.User, zona domain.Zona) error {

	res := domain.Table2{}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona.ID
	res.LlaveInterior = form.LlaveInterior
	res.Garrafon = form.Garrafon
	res.LlaveComunitaria = form.LlaveComunitaria

	err := u.DBForm.CreateForm2(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm3(form domain.Form3, encuestador domain.User, zona domain.Zona) error {

	res := domain.Table3{}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona.ID
	res.Imss = form.Imss
	res.Issste = form.Issste
	res.SeguroPopular = form.SeguroPopular
	res.Privado = form.Privado
	res.Ninguno = form.Ninguno

	err := u.DBForm.CreateForm3(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm4(form domain.Form4Request) error {

	res := []domain.Table4{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {
		res = append(res, domain.Table4{
			EncuestadorID:      encuestador.ID,
			ZonaID:             zona.ID,
			EnfermedadReferida: integrante.EnfermedadReferida,
			EnfermedadRecabada: integrante.EnfermedadRecabada,
		})
	}

	err = u.DBForm.CreateForm4(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm5(form domain.Form5Request) error {

	res := []domain.Table5{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {
		res = append(res, domain.Table5{
			EncuestadorID:      encuestador.ID,
			ZonaID:             zona.ID,
			EnfermedadReferida: integrante.EnfermedadReferida,
			EnfermedadRecabada: integrante.EnfermedadRecabada,
		})
	}

	err = u.DBForm.CreateForm5(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm6(form domain.Form6Request) error {

	res := []domain.Table6{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

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

	err = u.DBForm.CreateForm6(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm7(form domain.Form7Request) error {

	res := []domain.Table7{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Form7 {

		EdadDetection, err := strconv.Atoi(integrante.EdadDetection)
		if err != nil {
			fmt.Println(integrante.EdadDetection)
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

	err = u.DBForm.CreateForm7(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm8(form domain.Form8Request) error {

	res := []domain.Table8{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Form8 {

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

	err = u.DBForm.CreateForm8(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm8_1(form domain.Form8_1Request) error {

	res := []domain.Table8_1{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Form8_1 {

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

	err = u.DBForm.CreateForm8_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm9(form domain.Form9Request) error {

	res := []domain.Table9{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {
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

	err = u.DBForm.CreateForm9(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm9_1(form domain.Form9_1Request) error {

	res := []domain.Table9_1{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

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

	err = u.DBForm.CreateForm9_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm10(form domain.Form10Request) error {

	res := []domain.Table10{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, domain.Table10{
			EncuestadorID:        encuestador.ID,
			ZonaID:               zona.ID,
			DiscapacidadReferida: integrante.DiscapacidadReferida,
			DiscapacidadRecabada: integrante.DiscapacidadRecabada,
			TipoCondicion:        integrante.TipoCondicion,
		})
	}

	err = u.DBForm.CreateForm10(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm10_1(form domain.Form10_1Request) error {

	res := []domain.Table10_1{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, domain.Table10_1{
			EncuestadorID:    encuestador.ID,
			ZonaID:           zona.ID,
			SiDonde:          integrante.SiDonde,
			NoPorque:         integrante.NoPorque,
			TieneCertificado: integrante.TieneCertificado,
		})
	}

	err = u.DBForm.CreateForm10_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm11(form domain.Form11Request) error {

	res := []domain.Table11{}
	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, domain.Table11{
			EncuestadorID:  encuestador.ID,
			ZonaID:         zona.ID,
			NombreFarmaco:  integrante.NombreFarmaco,
			EsPrescrito:    integrante.EsPrescrito,
			MotivoReferido: integrante.MotivoReferido,
			MotivoRecabado: integrante.MotivoRecabado,
		})
	}

	err = u.DBForm.CreateForm11(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm1213(forms domain.Form1213Request) error {

	encuestador, err := u.DBUser.GetByEmail(forms.EncuestadorEmail)
	if err != nil {
		return err;
	}

	zona, err := u.DBZona.FirstOrCreateZona(forms.Zona)
	if err != nil {
		return err
	}

	err = u.CreateForm12(forms.Form12, *encuestador, *zona)
	if err != nil {
		return err
	}

	err = u.CreateForm13(forms.Form13, *encuestador, *zona)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm12(form domain.Form12, encuestador domain.User, zona domain.Zona) error {

	res := domain.Table12{}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona.ID
	res.HayProblema = form.HayProblema
	res.QueProblema = form.QueProblema

	err := u.DBForm.CreateForm12(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm13(form domain.Form13, encuestador domain.User, zona domain.Zona) error {

	res := domain.Table13{}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona.ID
	res.HayContaminacion = form.HayContaminacion
	res.CualEs = form.CualEs

	err := u.DBForm.CreateForm13(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *formService) CreateForm14(form domain.Form14Request) error {

	res := []domain.Table14{}

	encuestador, err := u.DBUser.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.DBZona.FirstOrCreateZona(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, domain.Table14{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona.ID,
			Tiempo:        integrante.Tiempo,
		})
	}

	err = u.DBForm.CreateForm14(res)
	if err != nil {
		return err
	}

	return nil
}
