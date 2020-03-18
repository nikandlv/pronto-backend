package services

import (
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type applicationService struct {
	repository contracts.IApplicationRepository
}

func NewApplicationService(deps dependencies.CommonDependencies,repository contracts.IApplicationRepository) applicationService {
	return applicationService{
		repository,
	}
}

func (service applicationService) Info() (payloads.ApplicationInfoPayload, error) {
	return service.repository.Info()
}

func (service applicationService) Ping() (payloads.ApplicationPingPayload, error) {
	return service.repository.Ping()
}
