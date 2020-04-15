package contracts

import "nikan.dev/pronto/entities"

type IPostService interface {
	List() ([]entities.Post, error)

}
