package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type AtivoValuationRepository struct {
	db *gorm.DB
}

func NewAtivoValuationRepository(database *gorm.DB) *AtivoValuationRepository {
	return &AtivoValuationRepository{db: database}
}

func (r *AtivoValuationRepository) GetValuationsByCodigo(codigo string) ([]model.AtivoValuation, error) {
	var valuations []model.AtivoValuation
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").Find(&valuations).Error; err != nil {
		return []model.AtivoValuation{}, err
	}
	return valuations, nil
}

func (r *AtivoValuationRepository) GetValuationByAtivoInformacao(ativoInformacaoId uint) (model.AtivoValuation, error) {
	var valuation model.AtivoValuation
	if err := r.db.Where("ativo_informacao_id = ?", ativoInformacaoId).First(&valuation).Error; err != nil {
		return model.AtivoValuation{}, err
	}
	return valuation, nil
}

func (r *AtivoValuationRepository) GetValuationMoreRecentByCodigo(codigo string) (model.AtivoValuation, error) {
	var valuation model.AtivoValuation
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").First(&valuation).Error; err != nil {
		return model.AtivoValuation{}, err
	}
	return valuation, nil
}

func (r *AtivoValuationRepository) CreateValuation(valuation model.AtivoValuation) (model.AtivoValuation, error) {
	if err := r.db.Create(&valuation).Error; err != nil {
		return model.AtivoValuation{}, err
	}
	return valuation, nil
}

func (r *AtivoValuationRepository) DeleteValuation(id uint) error {
	if err := r.db.Delete(&model.AtivoValuation{}, id).Error; err != nil {
		return err
	}
	return nil
}
