package entities

import "nikan.dev/pronto/internals/entity"

type Post struct {
	entity.BaseEntity
	Title      string
	Slug       string
	Content    string
	CategoryID uint
	UserID     uint
	Category   *Category
	User       *User
}
