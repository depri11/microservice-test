package repository

import (
	"github.com/depri11/lolipad/logistic/src/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() (*entity.Logistics, error) {
	var logistic entity.Logistics
	err := r.db.Find(&logistic).Error
	if err != nil {
		return nil, err
	}

	return &logistic, nil
}
