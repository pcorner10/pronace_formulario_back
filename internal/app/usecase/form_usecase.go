package usecase

import (
	apprepository "pronaces_back/internal/app/repository"
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"
	"strconv"
)

type FormUseCase struct {
	formRepository repository.FormRepository
	userRepo       repository.UserRepository
	zonaRepo       repository.ZonaRepository
}

func NewFormUseCase(formRepository repository.FormRepository) *FormUseCase {
	return &FormUseCase{
		formRepository: formRepository,
	}
}

func (u *FormUseCase) CreateForm0(integrantes apprepository.Form0Request) error {

	form := []entity.Table0{}

	encuestador, err := u.userRepo.GetByEmail(integrantes.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(integrantes.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range integrantes.Integrante {

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

		form = append(form, entity.Table0{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona,
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

	err = u.formRepository.CreateForm0(form)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm1(form apprepository.Form1) error {

	res := entity.Table1{}

	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	NumCuartos, err := strconv.Atoi(form.NumCuartos)
	if err != nil {
		return err
	}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona
	res.TieneElectricidad = form.TieneElectricidad
	res.FuentesAgua = form.FuentesAgua
	res.TieneGas = form.TieneGas
	res.EliminacionDesechos = form.EliminacionDesechos
	res.MaterialParedes = form.MaterialParedes
	res.MaterialPiso = form.MaterialPiso
	res.MaterialTecho = form.MaterialTecho
	res.NumCuartos = NumCuartos
	res.OloresDesagradables = form.OloresDesagradables

	err = u.formRepository.CreateForm1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm2(form apprepository.Form2) error {

	res := entity.Table2{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona
	res.LlaveInterior = form.LlaveInterior
	res.Garrafon = form.Garrafon
	res.LlaveComunitaria = form.LlaveComunitaria

	err = u.formRepository.CreateForm2(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm3(form apprepository.Form3) error {

	res := entity.Table3{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona
	res.Imss = form.Imss
	res.Issste = form.Issste
	res.SeguroPopular = form.SeguroPopular
	res.Privado = form.Privado
	res.Ninguno = form.Ninguno

	err = u.formRepository.CreateForm3(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm4(form apprepository.Form4Request) error {

	res := []entity.Table4{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {
		res = append(res, entity.Table4{
			EncuestadorID:      encuestador.ID,
			ZonaID:             zona,
			EnfermedadReferida: integrante.EnfermedadReferida,
			EnfermedadRecabada: integrante.EnfermedadRecabada,
		})
	}

	err = u.formRepository.CreateForm4(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm5(form apprepository.Form5Request) error {

	res := []entity.Table5{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {
		res = append(res, entity.Table5{
			EncuestadorID:      encuestador.ID,
			ZonaID:             zona,
			EnfermedadReferida: integrante.EnfermedadReferida,
			EnfermedadRecabada: integrante.EnfermedadRecabada,
		})
	}

	err = u.formRepository.CreateForm5(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm6(form apprepository.Form6Request) error {

	res := []entity.Table6{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
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

		res = append(res, entity.Table6{
			EncuestadorID:     encuestador.ID,
			ZonaID:            zona,
			Sexo:              integrante.Sexo,
			AñoFallecimiento:  AñoFallecimiento,
			EdadFallecimiento: EdadFallecimiento,
			CausaReferida:     integrante.CausaReferida,
			CausaRecabada:     integrante.CausaRecabada,
		})
	}

	err = u.formRepository.CreateForm6(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm7(form apprepository.Form7Request) error {

	res := []entity.Table7{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		EdadDetection, err := strconv.Atoi(integrante.EdadDetection)
		if err != nil {
			return err
		}

		AñoDetection, err := strconv.Atoi(integrante.AñoDetection)
		if err != nil {
			return err
		}

		res = append(res, entity.Table7{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona,
			Sexo:          integrante.Sexo,
			EdadDetection: EdadDetection,
			AñoDetection:  AñoDetection,
			TipoReferido:  integrante.TipoReferido,
			TipoRecabado:  integrante.TipoRecabado,
		})

	}

	err = u.formRepository.CreateForm7(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm8(form apprepository.Form8Request) error {

	res := []entity.Table8{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		AñoNacimientoPerdida, err := strconv.Atoi(integrante.AñoNacimientoPerdida)
		if err != nil {
			return err
		}

		res = append(res, entity.Table8{
			EncuestadorID:        encuestador.ID,
			ZonaID:               zona,
			AñoNacimientoPerdida: AñoNacimientoPerdida,
			EnCurso:              integrante.EnCurso,
			TipoParto:            integrante.TipoParto,
			TuvoComplicaciones:   integrante.TuvoComplicaciones,
			ComplicacionReferida: integrante.ComplicacionReferida,
			ComplicacionRecabada: integrante.ComplicacionRecabada,
		})

	}

	err = u.formRepository.CreateForm8(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm8_1(form apprepository.Form8_1Request) error {

	res := []entity.Table8_1{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		AñoPerdida, err := strconv.Atoi(integrante.AñoPerdida)
		if err != nil {
			return err
		}

		res = append(res, entity.Table8_1{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona,
			AñoPerdida:    AñoPerdida,
			Trimestre:     integrante.Trimestre,
		})
	}

	err = u.formRepository.CreateForm8_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm9(form apprepository.Form9Request) error {

	res := []entity.Table9{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {
		res = append(res, entity.Table9{
			EncuestadorID:        encuestador.ID,
			ZonaID:               zona,
			BajoPeso:             integrante.BajoPeso,
			Prematuro:            integrante.Prematuro,
			Malformacion:         integrante.Malformacion,
			MalformacionReferida: integrante.MalformacionReferida,
			MalformacionRecabada: integrante.MalformacionRecabada,
		})
	}

	err = u.formRepository.CreateForm9(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm9_1(form apprepository.Form9_1Request) error {

	res := []entity.Table9_1{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		año, err := strconv.Atoi(integrante.Año)
		if err != nil {
			return err
		}

		res = append(res, entity.Table9_1{
			EncuestadorID:    encuestador.ID,
			ZonaID:           zona,
			Año:              año,
			ProblemaReferido: integrante.ProblemaReferido,
			ProblemaRecabado: integrante.ProblemaRecabado,
		})
	}

	err = u.formRepository.CreateForm9_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm10(form apprepository.Form10Request) error {

	res := []entity.Table10{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, entity.Table10{
			EncuestadorID:        encuestador.ID,
			ZonaID:               zona,
			DiscapacidadReferida: integrante.DiscapacidadReferida,
			DiscapacidadRecabada: integrante.DiscapacidadRecabada,
			TipoCondicion:        integrante.TipoCondicion,
		})
	}

	err = u.formRepository.CreateForm10(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm10_1(form apprepository.Form10_1Request) error {

	res := []entity.Table10_1{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, entity.Table10_1{
			EncuestadorID:    encuestador.ID,
			ZonaID:           zona,
			SiDonde:          integrante.SiDonde,
			NoPorque:         integrante.NoPorque,
			TieneCertificado: integrante.TieneCertificado,
		})
	}

	err = u.formRepository.CreateForm10_1(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm11(form apprepository.Form11Request) error {

	res := []entity.Table11{}
	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, entity.Table11{
			EncuestadorID:  encuestador.ID,
			ZonaID:         zona,
			NombreFarmaco:  integrante.NombreFarmaco,
			EsPrescrito:    integrante.EsPrescrito,
			MotivoReferido: integrante.MotivoReferido,
			MotivoRecabado: integrante.MotivoRecabado,
		})
	}

	err = u.formRepository.CreateForm11(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm12(form apprepository.Form12) error {

	res := entity.Table12{}

	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona
	res.HayProblema = form.HayProblema
	res.QueProblema = form.QueProblema

	err = u.formRepository.CreateForm12(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm13(form apprepository.Form13) error {

	res := entity.Table13{}

	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	res.EncuestadorID = encuestador.ID
	res.ZonaID = zona
	res.HayContaminacion = form.HayContaminacion
	res.CualEs = form.CualEs

	err = u.formRepository.CreateForm13(res)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm14(form apprepository.Form14Request) error {

	res := []entity.Table14{}

	encuestador, err := u.userRepo.GetByEmail(form.EncuestadorEmail)
	if err != nil {
		return err
	}

	zona, err := u.zonaRepo.GetZonaID(form.Zona)
	if err != nil {
		return err
	}

	for _, integrante := range form.Integrante {

		res = append(res, entity.Table14{
			EncuestadorID: encuestador.ID,
			ZonaID:        zona,
			Tiempo:        integrante.Tiempo,
		})
	}

	err = u.formRepository.CreateForm14(res)
	if err != nil {
		return err
	}

	return nil
}
