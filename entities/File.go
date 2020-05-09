package entities

import (
	"nikan.dev/pronto/internals/entity"
)

type File struct {
	entity.BaseEntity
	Name  string
	Type  string
	Path string
	Storage string
	Size int64
}
