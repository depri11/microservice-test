package service

import (
	"github.com/depri11/lolipad/auth/src/entity"
	"github.com/depri11/lolipad/auth/src/helpers"
	"github.com/depri11/lolipad/auth/src/interfaces"
	"github.com/depri11/lolipad/auth/src/models"
	"github.com/google/uuid"
)

type service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *service {
	return &service{repository}
}

func (s *service) Register(input *models.InputRegisterUser) (*helpers.Response, error) {
	_, err := s.repository.GetByMsisdn(input.Msisdn)
	if err == nil {
		response := helpers.ResponseJSON("Failed", 401, "error", "msisdn already exists")
		return response, nil
	}

	_, err = s.repository.GetByUsername(input.Username)
	if err == nil {
		response := helpers.ResponseJSON("Failed", 401, "error", "username already exists")
		return response, nil
	}

	var user entity.User
	id := uuid.New()
	user.ID = id
	user.Msisdn = input.Msisdn
	user.Name = input.Name
	user.Username = input.Username

	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	dataUser, err := s.repository.SaveData(&user)
	if err != nil {
		return nil, err
	}

	var res models.ResponseRegisterUser
	res.ID = dataUser.ID
	res.Msisdn = dataUser.Msisdn
	res.Name = dataUser.Name
	res.Username = dataUser.Username

	response := helpers.ResponseJSON("Success", 200, "OK", res)
	return response, nil
}
