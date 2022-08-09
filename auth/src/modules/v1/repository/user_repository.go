package repository

import (
	"github.com/depri11/lolipad/auth/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveData(data *models.ResponseRegisterUser) (*models.ResponseRegisterUser, error) {
	err := r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
