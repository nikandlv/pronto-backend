package services

import (
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type applicationService struct {
	repository contracts.IApplicationRepository
	settingsRepository contracts.ISettingsRepository
}

func (service applicationService) Config() (payloads.ApplicationConfigPayload, error) {
	settings, err := service.settingsRepository.List()
	if err != nil {
		return payloads.ApplicationConfigPayload{}, err
	}
	return payloads.ApplicationConfigPayload{
		Version: "123",
		Settings: settings,
	}, nil
}

func NewApplicationService(deps dependencies.CommonDependencies,repository contracts.IApplicationRepository, settingsRepository contracts.ISettingsRepository) applicationService {
	return applicationService{
		repository, settingsRepository,
	}
}

func (service applicationService) Info() (payloads.ApplicationInfoPayload, error) {
	return service.repository.Info()
}

func (service applicationService) Ping() (payloads.ApplicationPingPayload, error) {
	return service.repository.Ping()
}
