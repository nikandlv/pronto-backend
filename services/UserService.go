package services

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/internals/entity"
	"nikan.dev/pronto/internals/hash"
	"nikan.dev/pronto/payloads"
	"path/filepath"
)

type userService struct {
	repository contracts.IUserRepository
	storageDependencies dependencies.StorageDependencies
	deps dependencies.CommonDependencies
}

func (service userService) UpdateAvatar(user entities.User, file entity.FileEntity) (entities.User, error) {
	file.Name = fmt.Sprintf("%v%v", uuid.New() ,filepath.Ext(file.Name))
	avatarFolder, err := service.deps.Configuration.Get("AvatarsFolder")
	if err != nil {
		return user, err
	}
	err = service.storageDependencies.Storage.Store(file,avatarFolder.(string) )
	if err != nil {
		return user, err
	}
	user.Avatar = file.Name
	return service.repository.Update(user)
}

func (service userService) Update(user entities.User, payload payloads.UserUpdatePayload) (entities.User, error) {
	if err := payload.Validate(service.deps.Validator); err != nil {
		return user, err
	}
	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.Email = payload.Email
    return  service.repository.Update(user)
}

func (service userService) UpdatePredefinedAvatar(user entities.User, payload payloads.UserUpdatePredefinedAvatarPayload) (entities.User, error) {
	if err := payload.Validate(service.deps.Validator); err != nil {
		return user, err
	}
	user.Avatar = payload.Avatar
	return service.repository.Update(user)
}

func (service userService) Register(payload payloads.UserRegisterPayload) (entities.User, error) {
	if err := payload.Validate(service.deps.Validator); err != nil {
		return entities.User{},err
	}

	exists, err := service.repository.EmailExists(payload.Email)
	if err != nil {
		return entities.User{}, exceptions.ServerError
	}
	if exists {
		return entities.User{}, exceptions.UserAlreadyExists
	}

	defaultUser, err := service.deps.Configuration.Get("DefaultRole")
	if err != nil {
		return entities.User{},err
	}
	hashBytes, err := hash.Generate(service.deps.Configuration, payload.Password)
	if err != nil {
		return entities.User{}, err
	}
	return service.repository.Create(entities.User{
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		Password:   base64.StdEncoding.EncodeToString(hashBytes),
		Email:      payload.Email,
		Role:       defaultUser.(string),
	})
}

func (service userService) CheckCredentials(payload payloads.UserCredentialsPayload) (entities.User, error) {
	if err := payload.Validate(service.deps.Validator); err != nil {
		return entities.User{},err
	}
	user, err := service.repository.GetByEmail(payload.Email)
	if err != nil {
		return entities.User{},err
	}
	hashBytes, err := hash.Generate(service.deps.Configuration, payload.Password)
	if err != nil {
		return entities.User{},err
	}
	if user.Password != base64.StdEncoding.EncodeToString(hashBytes) {
		return entities.User{}, exceptions.NotAllowed
	}
	return user, nil
}

func (service userService) Get(payload payloads.UserIDOnlyPayload) (entities.User, error) {
	if err := payload.Validate(service.deps.Validator); err != nil {
		return entities.User{}, err
	}
	return service.repository.Get(payload.ID)
}

func NewUserService(deps dependencies.CommonDependencies,storageDependencies dependencies.StorageDependencies,repository contracts.IUserRepository) userService {
	return userService{
		repository,
		storageDependencies,
		deps,
	}
}
