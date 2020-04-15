package contracts

import (
	"nikan.dev/pronto/entities"
)

type IPostRepository interface {
	List() ([]entities.Post, error)
	//Create(payload payloads.CategoryCreatePayload) (entities.Post, error)

}
