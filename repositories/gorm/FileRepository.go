package gorm

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type fileRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies

}

func (f fileRepository) Delete(file entities.File) (entities.File,error) {
	return file,f.pool.Delete(file).Error
}

func (f fileRepository) Get(ID uint) (entities.File, error) {
	var file entities.File
	if err := f.pool.First(&file,ID).Error; err != nil {
		return file, exceptions.FileNotFound
	}
	return file, nil
}

func (f fileRepository) List(payload payloads.PaginationPayload) (payloads.ChunkPayload, error) {
	var files []entities.File
	pager := pagination.Paging(&pagination.Param{
		DB:      f.pool,
		Page:    payload.Page,
		Limit:   payload.PageSize,
		OrderBy: []string{"id desc"},
	}, &files)
	return payloads.ChunkPayload(*pager), nil
}


func (f fileRepository) Insert(file entities.File) (entities.File, error ){
	if err := f.pool.Create(&file).Error; err != nil {
		return file, exceptions.ServerError
	}
	return file, nil
}

func NewFileRepository(deps dependencies.CommonDependencies, pool interface{}) fileRepository {
	return fileRepository{
		pool.(*gorm.DB),
		deps,
	}
}