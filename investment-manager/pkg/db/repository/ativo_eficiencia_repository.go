package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type AtivoEficienciaRepository struct {
	db *gorm.DB
}

func NewAtivoEficienciaRepository(database *gorm.DB) *AtivoEficienciaRepository {
	return &AtivoEficienciaRepository{db: database}
}

func (r *AtivoEficienciaRepository) GetEficienciasByCodigo(codigo string) ([]model.AtivoEficiencia, error) {
	var eficiencias []model.AtivoEficiencia
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").Find(&eficiencias).Error; err != nil {
		return []model.AtivoEficiencia{}, err
	}
	return eficiencias, nil
}

func (r *AtivoEficienciaRepository) GetEficienciaByAtivoInformacao(ativoInformacaoId uint) (model.AtivoEficiencia, error) {
	var ativoEficiencia model.AtivoEficiencia
	if err := r.db.Where("ativo_informacao_id = ?", ativoInformacaoId).First(&ativoEficiencia).Error; err != nil {
		return model.AtivoEficiencia{}, err
	}
	return ativoEficiencia, nil
}

func (r *AtivoEficienciaRepository) GetEficienciaMoreRecentByCodigo(codigo string) (model.AtivoEficiencia, error) {
	var eficiencia model.AtivoEficiencia
	if err := r.db.Where("codigo = ?", codigo).Order("data_calculo DESC").First(&eficiencia).Error; err != nil {
		return model.AtivoEficiencia{}, err
	}
	return eficiencia, nil
}

func (r *AtivoEficienciaRepository) CreateEficiencia(eficiencia model.AtivoEficiencia) (model.AtivoEficiencia, error) {
	if err := r.db.Create(&eficiencia).Error; err != nil {
		return model.AtivoEficiencia{}, err
	}
	return eficiencia, nil
}

func (r *AtivoEficienciaRepository) DeleteEficiencia(id uint) error {
	if err := r.db.Delete(&model.AtivoEficiencia{}, id).Error; err != nil {
		return err
	}
	return nil
}
