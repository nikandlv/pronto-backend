package contracts

import (
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/payloads"
)

type ICategoryService interface {
	List() ([]entities.Category, error)
	Create(payload payloads.CategoryCreatePayload) (entities.Category, error)
}
