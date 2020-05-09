package services

import (
	"encoding/base64"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/internals/hash"
	"nikan.dev/pronto/payloads"
)

type userService struct {
	repository contracts.IUserRepository
	deps dependencies.CommonDependencies
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
	user, err := service.repository.Create(entities.User{
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		Password:   base64.StdEncoding.EncodeToString(hashBytes),
		Email:      payload.Email,
		Status:     entities.PENDING,
		Role:       defaultUser.(string),
	})
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
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

func NewUserService(deps dependencies.CommonDependencies,repository contracts.IUserRepository) userService {
	return userService{
		repository,
		deps,
	}
}
