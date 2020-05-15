package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/entity"
	"nikan.dev/pronto/payloads"
)

type IStorageService interface {
	List(user entities.User,payload payloads.PaginationPayload) (payloads.ChunkPayload, error)
	Store(user entities.User,file entity.FileEntity, payload payloads.StoreFilePayload) (entities.File, error)
	Delete(user entities.User,payload payloads.DeleteFilePayload) (entities.File, error)
}
