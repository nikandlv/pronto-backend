package services

import (
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type settingsService struct {
	deps dependencies.CommonDependencies
	repository contracts.ISettingsRepository
}

func (s settingsService) Set(payloads payloads.SettingSetBatchPayload) ([]entities.Setting, error) {
	if err := payloads.Validate(s.deps.Validator); err != nil {
		return []entities.Setting{},err
	}
	var settings []entities.Setting
	for _,setting := range payloads.Settings {
		settings = append(settings, entities.Setting{
			Name: setting.Name,
			Value: setting.Value,
		})
	}
	return s.repository.SetBatch(settings)
}

func NewSettingsService(deps dependencies.CommonDependencies,repository contracts.ISettingsRepository) settingsService {
	return settingsService{
		deps,
		repository,
	}
}
