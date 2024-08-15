package repository

import (
	"github.com/MarcoADP/Investment-Manager/pkg/db/model"
	"gorm.io/gorm"
)

type CarteiraRepository struct {
	db *gorm.DB
}

func NewCarteiraRepository(database *gorm.DB) *CarteiraRepository {
	return &CarteiraRepository{db: database}
}

func (r *CarteiraRepository) GetAllCarteiras() ([]model.Carteira, error) {
	var carteiras []model.Carteira
	if err := r.db.Find(&carteiras).Error; err != nil {
		return nil, err
	}
	return carteiras, nil
}

func (r *CarteiraRepository) GetCarteiraByID(id uint) (model.Carteira, error) {
	var carteira model.Carteira
	if err := r.db.First(&carteira, id).Error; err != nil {
		return model.Carteira{}, err
	}
	return carteira, nil
}

func (r *CarteiraRepository) CreateCarteira(carteira model.Carteira) (model.Carteira, error) {
	if err := r.db.Create(&carteira).Error; err != nil {
		return model.Carteira{}, err
	}
	return carteira, nil
}

func (r *CarteiraRepository) UpdateCarteira(carteira model.Carteira) (model.Carteira, error) {
	if err := r.db.Save(&carteira).Error; err != nil {
		return model.Carteira{}, err
	}
	return carteira, nil
}

func (r *CarteiraRepository) DeleteCarteira(id uint) error {
	if err := r.db.Delete(&model.Carteira{}, id).Error; err != nil {
		return err
	}
	return nil
}
