package interfaces

import "github.com/depri11/lolipad/auth/src/entity"

type UserRepository interface {
	SaveData(data *entity.User) (*entity.User, error)
}
