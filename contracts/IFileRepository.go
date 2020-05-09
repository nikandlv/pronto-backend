package contracts

import "nikan.dev/pronto/entities"

type IFileRepository interface
{
	Get() (entities.File, error)
	Has(id uint) (bool, error)
	HasWithPath(id uint) (bool, error)
	Delete(id uint) (bool, error)
	Insert(file entities.File) (bool,error)
}