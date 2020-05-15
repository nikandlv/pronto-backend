package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/entity"
	"nikan.dev/pronto/payloads"
)

type IUserService interface {
	Register(payload payloads.UserRegisterPayload) (entities.User, error)
	CheckCredentials(payload payloads.UserCredentialsPayload) (entities.User, error)
	Get(payload payloads.UserIDOnlyPayload) (entities.User, error)
	Update(user entities.User,payload payloads.UserUpdatePayload) (entities.User, error)
	UpdateAvatar(user entities.User,file entity.FileEntity) (entities.User, error)
	UpdatePredefinedAvatar(user entities.User,payload payloads.UserUpdatePredefinedAvatarPayload) (entities.User, error)
}