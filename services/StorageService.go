package services

import (
	"fmt"
	"github.com/google/uuid"
	"nikan.dev/pronto/entities"
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/internals/entity"
	"nikan.dev/pronto/payloads"
	"path/filepath"
)

type storageService struct {
	storageDependencies dependencies.StorageDependencies
	deps dependencies.CommonDependencies
	guard internalContracts.IPermissionGuard
}

func (s storageService) List(user entities.User,payload payloads.PaginationPayload) (payloads.ChunkPayload, error) {
	if err:= s.guard.IsAuthorized(user,"STORAGE_LIST"); err != nil {
		return payloads.ChunkPayload{},err
	}

	if err := payload.Validate(s.deps.Validator); err != nil {
		return payloads.ChunkPayload{}, err
	}
	return s.storageDependencies.Repository.List(payload)
}

func (s storageService) Store(user entities.User,file entity.FileEntity, payload payloads.StoreFilePayload) (entities.File, error) {
	if err:= s.guard.IsAuthorized(user,"STORAGE_STORE"); err != nil {
		return entities.File{},err
	}

	if err := payload.Validate(s.deps.Validator); err != nil {
		return entities.File{}, err
	}
	file.Name = fmt.Sprintf("%v%v", uuid.New() ,filepath.Ext(file.Name))
	storage, err := s.deps.Configuration.Get("UploadsFolder")
	if err != nil {
		return entities.File{},err
	}
	basePath := fmt.Sprintf("%v/%v", storage.(string),payload.Scope)

	if err:= s.storageDependencies.Storage.Store(file,basePath); err != nil {
		return entities.File{},err
	}

	entity := entities.File{
		Name: file.Name,
		Mime : file.Mime,
		Scope: payload.Scope,
		Size : file.Size,
	}

	return s.storageDependencies.Repository.Insert(entity)
}

func (s storageService) Delete(user entities.User,payload payloads.DeleteFilePayload) (entities.File, error) {
	if err:= s.guard.IsAuthorized(user,"STORAGE_DELETE"); err != nil {
		return entities.File{},err
	}
	if err := payload.Validate(s.deps.Validator); err != nil {
		return entities.File{}, err
	}
	file, getErr := s.storageDependencies.Repository.Get(payload.ID)
	if getErr != nil {
		return entities.File{}, getErr
	}
	return s.storageDependencies.Repository.Delete(file)
}

func NewStorageService(deps dependencies.CommonDependencies, storageDependencies dependencies.StorageDependencies, guard internalContracts.IPermissionGuard) storageService {
	return storageService{
		storageDependencies,
		deps,
		guard,
	}
}