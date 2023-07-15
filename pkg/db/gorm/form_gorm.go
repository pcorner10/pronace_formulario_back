package db

import (
	"fmt"
	"pronaces_back/pkg/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormStore struct {
	db *gorm.DB
}

func NewGormStore() (domain.FormDB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetString("database.port"))

	dbHandler, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sql, err := dbHandler.DB()
	if err != nil {
		panic(err)
	}

	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxIdleTime(30)
	sql.SetConnMaxLifetime(60)

	fmt.Println("Successfully connected to database!")
	return &gormStore{db: dbHandler}, nil
}

func (r *gormStore) CreateForm0(table0 []domain.Table0) error {

	err := r.db.Create(&table0).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm1(table1 domain.Table1) error {

	err := r.db.Create(&table1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm2(table2 domain.Table2) error {

	err := r.db.Create(&table2).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm3(table3 domain.Table3) error {

	err := r.db.Create(&table3).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm4(table4 []domain.Table4) error {

	err := r.db.Create(&table4).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm5(table5 []domain.Table5) error {

	err := r.db.Create(&table5).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm6(table6 []domain.Table6) error {

	err := r.db.Create(&table6).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm7(table7 []domain.Table7) error {

	err := r.db.Create(&table7).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm8(table8 []domain.Table8) error {

	err := r.db.Create(&table8).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm8_1(table8_1 []domain.Table8_1) error {

	err := r.db.Create(&table8_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm9(table9 []domain.Table9) error {

	err := r.db.Create(&table9).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm9_1(table9_1 []domain.Table9_1) error {

	err := r.db.Create(&table9_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm10(table10 []domain.Table10) error {

	err := r.db.Create(&table10).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm10_1(table10_1 []domain.Table10_1) error {

	err := r.db.Create(&table10_1).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm11(table11 []domain.Table11) error {

	err := r.db.Create(&table11).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm12(table12 domain.Table12) error {

	err := r.db.Create(&table12).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm13(table13 domain.Table13) error {

	err := r.db.Create(&table13).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm14(table14 []domain.Table14) error {

	err := r.db.Create(&table14).Error
	if err != nil {
		return err
	}

	return nil
}
