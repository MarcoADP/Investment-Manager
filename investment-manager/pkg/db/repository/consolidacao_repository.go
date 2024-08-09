package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type ConsolidacaoRepository struct {
	db *gorm.DB
}

func NewConsolidacaoRepository(database *gorm.DB) *ConsolidacaoRepository {
	return &ConsolidacaoRepository{db: database}
}

func (r *ConsolidacaoRepository) GetAllConsolidacoes() ([]model.Consolidacao, error) {
	var consolidacoes []model.Consolidacao
	if err := r.db.Find(&consolidacoes).Error; err != nil {
		return nil, err
	}
	return consolidacoes, nil
}

func (r *ConsolidacaoRepository) GetConsolidacaoByCodigo(codigo string) (model.Consolidacao, error) {
	var consolidacao model.Consolidacao
	if err := r.db.Where("codigo = ?", codigo).First(&consolidacao).Error; err != nil {
		return model.Consolidacao{}, err
	}
	return consolidacao, nil
}

func (r *ConsolidacaoRepository) CreateConsolidacao(consolidacao model.Consolidacao) (model.Consolidacao, error) {
	if err := r.db.Create(&consolidacao).Error; err != nil {
		return model.Consolidacao{}, err
	}
	return consolidacao, nil
}

func (r *ConsolidacaoRepository) UpdateConsolidacao(consolidacao model.Consolidacao) (model.Consolidacao, error) {
	if err := r.db.Save(&consolidacao).Error; err != nil {
		return model.Consolidacao{}, err
	}
	return consolidacao, nil
}
