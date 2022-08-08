package auth

import (
	"github.com/depri11/lolipad/src/helpers"
	"github.com/depri11/lolipad/src/interfaces"
	"github.com/depri11/lolipad/src/models"
)

type service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *service {
	return &service{repository}
}

func (s *service) Register(user *models.User) *helpers.Response {
	result, err := s.repository.Register(user)
	if err != nil {
		return helpers.ResponseJSON("error", 401, err.Error())
	}

	return helpers.ResponseJSON("error", 401, result)
}
