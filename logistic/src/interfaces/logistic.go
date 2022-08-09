package interfaces

import (
	"github.com/depri11/lolipad/logistic/src/entity"
	"github.com/depri11/lolipad/logistic/src/helpers"
	"github.com/depri11/lolipad/logistic/src/models"
)

type LogisticRepository interface {
	GetAll() (*entity.Logistics, error)
	GetData(data *entity.Logistic) (*entity.Logistic, error)
}

type LogisticService interface {
	FindAll() (*helpers.Response, error)
	GetData(input *models.InputLogisticData) (*helpers.Response, error)
}
