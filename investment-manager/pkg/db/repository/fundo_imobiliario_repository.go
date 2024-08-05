package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type FundoImobiliarioRepository struct {
	db *gorm.DB
}

func NewFundoImobiliarioRepository(database *gorm.DB) *FundoImobiliarioRepository {
	return &FundoImobiliarioRepository{db: database}
}

func (r *FundoImobiliarioRepository) GetAllFundosImobiliarios() ([]model.FundoImobiliario, error) {
	var fundos []model.FundoImobiliario
	if err := r.db.Find(&fundos).Error; err != nil {
		return nil, err
	}
	return fundos, nil
}

func (r *FundoImobiliarioRepository) GetFundoImobiliarioByID(id uint) (model.FundoImobiliario, error) {
	var fundo model.FundoImobiliario
	if err := r.db.First(&fundo, id).Error; err != nil {
		return model.FundoImobiliario{}, err
	}
	return fundo, nil
}

func (r *FundoImobiliarioRepository) CreateFundoImobiliario(fundo model.FundoImobiliario) (model.FundoImobiliario, error) {
	if err := r.db.Create(&fundo).Error; err != nil {
		return model.FundoImobiliario{}, err
	}
	return fundo, nil
}

func (r *FundoImobiliarioRepository) UpdateFundoImobiliario(fundo model.FundoImobiliario) (model.FundoImobiliario, error) {
	if err := r.db.Save(&fundo).Error; err != nil {
		return model.FundoImobiliario{}, err
	}
	return fundo, nil
}

func (r *FundoImobiliarioRepository) DeleteFundoImobiliario(id uint) error {
	if err := r.db.Delete(&model.FundoImobiliario{}, id).Error; err != nil {
		return err
	}
	return nil
}
