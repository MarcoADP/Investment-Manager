package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type AtivoRentabilidadeRepository struct {
	db *gorm.DB
}

func NewAtivoRentabilidadeRepository(database *gorm.DB) *AtivoRentabilidadeRepository {
	return &AtivoRentabilidadeRepository{db: database}
}

func (r *AtivoRentabilidadeRepository) GetRentabilidadesByCodigo(codigo string) ([]model.AtivoRentabilidade, error) {
	var rentabilidades []model.AtivoRentabilidade
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").Find(&rentabilidades).Error; err != nil {
		return []model.AtivoRentabilidade{}, err
	}
	return rentabilidades, nil
}

func (r *AtivoRentabilidadeRepository) GetRentabilidadeByAtivoInformacao(ativoInformacaoId uint) (model.AtivoRentabilidade, error) {
	var ativoRentabilidade model.AtivoRentabilidade
	if err := r.db.Where("ativo_informacao_id = ?", ativoInformacaoId).First(&ativoRentabilidade).Error; err != nil {
		return model.AtivoRentabilidade{}, err
	}
	return ativoRentabilidade, nil
}

func (r *AtivoRentabilidadeRepository) GetRentabilidadeMoreRecentByCodigo(codigo string) (model.AtivoRentabilidade, error) {
	var rentabilidade model.AtivoRentabilidade
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").First(&rentabilidade).Error; err != nil {
		return model.AtivoRentabilidade{}, err
	}
	return rentabilidade, nil
}

func (r *AtivoRentabilidadeRepository) CreateRentabilidade(rentabilidade model.AtivoRentabilidade) (model.AtivoRentabilidade, error) {
	if err := r.db.Create(&rentabilidade).Error; err != nil {
		return model.AtivoRentabilidade{}, err
	}
	return rentabilidade, nil
}

func (r *AtivoRentabilidadeRepository) DeleteRentabilidade(id uint) error {
	if err := r.db.Delete(&model.AtivoRentabilidade{}, id).Error; err != nil {
		return err
	}
	return nil
}
