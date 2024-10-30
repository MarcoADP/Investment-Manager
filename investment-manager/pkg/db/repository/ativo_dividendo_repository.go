package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type AtivoDividendoRepository struct {
	db *gorm.DB
}

func NewAtivoDividendoRepository(database *gorm.DB) *AtivoDividendoRepository {
	return &AtivoDividendoRepository{db: database}
}

func (r *AtivoDividendoRepository) GetDividendosByCodigo(codigo string) ([]model.AtivoDividendo, error) {
	var dividendos []model.AtivoDividendo
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").Find(&dividendos).Error; err != nil {
		return []model.AtivoDividendo{}, err
	}
	return dividendos, nil
}

func (r *AtivoDividendoRepository) GetDividendoByAtivoInformacao(ativoInformacaoId uint) (model.AtivoDividendo, error) {
	var ativoDividendo model.AtivoDividendo
	if err := r.db.Where("ativo_informacao_id = ?", ativoInformacaoId).First(&ativoDividendo).Error; err != nil {
		return model.AtivoDividendo{}, err
	}
	return ativoDividendo, nil
}

func (r *AtivoDividendoRepository) GetDividendoMoreRecentByCodigo(codigo string) (model.AtivoDividendo, error) {
	var dividendo model.AtivoDividendo
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").First(&dividendo).Error; err != nil {
		return model.AtivoDividendo{}, err
	}
	return dividendo, nil
}

func (r *AtivoDividendoRepository) CreateDividendo(dividendo model.AtivoDividendo) (model.AtivoDividendo, error) {
	if err := r.db.Create(&dividendo).Error; err != nil {
		return model.AtivoDividendo{}, err
	}
	return dividendo, nil
}

func (r *AtivoDividendoRepository) DeleteDividendo(id uint) error {
	if err := r.db.Delete(&model.AtivoDividendo{}, id).Error; err != nil {
		return err
	}
	return nil
}
