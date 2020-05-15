package internalContracts

import "nikan.dev/pronto/entities"

type IPermissionGuard interface {

	IsAuthorized(user entities.User, action string) error
}
