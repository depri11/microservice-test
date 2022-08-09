package repository

import (
	"github.com/depri11/lolipad/auth/src/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByMsisdn(msisdn string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("msisdn = ?", msisdn).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) SaveData(data *entity.User) (*entity.User, error) {
	err := r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
