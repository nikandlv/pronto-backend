package mysql

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type applicationRepository struct {
	pool internalContracts.IDataPool
	deps dependencies.CommonDependencies
}

func NewApplicationRepository(deps dependencies.CommonDependencies, pool internalContracts.IDataPool) applicationRepository {
	return applicationRepository{pool, deps}
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
