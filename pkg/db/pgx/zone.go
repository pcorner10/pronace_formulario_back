package db

import (
	"database/sql"
	"pronaces_back/pkg/domain"
	"time"
)

func (r *psxStore) FirstOrCreateZona(zona domain.Zona) (*domain.Zona, error) {
	// implementa la logica para obtener una zona de la base de datos
	// si no existe, la crea y la devuelve usando Psx

	var zonaDB domain.Zona

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
	SELECT * FROM zona WHERE municipio = $1 AND localidad = $2 AND ageb = $3 AND manzana = $4 AND lote = $5`

	err = tx.QueryRow(query, zona.Municipio, zona.Localidad, zona.Ageb, zona.Manzana,
		zona.Lote).Scan(&zonaDB.ID, &zonaDB.Municipio, &zonaDB.Localidad, &zonaDB.Ageb,
		&zonaDB.Manzana, &zonaDB.Lote, &zonaDB.CreatedAt, &zonaDB.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			insertQuery := `
				INSERT INTO zona (municipio, localidad, ageb, manzana, lote) 
				VALUES ($1, $2, $3, $4, $5) RETURNING id`

			err = tx.QueryRow(insertQuery, zona.Municipio, zona.Localidad, zona.Ageb, zona.Manzana, zona.Lote, time.Now(), time.Now()).
				Scan(&zonaDB.ID)

			if err != nil {
				tx.Rollback()
				return nil, err
			}
		} else if err != nil {
			// An unexpected error occurred, rollback the transaction and return the error.
			tx.Rollback()
			return nil, err
		}
	}

	// Commit the transaction since everything went well.
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &zonaDB, nil
}

func (r *psxStore) CreateZona(zona domain.Zona) (*domain.Zona, error) {

	var zonaDB domain.Zona

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	insertQuery := `
	INSERT INTO zona (municipio, localidad, ageb, manzana, lote, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = r.db.QueryRow(insertQuery, zona.Municipio, zona.Localidad, zona.Ageb, zona.Manzana, zona.Lote, time.Now(), time.Now()).
		Scan(&zonaDB.ID, &zonaDB.Municipio, &zonaDB.Localidad, &zonaDB.Ageb, &zonaDB.Manzana, &zonaDB.Lote, &zonaDB.CreatedAt, &zonaDB.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction since everything went well.
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &zonaDB, nil
}

func (r *psxStore) GetZonaByID(zonaID uint) (*domain.Zona, error) {

	var zona domain.Zona
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
	SELECT id FROM zona WHERE 
		municipio = $1 AND 
		localidad = $2 AND 
		ageb = $3 AND
		manzana = $4 AND
		lote = $5`

	err = tx.QueryRow(query, zona.Municipio, zona.Localidad, zona.Ageb, zona.Manzana,
		zona.Lote).Scan(&zona.ID, &zona.Municipio, &zona.Localidad, &zona.Ageb, &zona.Manzana,
		&zona.Lote, &zona.CreatedAt, &zona.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &zona, nil
}

func (r *psxStore) GetZonaID(zona domain.Zona) (uint, error) {

	var zonaID uint

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `
	SELECT id FROM zona WHERE 
		municipio = $1 AND 
		localidad = $2 AND 
		ageb = $3 AND
		manzana = $4 AND
		lote = $5`

	err = tx.QueryRow(query, zona.Municipio, zona.Localidad, zona.Ageb, zona.Manzana,
		zona.Lote).Scan(&zonaID)

	if err == sql.ErrNoRows {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	return zonaID, nil
}
