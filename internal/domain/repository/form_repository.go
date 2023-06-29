package repository

import "pronaces_back/internal/domain/entity"

// FormRepository define las operaciones para acceder y manipular los datos de los formularios
type FormRepository interface {
	CreateForm0(table0 []entity.Table0) error
	CreateForm1(table1 entity.Table1) (*entity.Table1, error)
	CreateForm2(table2 entity.Table2) (*entity.Table2, error)
	CreateForm3(table3 entity.Table3) (*entity.Table3, error)
	CreateForm4(table4 entity.Table4) (*entity.Table4, error)
	CreateForm5(table5 entity.Table5) (*entity.Table5, error)
	CreateForm6(table6 entity.Table6) (*entity.Table6, error)
	CreateForm7(table7 entity.Table7) (*entity.Table7, error)
	CreateForm8(table8 entity.Table8) (*entity.Table8, error)
	CreateForm8_1(table8_1 entity.Table8_1) (*entity.Table8_1, error)
	CreateForm9(table9 entity.Table9) (*entity.Table9, error)
	CreateForm9_1(table9 entity.Table9_1) (*entity.Table9_1, error)
	CreateForm10(table10 entity.Table10) (*entity.Table10, error)
	CreateForm10_1(table10 entity.Table10_1) (*entity.Table10_1, error)
	CreateForm11(table11 entity.Table11) (*entity.Table11, error)
	CreateForm12(table12 entity.Table12) (*entity.Table12, error)
	CreateForm13(table13 entity.Table13) (*entity.Table13, error)
	CreateForm14(table14 entity.Table14) (*entity.Table14, error)
}
