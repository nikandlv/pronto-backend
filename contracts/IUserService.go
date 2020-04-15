package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/payloads"
)

type IUserService interface {
	Register(payload payloads.UserRegisterPayload) (entities.User, error)
	CheckCredentials(payload payloads.UserCredentialsPayload) (entities.User, error)
	Get(payload payloads.UserIDOnlyPayload) (entities.User, error)
}
