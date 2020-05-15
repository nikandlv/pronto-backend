package entities

import (
	"nikan.dev/pronto/internals/entity"
)

type File struct {
	entity.BaseEntity
	Name  string
	Mime  string
	Scope string
	Size int64
}
