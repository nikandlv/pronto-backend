package http

import (
	"github.com/labstack/echo/v4"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
	"strconv"
)

type storageEndpoint struct {
	deps dependencies.CommonDependencies
	service contracts.IStorageService
}

func (endpoint storageEndpoint) Boot(transport interface{}) {
	t := transport.(*echo.Group)
	group := drivers.JwtGroup(endpoint.deps.Configuration,t.Group("/storage"))
	group.GET("/list", endpoint.list).Name = "Storage.list"
	group.POST("/store", endpoint.store).Name = "Storage.store"
	group.DELETE("/file/:id", endpoint.delete).Name = "Storage.delete"
}

func NewStorageEndpoint(deps dependencies.CommonDependencies, service contracts.IStorageService) storageEndpoint {
	return storageEndpoint{deps,service }
}

func (endpoint storageEndpoint) list(ctx echo.Context) error {
	user:= drivers.GetUser(ctx)

	payload, err := drivers.RequestToPayload(ctx, new(payloads.PaginationPayload))
	if err != nil {
		return err
	}
	paginationPayload := payload.(*payloads.PaginationPayload)

	chunk, err := endpoint.service.List(user,*paginationPayload)
	return drivers.TypeToResponse(ctx, chunk, err)
}

func (endpoint storageEndpoint) store(ctx echo.Context) error {
	user:= drivers.GetUser(ctx)

	payload, err := drivers.RequestToPayload(ctx, new(payloads.StoreFilePayload))
	if err != nil {
		return err
	}
	storePayload := payload.(*payloads.StoreFilePayload)
	file,err := drivers.FileFromRequest(ctx, "File")
	if err != nil {
		return err
	}
	storedFile, err := endpoint.service.Store(user,file,*storePayload)

	return drivers.TypeToResponse(ctx, storedFile, err)
}

func (endpoint storageEndpoint) delete(ctx echo.Context) error {
	user:= drivers.GetUser(ctx)

	ID := ctx.Param("id")
	intID, err:=  strconv.Atoi(ID)
	if err != nil {
		return err
	}
	if intID <= 0 {
		return exceptions.InvalidInput
	}
	deletePayload := payloads.DeleteFilePayload{
		ID: uint(intID),
	}

	deletedFile, err := endpoint.service.Delete(user,deletePayload)

	return drivers.TypeToResponse(ctx, deletedFile, err)
}