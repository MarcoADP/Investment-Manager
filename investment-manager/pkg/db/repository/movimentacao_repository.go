package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type MovimentacaoRepository struct {
	db *gorm.DB
}

func NewMovimentacaoRepository(database *gorm.DB) *MovimentacaoRepository {
	return &MovimentacaoRepository{db: database}
}

func (r *MovimentacaoRepository) GetAllMovimentacoes() ([]model.Movimentacao, error) {
	var movimentacoes []model.Movimentacao
	if err := r.db.Find(&movimentacoes).Error; err != nil {
		return nil, err
	}
	return movimentacoes, nil
}

func (r *MovimentacaoRepository) GetMovimentacaoByID(id uint) (model.Movimentacao, error) {
	var movimentacao model.Movimentacao
	if err := r.db.First(&movimentacao, id).Error; err != nil {
		return model.Movimentacao{}, err
	}
	return movimentacao, nil
}

func (r *MovimentacaoRepository) GetMovimentacaoByCodigo(codigo string) ([]model.Movimentacao, error) {
	var movimentacoes []model.Movimentacao
	if err := r.db.Where("codigo = ?", codigo).Find(&movimentacoes).Error; err != nil {
		return []model.Movimentacao{}, err
	}
	return movimentacoes, nil
}

func (r *MovimentacaoRepository) CreateMovimentacao(movimentacao model.Movimentacao) (model.Movimentacao, error) {
	if err := r.db.Create(&movimentacao).Error; err != nil {
		return model.Movimentacao{}, err
	}
	return movimentacao, nil
}

func (r *MovimentacaoRepository) DeleteMovimentacao(id uint) error {
	if err := r.db.Delete(&model.Movimentacao{}, id).Error; err != nil {
		return err
	}
	return nil
}
