package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type AtivoEndividamentoRepository struct {
	db *gorm.DB
}

func NewAtivoEndividamentoRepository(database *gorm.DB) *AtivoEndividamentoRepository {
	return &AtivoEndividamentoRepository{db: database}
}

func (r *AtivoEndividamentoRepository) GetEndividamentosByCodigo(codigo string) ([]model.AtivoEndividamento, error) {
	var endividamentos []model.AtivoEndividamento
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").Find(&endividamentos).Error; err != nil {
		return []model.AtivoEndividamento{}, err
	}
	return endividamentos, nil
}

func (r *AtivoEndividamentoRepository) GetEndividamentoByAtivoInformacao(ativoInformacaoId uint) (model.AtivoEndividamento, error) {
	var endividamento model.AtivoEndividamento
	if err := r.db.Where("ativo_informacao_id = ?", ativoInformacaoId).First(&endividamento).Error; err != nil {
		return model.AtivoEndividamento{}, err
	}
	return endividamento, nil
}

func (r *AtivoEndividamentoRepository) GetEndividamentoMoreRecentByCodigo(codigo string) (model.AtivoEndividamento, error) {
	var endividamento model.AtivoEndividamento
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").First(&endividamento).Error; err != nil {
		return model.AtivoEndividamento{}, err
	}
	return endividamento, nil
}

func (r *AtivoEndividamentoRepository) CreateEndividamento(endividamento model.AtivoEndividamento) (model.AtivoEndividamento, error) {
	if err := r.db.Create(&endividamento).Error; err != nil {
		return model.AtivoEndividamento{}, err
	}
	return endividamento, nil
}

func (r *AtivoEndividamentoRepository) DeleteEndividamento(id uint) error {
	if err := r.db.Delete(&model.AtivoEndividamento{}, id).Error; err != nil {
		return err
	}
	return nil
}
