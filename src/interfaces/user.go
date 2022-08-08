package interfaces

import "github.com/depri11/lolipad/src/models"

type UserRepository interface {
	Register(user *models.User) (int64, error)
}
