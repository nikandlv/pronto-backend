package http

import (
	"github.com/labstack/echo/v4"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type userEndpoint struct {
	deps dependencies.CommonDependencies
	service contracts.IUserService
}

func (endpoint userEndpoint) Boot(transport interface{}) {
	t := transport.(*echo.Group)
	group := t.Group("/user")
	group.GET("/info", endpoint.get).Name = "User.Info"
	group.POST("/login", endpoint.login).Name = "User.Login"
	group.POST("/register", endpoint.register).Name = "User.Register"
	group.POST("/refresh", endpoint.refresh).Name = "User.Refresh"
}

func NewUserEndpoint(deps dependencies.CommonDependencies, service contracts.IUserService) userEndpoint {
	return userEndpoint{deps,service }
}

func (endpoint userEndpoint) get(ctx echo.Context) error {
	payload, err := endpoint.service.Get(payloads.UserIDOnlyPayload{ID:1})
	if err != nil {
		return err
	}
	return drivers.TypeToResponse(ctx, payload, err)
}

func (endpoint userEndpoint) login(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.UserCredentialsPayload))
	if err != nil {
		return err
	}
	user, err := endpoint.service.CheckCredentials(*payload.(*payloads.UserCredentialsPayload))
	if err != nil {
		return exceptions.InvalidCredentials
	}
	token, err := drivers.GenerateJWT(endpoint.deps.Configuration,user)
	return drivers.TypeToResponse(ctx, token, err)
}

func (endpoint userEndpoint) register(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.UserRegisterPayload))
	if err != nil {
		return err
	}
	user, err := endpoint.service.Register(*payload.(*payloads.UserRegisterPayload))
	if err != nil {
		return err
	}
	return drivers.TypeToResponse(ctx, user, err)
}

func (endpoint userEndpoint) refresh(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.JWTRefreshPayload))
	if err != nil {
		return err
	}
	token, err := drivers.RefreshJWT(endpoint.deps,endpoint.service,*payload.(*payloads.JWTRefreshPayload))
	return drivers.TypeToResponse(ctx, token, err)
}