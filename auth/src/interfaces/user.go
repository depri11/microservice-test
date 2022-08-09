package interfaces

import (
	"github.com/depri11/lolipad/auth/src/models"
)

type UserRepository interface {
	SaveData(data models.ResponseRegisterUser) (*models.ResponseRegisterUser, error)
}
