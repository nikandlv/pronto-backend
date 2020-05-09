package gorm

import (
	"github.com/jinzhu/gorm"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
)

type postRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies
}

func (p postRepository) List() ([]entities.Post, error) {
	panic("implement me")
}

func NewPostRepository(deps dependencies.CommonDependencies, pool interface{}) postRepository {
	return postRepository{pool: pool.(*gorm.DB), deps: deps}
}
