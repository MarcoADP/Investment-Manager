package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type GrahamFormulaRepository struct {
	db *gorm.DB
}

func NewGrahamFormulaRepository(database *gorm.DB) *GrahamFormulaRepository {
	return &GrahamFormulaRepository{db: database}
}

func (r *GrahamFormulaRepository) GetGrahamFormulaByCodigo(codigo string) ([]model.GrahamFormula, error) {
	var calculos []model.GrahamFormula
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").Find(&calculos).Error; err != nil {
		return []model.GrahamFormula{}, err
	}
	return calculos, nil
}

func (r *GrahamFormulaRepository) GetGrahamFormulaMoreRecentByCodigo(codigo string) (model.GrahamFormula, error) {
	var grahamFormula model.GrahamFormula
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").First(&grahamFormula).Error; err != nil {
		return model.GrahamFormula{}, err
	}
	return grahamFormula, nil
}

func (r *GrahamFormulaRepository) CreateGrahamFormula(grahamFormula model.GrahamFormula) (model.GrahamFormula, error) {
	if err := r.db.Create(&grahamFormula).Error; err != nil {
		return model.GrahamFormula{}, err
	}
	return grahamFormula, nil
}

func (r *GrahamFormulaRepository) DeleteGrahamFormula(id uint) error {
	if err := r.db.Delete(&model.GrahamFormula{}, id).Error; err != nil {
		return err
	}
	return nil
}
