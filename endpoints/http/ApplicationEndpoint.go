package http

import (
	"github.com/labstack/echo/v4"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type applicationEndpoint struct {
	deps dependencies.CommonDependencies
	service contracts.IApplicationService
}

func (endpoint applicationEndpoint) Boot(transport interface{}) {
	t := transport.(*echo.Group)
	group := t.Group("/application")
	group.GET("/info", endpoint.info).Name = "Application.Info"
	group.GET("/config", endpoint.config).Name = "Application.Config"
	group.GET("/ping", endpoint.ping).Name = "Application.Ping"
	group.GET("/echo", endpoint.echo).Name = "Application.Echo"
}

func NewApplicationEndpoint(deps dependencies.CommonDependencies, service contracts.IApplicationService) applicationEndpoint {
	return applicationEndpoint{deps,service }
}

func (endpoint applicationEndpoint) info(ctx echo.Context) error {
	payload, err := endpoint.service.Info()
	return drivers.PayloadToResponse(ctx, payload, err)
}

func (endpoint applicationEndpoint) config(ctx echo.Context) error {
	payload, err := endpoint.service.Config()
	return drivers.PayloadToResponse(ctx, payload, err)
}

func (endpoint applicationEndpoint) ping(ctx echo.Context) error {
	payload, err := endpoint.service.Ping()
	return drivers.PayloadToResponse(ctx, payload, err)
}

func (endpoint applicationEndpoint) echo(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.MessagePayload))
	if err != nil {
		return err
	}
	validationErr := payload.Validate(endpoint.deps.Validator)
	return drivers.PayloadToResponse(ctx, payload, validationErr)
}