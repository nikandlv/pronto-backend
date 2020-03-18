package http

import (
	"github.com/labstack/echo/v4"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/drivers"
)

type applicationEndpoint struct {
	service contracts.IApplicationService
}

func (endpoint applicationEndpoint) Boot(transport interface{}) {
	t := transport.(*echo.Group)
	group := t.Group("/application")
	group.GET("/info", endpoint.info)
	group.GET("/ping", endpoint.ping)
}

func NewApplicationEndpoint(service contracts.IApplicationService) applicationEndpoint {
	return applicationEndpoint{service }
}

func (endpoint applicationEndpoint) info(ctx echo.Context) error {
	payload, err := endpoint.service.Info()
	return drivers.PayloadToResponse(ctx, payload, err)
}

func (endpoint applicationEndpoint) ping(ctx echo.Context) error {
	payload, err := endpoint.service.Ping()
	return drivers.PayloadToResponse(ctx, payload, err)
}