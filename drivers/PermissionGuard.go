package drivers

import (
	"fmt"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
)

type permissionGuard struct {
	deps dependencies.CommonDependencies
}

func (p permissionGuard) IsAuthorized(user entities.User, action string) error {
	hasPermission:= p.deps.Configuration.Contains(fmt.Sprintf("%v.%v.%v","Permissions",user.Role, action))
	if !hasPermission {
		return exceptions.NotAllowed
	}
	return nil
}

func NewPermissionGuard(deps dependencies.CommonDependencies) permissionGuard {
	return permissionGuard{
		deps,
	}
}