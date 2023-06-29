package usecase

import (
	apprepository "pronaces_back/internal/app/repository"
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/domain/repository"
)

type FormUseCase struct {
	formRepository repository.FormRepository
}

func NewFormUseCase(formRepository repository.FormRepository) *FormUseCase {
	return &FormUseCase{
		formRepository: formRepository,
	}
}

func (u *FormUseCase) CreateForm0(integrantes apprepository.Form0Request) error {

	form := make([]entity.Table0, len(integrantes.Integrante))
	copy(form, integrantes.Integrante)

	err := u.formRepository.CreateForm0(form)
	if err != nil {
		return err
	}

	return nil
}

func (u *FormUseCase) CreateForm1(form entity.Table1) (*entity.Table1, error) {

	res := &entity.Table1{}
	res, err := u.formRepository.CreateForm1(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm2(form entity.Table2) (*entity.Table2, error) {

	res := &entity.Table2{}
	res, err := u.formRepository.CreateForm2(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm3(form entity.Table3) (*entity.Table3, error) {

	res := &entity.Table3{}
	res, err := u.formRepository.CreateForm3(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm4(form entity.Table4) (*entity.Table4, error) {

	res := &entity.Table4{}
	res, err := u.formRepository.CreateForm4(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm5(form entity.Table5) (*entity.Table5, error) {

	res := &entity.Table5{}
	res, err := u.formRepository.CreateForm5(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm6(form entity.Table6) (*entity.Table6, error) {

	res := &entity.Table6{}
	res, err := u.formRepository.CreateForm6(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm7(form entity.Table7) (*entity.Table7, error) {

	res := &entity.Table7{}
	res, err := u.formRepository.CreateForm7(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm8(form entity.Table8) (*entity.Table8, error) {

	res := &entity.Table8{}
	res, err := u.formRepository.CreateForm8(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm8_1(form entity.Table8_1) (*entity.Table8_1, error) {

	res := &entity.Table8_1{}
	res, err := u.formRepository.CreateForm8_1(form)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm9(table9 entity.Table9) (*entity.Table9, error) {

	res := &entity.Table9{}
	res, err := u.formRepository.CreateForm9(table9)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm9_1(table9_1 entity.Table9_1) (*entity.Table9_1, error) {

	res := &entity.Table9_1{}
	res, err := u.formRepository.CreateForm9_1(table9_1)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm10(table10 entity.Table10) (*entity.Table10, error) {

	res := &entity.Table10{}
	res, err := u.formRepository.CreateForm10(table10)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm10_1(table10_1 entity.Table10_1) (*entity.Table10_1, error) {

	res := &entity.Table10_1{}
	res, err := u.formRepository.CreateForm10_1(table10_1)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm11(table11 entity.Table11) (*entity.Table11, error) {

	res := &entity.Table11{}
	res, err := u.formRepository.CreateForm11(table11)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm12(table12 entity.Table12) (*entity.Table12, error) {

	res := &entity.Table12{}
	res, err := u.formRepository.CreateForm12(table12)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm13(table13 entity.Table13) (*entity.Table13, error) {

	res := &entity.Table13{}
	res, err := u.formRepository.CreateForm13(table13)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FormUseCase) CreateForm14(table14 entity.Table14) (*entity.Table14, error) {

	res := &entity.Table14{}
	res, err := u.formRepository.CreateForm14(table14)
	if err != nil {
		return nil, err
	}

	return res, nil
}
