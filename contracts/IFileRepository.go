package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/payloads"
)

type IFileRepository interface
{
	Get(ID uint) (entities.File,error)
	List(payload payloads.PaginationPayload) (payloads.ChunkPayload , error)
	Delete(file entities.File) (entities.File, error)
	Insert(file entities.File) (entities.File,error)
}