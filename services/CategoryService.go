package services

import (
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type categoryService struct {
	repository contracts.ICategoryRepository
	deps dependencies.CommonDependencies
}

func NewCategoryService(deps dependencies.CommonDependencies,repository contracts.ICategoryRepository) categoryService {
	return categoryService{
		repository,
		deps,
	}
}

func (c categoryService) List() ([]entities.Category, error) {
	return c.repository.List()
}

func (c categoryService) Create(payload payloads.CategoryCreatePayload) (entities.Category, error) {
	if err := c.deps.Validator.Validate(payload); err != nil {
		return entities.Category{}, err
	}
	exists, err := c.repository.SlugExists(payload.Slug)
	if err != nil {
		return entities.Category{},err
	}
	 if exists {
		return entities.Category{}, exceptions.CategoryAlreadyExists
	 }
	if payload.ParentID > 0 {
		exists, err := c.repository.Exists(payload.ParentID)
		if err != nil {
			return entities.Category{},err
		}
		if !exists {
			return entities.Category{}, exceptions.CategoryParentNotFound
		}
	}
	 return c.repository.Create(entities.Category{
		Title:      payload.Title,
		Slug:       payload.Slug,
		ParentID:   payload.ParentID,
	})
}
