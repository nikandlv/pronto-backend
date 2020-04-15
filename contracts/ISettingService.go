package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/payloads"
)

type ISettingService interface {
	Set(payload payloads.SettingSetPayload) entities.Setting
	Get(payload payloads.SettingGetPayload) entities.Setting
}
