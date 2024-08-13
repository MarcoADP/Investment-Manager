package repository

import (
	"time"

	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type CotacaoHistoricoRepository struct {
	db *gorm.DB
}

func NewCotacaoHistoricoRepository(database *gorm.DB) *CotacaoHistoricoRepository {
	return &CotacaoHistoricoRepository{db: database}
}

func (r *CotacaoHistoricoRepository) GetAllCotacoes() ([]model.CotacaoHistorico, error) {
	var cotacoes []model.CotacaoHistorico
	if err := r.db.Find(&cotacoes).Order("data_preco DESC").Error; err != nil {
		return nil, err
	}
	return cotacoes, nil
}

func (r *CotacaoHistoricoRepository) GetCotacoesByCodigo(codigo string) ([]model.CotacaoHistorico, error) {
	var cotacoes []model.CotacaoHistorico
	if err := r.db.Where("codigo = ?", codigo).Order("data_preco DESC").Find(&cotacoes).Error; err != nil {
		return []model.CotacaoHistorico{}, err
	}
	return cotacoes, nil
}

func (r *CotacaoHistoricoRepository) GetCotacoesByCodigoAndData(codigo string, data time.Time) (model.CotacaoHistorico, error) {
	var cotacao model.CotacaoHistorico
	if err := r.db.Where("codigo = ? and data_preco = ?", codigo, data).First(&cotacao).Error; err != nil {
		return model.CotacaoHistorico{}, err
	}
	return cotacao, nil
}

func (r *CotacaoHistoricoRepository) GetCotacaoMoreRecentByCodigo(codigo string) (model.CotacaoHistorico, error) {
	var cotacao model.CotacaoHistorico
	if err := r.db.Where("codigo = ?", codigo).Order("data_preco DESC").First(&cotacao).Error; err != nil {
		return model.CotacaoHistorico{}, err
	}
	return cotacao, nil
}

func (r *CotacaoHistoricoRepository) CreateCotacao(cotacao model.CotacaoHistorico) (model.CotacaoHistorico, error) {
	if err := r.db.Create(&cotacao).Error; err != nil {
		return model.CotacaoHistorico{}, err
	}
	return cotacao, nil
}

func (r *CotacaoHistoricoRepository) UpdateCotacaoHistorico(cotacao model.CotacaoHistorico) (model.CotacaoHistorico, error) {
	if err := r.db.Save(&cotacao).Error; err != nil {
		return model.CotacaoHistorico{}, err
	}
	return cotacao, nil
}

func (r *CotacaoHistoricoRepository) DeleteCotacaoHistorico(id uint) error {
	if err := r.db.Delete(&model.CotacaoHistorico{}, id).Error; err != nil {
		return err
	}
	return nil
}
