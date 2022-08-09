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
	_, err := s.repository.GetByMsisdn("62" + input.Msisdn)
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
	user.Msisdn = "62" + input.Msisdn
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

func (s *service) Login(input *models.InputLoginUser) (*helpers.Response, error) {
	user, err := s.repository.GetByMsisdn("62" + input.Msisdn)
	if err != nil {
		response := helpers.ResponseJSON("Bad Request", 401, "error", "msisdn not found")
		return response, nil
	}

	if !helpers.CheckPassword(user.Password, input.Password) {
		response := helpers.ResponseJSON("Bad Request", 401, "error", "please check your msisdn/password")
		return response, nil
	}

	token := helpers.NewToken(user.ID.String())
	tokens, err := token.Create()
	if err != nil {
		response := helpers.ResponseJSON("Internal Server Error", 500, "error", nil)
		return response, nil
	}

	response := helpers.ResponseJSON("Success", 200, "OK", tokens)
	return response, nil
}

func (s *service) CheckToken(token string) (*helpers.Response, error) {
	check, err := helpers.CheckToken(token)
	if err != nil {
		response := helpers.ResponseJSON("Bad Request", 401, "error", err.Error())
		return response, nil
	}

	response := helpers.ResponseJSON("Success", 200, "OK", check)
	return response, nil
}
