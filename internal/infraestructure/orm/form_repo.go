package orm

import (
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"

	"gorm.io/gorm"
)

type formRepository struct {
	db *gorm.DB
}

func NewFormRepository(db *gorm.DB) repository.FormRepository {
	return &formRepository{
		db: db,
	}
}

func (r *formRepository) CreateForm0(table0 []entity.Table0) error {

	err := r.db.Create(&table0).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm1(table1 entity.Table1) error {

	err := r.db.Create(&table1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm2(table2 entity.Table2) error {

	err := r.db.Create(&table2).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm3(table3 entity.Table3) error {

	err := r.db.Create(&table3).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm4(table4 entity.Table4) error {

	err := r.db.Create(&table4).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm5(table5 entity.Table5) error {

	err := r.db.Create(&table5).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm6(table6 entity.Table6) error {

	err := r.db.Create(&table6).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm7(table7 entity.Table7) error {

	err := r.db.Create(&table7).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm8(table8 entity.Table8) error {

	err := r.db.Create(&table8).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm8_1(table8_1 entity.Table8_1) error {

	err := r.db.Create(&table8_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm9(table9 entity.Table9) error {

	err := r.db.Create(&table9).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm9_1(table9_1 entity.Table9_1) error {

	err := r.db.Create(&table9_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm10(table10 entity.Table10) error {

	err := r.db.Create(&table10).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm10_1(table10_1 entity.Table10_1) error {

	err := r.db.Create(&table10_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm11(table11 entity.Table11) error {

	err := r.db.Create(&table11).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm12(table12 entity.Table12) error {

	err := r.db.Create(&table12).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm13(table13 entity.Table13) error {

	err := r.db.Create(&table13).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *formRepository) CreateForm14(table14 entity.Table14) error {

	err := r.db.Create(&table14).Error
	if err != nil {
		return err
	}

	return nil
}
