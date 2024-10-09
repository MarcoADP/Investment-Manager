package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type SetorRepository struct {
	db *gorm.DB
}

func NewSetorRepository(database *gorm.DB) *SetorRepository {
	return &SetorRepository{db: database}
}

func (r *SetorRepository) GetAllSetores() ([]model.Setor, error) {
	var setores []model.Setor
	if err := r.db.Find(&setores).Error; err != nil {
		return nil, err
	}
	return setores, nil
}

func (r *SetorRepository) GetSetorByID(id uint) (model.Setor, error) {
	var setor model.Setor
	if err := r.db.First(&setor, id).Error; err != nil {
		return model.Setor{}, err
	}
	return setor, nil
}

func (r *SetorRepository) GetSetorByNome(nome string) (model.Setor, error) {
	var setor model.Setor
	if err := r.db.Where("nome = ?", nome).First(&setor).Error; err != nil {
		return model.Setor{}, err
	}
	return setor, nil
}

func (r *SetorRepository) CreateSetor(setor model.Setor) (model.Setor, error) {
	if err := r.db.Create(&setor).Error; err != nil {
		return model.Setor{}, err
	}
	return setor, nil
}
