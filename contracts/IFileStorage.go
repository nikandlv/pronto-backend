package contracts

import "nikan.dev/pronto/internals/entity"

type IFileStorage interface
{
	Store(file entity.FileEntity, path string) error
	Remove(name string) (bool, error)
}