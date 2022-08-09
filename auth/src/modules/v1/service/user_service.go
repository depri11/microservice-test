package service

import (
	"github.com/depri11/lolipad/auth/src/entity"
	"github.com/depri11/lolipad/auth/src/helpers"
	"github.com/depri11/lolipad/auth/src/interfaces"
	"github.com/depri11/lolipad/auth/src/models"
)

type service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *service {
	return &service{repository}
}

func (s *service) Register(input *models.InputRegisterUser) (*helpers.Response, error) {
	var user entity.User
	user.Msisdn = input.Msisdn
	user.Name = input.Name
	user.Username = input.Username

	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	data, err := s.repository.SaveData(&user)
	if err != nil {
		return nil, err
	}

	response := helpers.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}
