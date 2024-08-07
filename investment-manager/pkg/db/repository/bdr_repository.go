package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type BdrRepository struct {
	db *gorm.DB
}

func NewBdrRepository(database *gorm.DB) *BdrRepository {
	return &BdrRepository{db: database}
}

func (r *BdrRepository) GetAllBdrs() ([]model.Bdr, error) {
	var bdrs []model.Bdr
	if err := r.db.Find(&bdrs).Error; err != nil {
		return nil, err
	}
	return bdrs, nil
}

func (r *BdrRepository) GetBdrByID(id uint) (model.Bdr, error) {
	var bdr model.Bdr
	if err := r.db.First(&bdr, id).Error; err != nil {
		return model.Bdr{}, err
	}
	return bdr, nil
}

func (r *BdrRepository) CreateBdr(bdr model.Bdr) (model.Bdr, error) {
	if err := r.db.Create(&bdr).Error; err != nil {
		return model.Bdr{}, err
	}
	return bdr, nil
}

func (r *BdrRepository) UpdateBdr(bdr model.Bdr) (model.Bdr, error) {
	if err := r.db.Save(&bdr).Error; err != nil {
		return model.Bdr{}, err
	}
	return bdr, nil
}

func (r *BdrRepository) DeleteBdr(id uint) error {
	if err := r.db.Delete(&model.Bdr{}, id).Error; err != nil {
		return err
	}
	return nil
}
