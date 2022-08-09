package interfaces

import (
	"github.com/depri11/lolipad/logistic/src/entity"
	"github.com/depri11/lolipad/logistic/src/helpers"
)

type LogisticRepository interface {
	GetAll() (*entity.Logistics, error)
}

type LogisticService interface {
	FindAll() (*helpers.Response, error)
}
