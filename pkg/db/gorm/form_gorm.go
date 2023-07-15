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

func (r *gormStore) CreateForm0(req interface{}) error {

	form, ok := req.([]domain.Table0)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table0")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO
func (r *gormStore) CreateForm123(req interface{}) error {
	return nil
}

func (r *gormStore) CreateForm1(req interface{}) error {

	form, ok := req.(domain.Table1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table1")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm2(req interface{}) error {

	form, ok := req.(domain.Table2)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table2")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm3(req interface{}) error {

	form, ok := req.(domain.Table3)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table3")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm4(req interface{}) error {

	form, ok := req.([]domain.Table4)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table4")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm5(req interface{}) error {

	form, ok := req.([]domain.Table5)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table5")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm6(req interface{}) error {

	form, ok := req.([]domain.Table6)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table6")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm7(req interface{}) error {

	form, ok := req.([]domain.Table7)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table7")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm8(req interface{}) error {

	form, ok := req.([]domain.Table8)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table8")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm8_1(req interface{}) error {

	form, ok := req.([]domain.Table8_1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table8_1")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm9(req interface{}) error {

	form, ok := req.([]domain.Table9)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table9")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm9_1(req interface{}) error {

	form, ok := req.([]domain.Table9_1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table9_1")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm10(req interface{}) error {

	form, ok := req.([]domain.Table10)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table10")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm10_1(req interface{}) error {

	form, ok := req.([]domain.Table10_1)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table10_1")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm11(req interface{}) error {

	form, ok := req.([]domain.Table11)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table11")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO
func (r *gormStore) CreateForm1213(req interface{}) error {
	return nil
}

func (r *gormStore) CreateForm12(req interface{}) error {

	form, ok := req.(domain.Table12)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table12")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm13(req interface{}) error {

	form, ok := req.(domain.Table13)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table13")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gormStore) CreateForm14(req interface{}) error {

	form, ok := req.([]domain.Table14)
	if !ok {
		return fmt.Errorf("error casting to []domain.Table14")
	}

	err := r.db.Create(&form).Error
	if err != nil {
		return err
	}

	return nil
}
