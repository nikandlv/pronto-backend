package gorm

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
)

type fileRepository struct {
	deps dependencies.CommonDependencies
}

func (f fileRepository) Get() (entities.File, error) {
	panic("implement me")
}

func (f fileRepository) Has(id uint) (bool, error) {
	panic("implement me")
}

func (f fileRepository) HasWithPath(id uint) (bool, error) {
	panic("implement me")
}

func (f fileRepository) Delete(id uint) (bool, error) {
	panic("implement me")
}

func (f fileRepository) Insert(file entities.File) (bool, error) {
	panic("implement me")
}

func NewFileRepository(deps dependencies.CommonDependencies) fileRepository {
	return fileRepository{
		deps,
	}
}