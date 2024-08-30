package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type AtivoInformacaoRepository struct {
	db *gorm.DB
}

func NewAtivoInformacaoRepository(database *gorm.DB) *AtivoInformacaoRepository {
	return &AtivoInformacaoRepository{db: database}
}

func (r *AtivoInformacaoRepository) GetInformacoesByCodigo(codigo string) ([]model.AtivoInformacao, error) {
	var informacoes []model.AtivoInformacao
	if err := r.db.Where("codigo = ?", codigo).Order("data_informacao DESC").Find(&informacoes).Error; err != nil {
		return []model.AtivoInformacao{}, err
	}
	return informacoes, nil
}

func (r *AtivoInformacaoRepository) GetInformacaoMoreRecentByCodigo(codigo string) (model.AtivoInformacao, error) {
	var informacao model.AtivoInformacao
	if err := r.db.Where("codigo = ?", codigo).Order("data_informacao DESC").First(&informacao).Error; err != nil {
		return model.AtivoInformacao{}, err
	}
	return informacao, nil
}

func (r *AtivoInformacaoRepository) CreateInformacao(informacao model.AtivoInformacao) (model.AtivoInformacao, error) {
	if err := r.db.Create(&informacao).Error; err != nil {
		return model.AtivoInformacao{}, err
	}
	return informacao, nil
}

func (r *AtivoInformacaoRepository) DeleteInformacao(id uint) error {
	if err := r.db.Delete(&model.AtivoInformacao{}, id).Error; err != nil {
		return err
	}
	return nil
}
