package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/payloads"
)

type ISettingRepository interface {
	Set(setting entities.Setting) entities.Setting
	SetBatch(setting []entities.Setting) []entities.Setting
	Get(payload payloads.SettingGetPayload) entities.Setting
	List() []entities.Setting
}
