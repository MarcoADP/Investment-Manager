package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type CarteiraAtivoRepository struct {
	db *gorm.DB
}

func NewCarteiraAtivoRepository(database *gorm.DB) *CarteiraAtivoRepository {
	return &CarteiraAtivoRepository{db: database}
}

func (r *CarteiraAtivoRepository) GetCarteiraAtivoByID(id uint) (model.CarteiraAtivo, error) {
	var ativo model.CarteiraAtivo
	if err := r.db.First(&ativo, id).Error; err != nil {
		return model.CarteiraAtivo{}, err
	}
	return ativo, nil
}

func (r *CarteiraAtivoRepository) GetAtivosByCarteiraID(carteiraId uint) ([]model.CarteiraAtivo, error) {
	var ativos []model.CarteiraAtivo
	if err := r.db.Where("carteira_id = ?", carteiraId).Find(&ativos).Error; err != nil {
		return []model.CarteiraAtivo{}, err
	}
	return ativos, nil
}

func (r *CarteiraAtivoRepository) GetAtivoByCodigoAndCarteiraID(codigo string, carteiraId uint) (model.CarteiraAtivo, error) {
	var ativo model.CarteiraAtivo
	if err := r.db.Where("codigo = ? AND carteira_id = ?", codigo, carteiraId).First(&ativo).Error; err != nil {
		return model.CarteiraAtivo{}, err
	}
	return ativo, nil
}

func (r *CarteiraAtivoRepository) CreateCarteiraAtivo(ativo model.CarteiraAtivo) (model.CarteiraAtivo, error) {
	if err := r.db.Create(&ativo).Error; err != nil {
		return model.CarteiraAtivo{}, err
	}
	return ativo, nil
}

func (r *CarteiraAtivoRepository) UpdateCarteiraAtivo(ativo model.CarteiraAtivo) (model.CarteiraAtivo, error) {
	if err := r.db.Save(&ativo).Error; err != nil {
		return model.CarteiraAtivo{}, err
	}
	return ativo, nil
}

func (r *CarteiraAtivoRepository) DeleteCarteiraAtivo(id uint) error {
	if err := r.db.Delete(&model.CarteiraAtivo{}, id).Error; err != nil {
		return err
	}
	return nil
}
