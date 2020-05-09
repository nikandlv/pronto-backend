package entities

import (
	"nikan.dev/pronto/internals/entity"
)

type Setting struct {
	entity.SimpleEntity
	Name  string
	Value string
}
