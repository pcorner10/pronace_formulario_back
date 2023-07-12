package db

import (
	"pronaces_back/pkg/domain"

	"gorm.io/gorm"
)

type formGorm struct {
	db *gorm.DB
}

func NewFormGorm(db *gorm.DB) (domain.FormDB, error) {
	return &formGorm{
		db: db,
	}, nil
}

func (r *formGorm) CreateForm0(table0 []domain.Table0) error {

	err := r.db.Create(&table0).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm1(table1 domain.Table1) error {

	err := r.db.Create(&table1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm2(table2 domain.Table2) error {

	err := r.db.Create(&table2).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm3(table3 domain.Table3) error {

	err := r.db.Create(&table3).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm4(table4 []domain.Table4) error {

	err := r.db.Create(&table4).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm5(table5 []domain.Table5) error {

	err := r.db.Create(&table5).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm6(table6 []domain.Table6) error {

	err := r.db.Create(&table6).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm7(table7 []domain.Table7) error {

	err := r.db.Create(&table7).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm8(table8 []domain.Table8) error {

	err := r.db.Create(&table8).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm8_1(table8_1 []domain.Table8_1) error {

	err := r.db.Create(&table8_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm9(table9 []domain.Table9) error {

	err := r.db.Create(&table9).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm9_1(table9_1 []domain.Table9_1) error {

	err := r.db.Create(&table9_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm10(table10 []domain.Table10) error {

	err := r.db.Create(&table10).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm10_1(table10_1 []domain.Table10_1) error {

	err := r.db.Create(&table10_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm11(table11 []domain.Table11) error {

	err := r.db.Create(&table11).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm12(table12 domain.Table12) error {

	err := r.db.Create(&table12).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm13(table13 domain.Table13) error {

	err := r.db.Create(&table13).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formGorm) CreateForm14(table14 []domain.Table14) error {

	err := r.db.Create(&table14).Error
	if err != nil {
		return err
	}

	return nil
}
