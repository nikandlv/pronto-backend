package mysql

import (
	"github.com/jinzhu/gorm"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type applicationRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies
}

func NewApplicationRepository(deps dependencies.CommonDependencies, pool interface{}) applicationRepository {
	return applicationRepository{pool: pool.(*gorm.DB), deps: deps}
}

func (repo applicationRepository) Info() (payloads.ApplicationInfoPayload, error) {
	ver, err := repo.deps.Configuration.Get("Version")
	if err != nil {
		return payloads.ApplicationInfoPayload{}, err
	}
	return payloads.ApplicationInfoPayload{Version: ver.(string)}, err
}

func (repo applicationRepository) Ping() (payloads.ApplicationPingPayload, error) {
	return payloads.ApplicationPingPayload{Ping:"Pong"}, nil
}
