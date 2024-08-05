package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type AcaoBrRepository struct {
	db *gorm.DB
}

func NewAcaoBrRepository(database *gorm.DB) *AcaoBrRepository {
	return &AcaoBrRepository{db: database}
}

func (r *AcaoBrRepository) GetAllAcaoBrs() ([]model.AcaoBr, error) {
	var acoes []model.AcaoBr
	if err := r.db.Find(&acoes).Error; err != nil {
		return nil, err
	}
	return acoes, nil
}

func (r *AcaoBrRepository) GetAcaoBrByID(id uint) (model.AcaoBr, error) {
	var acao model.AcaoBr
	if err := r.db.First(&acao, id).Error; err != nil {
		return model.AcaoBr{}, err
	}
	return acao, nil
}

func (r *AcaoBrRepository) CreateAcaoBr(acao model.AcaoBr) (model.AcaoBr, error) {
	if err := r.db.Create(&acao).Error; err != nil {
		return model.AcaoBr{}, err
	}
	return acao, nil
}

func (r *AcaoBrRepository) UpdateAcaoBr(acao model.AcaoBr) (model.AcaoBr, error) {
	if err := r.db.Save(&acao).Error; err != nil {
		return model.AcaoBr{}, err
	}
	return acao, nil
}

func (r *AcaoBrRepository) DeleteAcaoBr(id uint) error {
	if err := r.db.Delete(&model.AcaoBr{}, id).Error; err != nil {
		return err
	}
	return nil
}
