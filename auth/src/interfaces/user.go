package interfaces

import (
	"github.com/depri11/lolipad/auth/src/entity"
	"github.com/depri11/lolipad/auth/src/helpers"
	"github.com/depri11/lolipad/auth/src/models"
)

type UserRepository interface {
	SaveData(data *entity.User) (*entity.User, error)
}

type UserService interface {
	Register(input *models.InputRegisterUser) (*helpers.Response, error)
}
