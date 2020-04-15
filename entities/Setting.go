package entities

import (
	"nikan.dev/pronto/internals/entity"
)

type Setting struct {
	entity.BaseEntity
	Name  string
	Value string
}
