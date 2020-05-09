package contracts

import (
	"nikan.dev/pronto/entities"
)

type ISettingsRepository interface {
	SetBatch(setting []entities.Setting) ([]entities.Setting, error)
	List() ([]entities.Setting, error)
}
