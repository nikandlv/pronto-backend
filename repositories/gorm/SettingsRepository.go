package gorm

import (
	"github.com/jinzhu/gorm"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
)

type settingsRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies
}

func (s settingsRepository) SetBatch(setting []entities.Setting) ([]entities.Setting, error) {
	panic("implement me")
}

func (s settingsRepository) List() ([]entities.Setting, error) {
	var settings []entities.Setting
	s.pool.Select([]string{"id","name", "value"}).Find(&settings)
	return settings, nil
}

func NewSettingsRepository(deps dependencies.CommonDependencies, pool interface{}) settingsRepository {
	return settingsRepository{pool: pool.(*gorm.DB), deps: deps}
}
