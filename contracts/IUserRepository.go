package contracts

import (
	"nikan.dev/pronto/entities"
)

type IUserRepository interface {
	Get(ID uint) (entities.User, error)
	GetByEmail(Email string) (entities.User, error)
	EmailExists(Email string) (bool, error)
	Create(user entities.User) (entities.User, error)
}
