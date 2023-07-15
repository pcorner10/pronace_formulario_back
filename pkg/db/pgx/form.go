package db

import (
	"fmt"
	"pronaces_back/pkg/domain"
)

func (r *psxStore) CreateForm0(req interface{}) error {

	form, ok := req.([]domain.Table0)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table0")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table0 (encuestador_id, zona_id, sexo, edad, parentesco, 
		escolaridad, trabaja, horas_dia, dias_semana, ocupacion, fumador,
		tiempo_ciudad) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	for _, f := range form {

		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.Sexo, f.Edad, f.Parentesco,
			f.Escolaridad, f.Trabaja, f.HorasPorDia, f.DiasPorSemana, f.Ocupacion, f.Fumador,
			f.TiempoCiudad)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// TODO
func (r *psxStore) CreateForm123(req interface{}) error {
	return nil
}

func (r *psxStore) CreateForm1(req interface{}) error {
	form, ok := req.(domain.Table1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table1")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table1 (encuestador_id, zona_id, tiene_electricidad, fuentes_agua, 
		tiene_gas, eliminacion_desechos, material_techo, material_piso, material_paredes,
		num_cuartos, olores_desagradables) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err = tx.Exec(query, form.EncuestadorID, form.ZonaID, form.TieneElectricidad, form.FuentesAgua,
		form.TieneGas, form.EliminacionDesechos, form.MaterialTecho, form.MaterialPiso, form.MaterialParedes,
		form.NumCuartos, form.OloresDesagradables)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm2(req interface{}) error {
	form, ok := req.(domain.Table2)
	if !ok {
		return fmt.Errorf("error casting to domain.Table2")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table2 (encuestador_id, zona_id, llave_interior, garrafon, 
			llave_comunitaria) VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(query, form.EncuestadorID, form.ZonaID, form.LlaveInterior, form.Garrafon,
		form.LlaveComunitaria)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm3(req interface{}) error {
	form, ok := req.(domain.Table3)
	if !ok {
		return fmt.Errorf("error casting to domain.Table3")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table3 (encuestador_id, zona_id, imss, issste, seguro_popular, 
			privado, ninguno) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = tx.Exec(query, form.EncuestadorID, form.ZonaID, form.Imss, form.Issste, form.SeguroPopular,
		form.Privado, form.Ninguno)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm4(req interface{}) error {
	form, ok := req.([]domain.Table4)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table4")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table4 (encuestador_id, zona_id, enfermedad_referida, enfermedad_recabada)
		VALUES ($1, $2, $3, $4)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.EnfermedadReferida, f.EnfermedadRecabada)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm5(req interface{}) error {
	form, ok := req.([]domain.Table5)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table5")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table5 (encuestador_id, zona_id, enfermedad_referida, enfermedad_recabada)
		VALUES ($1, $2, $3, $4)`
	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.EnfermedadReferida, f.EnfermedadRecabada)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *psxStore) CreateForm6(req interface{}) error {
	form, ok := req.([]domain.Table6)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table6")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table6 (encuestador_id, zona_id, sexo, edad_fallecimiento, año_fallecimiento,
			causa_referida, causa_recabada)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.Sexo, f.EdadFallecimiento, f.AñoFallecimiento,
			f.CausaReferida, f.CausaRecabada)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm7(req interface{}) error {
	form, ok := req.([]domain.Table7)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table7")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table7 (encuestador_id, zona_id, sexo, edad_detection, año_detection,
			tipo_referido, tipo_recabado)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.Sexo, f.EdadDetection, f.AñoDetection, 
			f.TipoReferido, f.TipoRecabado)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm8(req interface{}) error {
	form, ok := req.([]domain.Table8)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table8")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table8 (encuestador_id, zona_id, año_nacimiento_perdida, en_curso, 
			tipo_parto, tuvo_complicaciones, complicacion_referida, complicacion_recabada)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.AñoNacimientoPerdida, f.EnCurso, f.TipoParto,
			f.TuvoComplicaciones, f.ComplicacionReferida, f.ComplicacionRecabada)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm8_1(req interface{}) error {
	form, ok := req.([]domain.Table8_1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table8_1")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table8_1 (encuestador_id, zona_id, año_perdida, trimestre)
		VALUES ($1, $2, $3, $4)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.AñoPerdida, f.Trimestre)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm9(req interface{}) error {
	form, ok := req.([]domain.Table9)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table9")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table9 (encuestador_id, zona_id, bajo_peso, prematuro, malformacion,
			malformacion_referida, malformacion_recabada)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.BajoPeso, f.Prematuro, f.Malformacion,
			f.MalformacionReferida, f.MalformacionRecabada)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm9_1(req interface{}) error {
	form, ok := req.([]domain.Table9_1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table9_1")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table9 (encuestador_id, zona_id, año, problema_referido, problema_recabado)
		VALUES ($1, $2, $3, $4, $5)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.Año, f.ProblemaReferido, f.ProblemaRecabado)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm10(req interface{}) error {
	form, ok := req.([]domain.Table10)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table10")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table10 (encuestador_id, zona_id, discapacidad_referida, 
			discapacidad_recabada, tipo_condicion)
		VALUES ($1, $2, $3, $4, $5)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.DiscapacidadReferida, f.DiscapacidadRecabada,
			f.TipoCondicion)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm10_1(req interface{}) error {
	form, ok := req.([]domain.Table10_1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table10_1")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table10_1 (encuestador_id, zona_id, si_donde, no_porque, tiene_certificado)
		VALUES ($1, $2, $3, $4, $5)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.SiDonde, f.NoPorque, f.TieneCertificado)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm11(req interface{}) error {
	form, ok := req.([]domain.Table11)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table11")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table11 (encuestador_id, zona_id, nombre_farmaco, es_prescrito, 
			motivo_referido, motivo_recabado)
		VALUES ($1, $2, $3, $4, $5, $6)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.NombreFarmaco, f.EsPrescrito, f.MotivoReferido,
			f.MotivoRecabado)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// TODO
func (r *psxStore) CreateForm1213(req interface{}) error {
	return nil
}

func (r *psxStore) CreateForm12(req interface{}) error {
	form, ok := req.(domain.Table12)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table12")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table13 (encuestador_id, zona_id, hay_problema, que_problema)
		VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(query, form.EncuestadorID, form.ZonaID, form.HayProblema, form.QueProblema)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm13(req interface{}) error {
	form, ok := req.(domain.Table13)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table13")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table13 (encuestador_id, zona_id, hay_contaminacion, cual_es)
		VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(query, form.EncuestadorID, form.ZonaID, form.HayContaminacion, form.CualEs)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *psxStore) CreateForm14(req interface{}) error {
	form, ok := req.([]domain.Table14)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table14")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO table6 (encuestador_id, zona_id, tiempo)
		VALUES ($1, $2, $3)`

	for _, f := range form {
		_, err := tx.Exec(query, f.EncuestadorID, f.ZonaID, f.Tiempo)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
