package service

import (
	"fmt"

	"github.com/depri11/lolipad/logistic/src/entity"
	"github.com/depri11/lolipad/logistic/src/helpers"
	"github.com/depri11/lolipad/logistic/src/interfaces"
	"github.com/depri11/lolipad/logistic/src/models"
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

func (s *service) GetData(input *models.InputLogisticData) (*helpers.Response, error) {
	var logistic entity.Logistic
	logistic.OriginName = input.OriginName
	logistic.DestinationName = input.DestinationName
	fmt.Println(logistic)

	data, err := s.repository.GetData(&logistic)
	if err != nil {
		return nil, err
	}

	response := helpers.ResponseJSON("Success", 200, "OK", data)
	return response, nil
}
