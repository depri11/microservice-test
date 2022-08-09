package service

import (
	"github.com/depri11/lolipad/logistic/src/helpers"
	"github.com/depri11/lolipad/logistic/src/interfaces"
)

type service struct {
	repository interfaces.LogisticRepository
}

func NewService(repository interfaces.LogisticRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() (*helpers.Response, error) {
	logistic, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	response := helpers.ResponseJSON("Success", 200, "OK", logistic)
	return response, nil
}
