package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/payloads"
)

type ISettingsService interface {
	Set(payloads payloads.SettingSetBatchPayload) ([]entities.Setting, error)
}
