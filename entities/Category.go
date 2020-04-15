package entities

import "nikan.dev/pronto/internals/entity"

type Category struct {
	entity.BaseEntity
	Title string
	Slug string
	ParentID uint
	Parent *Category
}

