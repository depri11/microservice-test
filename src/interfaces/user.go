package interfaces

import "github.com/depri11/lolipad/src/models"

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
	Register(user *models.User) (int64, error)
}
