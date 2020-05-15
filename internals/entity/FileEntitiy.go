package entity

import "io"

type FileEntity struct {
	Reader io.Reader
	Name string
	Size int64
	Mime string
}