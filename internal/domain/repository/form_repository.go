package repository

import "pronaces_back/internal/domain/entity"

// FormRepository define las operaciones para acceder y manipular los datos de los formularios
type FormRepository interface {
	CreateForm(table0 entity.Table0) (*entity.Table0, error)
}
