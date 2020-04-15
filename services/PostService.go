package services

import (
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
)

type postService struct {
	repository contracts.IPostRepository
}

func NewPostService(deps dependencies.CommonDependencies,repository contracts.IPostRepository) postService {
	return postService{
		repository,
	}
}

func (p postService) List() ([]entities.Post, error) {
	return p.repository.List()
}
