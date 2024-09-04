package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type DividendoRepository struct {
	db *gorm.DB
}

func NewDividendoRepository(database *gorm.DB) *DividendoRepository {
	return &DividendoRepository{db: database}
}

func (r *DividendoRepository) GetAllDividendos() ([]model.Dividendo, error) {
	var dividendos []model.Dividendo
	if err := r.db.Find(&dividendos).Error; err != nil {
		return nil, err
	}
	return dividendos, nil
}

func (r *DividendoRepository) GetDividendoByID(id uint) (model.Dividendo, error) {
	var dividendo model.Dividendo
	if err := r.db.First(&dividendo, id).Error; err != nil {
		return model.Dividendo{}, err
	}
	return dividendo, nil
}

func (r *DividendoRepository) CreateDividendo(dividendo model.Dividendo) (model.Dividendo, error) {
	if err := r.db.Create(&dividendo).Error; err != nil {
		return model.Dividendo{}, err
	}
	return dividendo, nil
}

func (r *DividendoRepository) UpdateDividendo(dividendo model.Dividendo) (model.Dividendo, error) {
	if err := r.db.Save(&dividendo).Error; err != nil {
		return model.Dividendo{}, err
	}
	return dividendo, nil
}

func (r *DividendoRepository) DeleteDividendo(id uint) error {
	if err := r.db.Delete(&model.Dividendo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *DividendoRepository) GetDividendosByCodigo(codigo string) ([]model.Dividendo, error) {
	var dividendos []model.Dividendo
	if err := r.db.Where("codigo = ?", codigo).Order("data_pagamento DESC").Find(&dividendos).Error; err != nil {
		return []model.Dividendo{}, err
	}
	return dividendos, nil
}
