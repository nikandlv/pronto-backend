package http

import (
	"github.com/dgrijalva/jwt-go"
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
	protected := drivers.JwtGroup(endpoint.deps.Configuration,group)
	protected.GET("/info", endpoint.get).Name = "User.Info"
	protected.PATCH("/info", endpoint.update).Name = "User.Update"
	protected.POST("/avatar", endpoint.updateAvatar).Name = "User.Avatar"
	protected.PATCH("/avatar", endpoint.updatePredefinedAvatar).Name = "User.PredefinedAvatar"
	group.POST("/login", endpoint.login).Name = "User.Login"
	group.POST("/register", endpoint.register).Name = "User.Register"
	group.POST("/refresh", endpoint.refresh).Name = "User.Refresh"
}

func NewUserEndpoint(deps dependencies.CommonDependencies, service contracts.IUserService) userEndpoint {
	return userEndpoint{deps,service }
}

func (endpoint userEndpoint) get(ctx echo.Context) error {
	user:= drivers.GetUser(ctx)
	return drivers.TypeToResponse(ctx, user, nil)
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
	token, err := drivers.GenerateJWT(endpoint.deps.Configuration,user)
	return drivers.TypeToResponse(ctx, token, err)
}

func (endpoint userEndpoint) refresh(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.JWTRefreshPayload))
	if err != nil {
		return err
	}
	token, err := drivers.RefreshJWT(endpoint.deps,endpoint.service,*payload.(*payloads.JWTRefreshPayload))
	return drivers.TypeToResponse(ctx, token, err)
}

func (endpoint userEndpoint) update(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.UserUpdatePayload))
	if err != nil {
		return err
	}

	updatePayload := payload.(*payloads.UserUpdatePayload)
	user := ctx.Get("user").(*jwt.Token)
	claims := drivers.GetClaims(user.Claims)
	u, err := endpoint.service.Update(*claims.User,*updatePayload)
	return drivers.TypeToResponse(ctx, u, err)
}
func (endpoint userEndpoint) updateAvatar(ctx echo.Context) error {
	file, uploadError := drivers.MimeFileFromRequest(ctx, "Avatar", "image");
	if uploadError != nil {
		return uploadError
	}
	user := ctx.Get("user").(*jwt.Token)
	claims := drivers.GetClaims(user.Claims)
	u, err := endpoint.service.UpdateAvatar(*claims.User, file)
	return drivers.TypeToResponse(ctx, u, err)
}
func (endpoint userEndpoint) updatePredefinedAvatar(ctx echo.Context) error {
	payload, err := drivers.RequestToPayload(ctx, new(payloads.UserUpdatePredefinedAvatarPayload));
	if err != nil {
		return err
	}
	updatePayload := payload.(*payloads.UserUpdatePredefinedAvatarPayload)
	user := ctx.Get("user").(*jwt.Token)
	claims := drivers.GetClaims(user.Claims)
	u, err := endpoint.service.UpdatePredefinedAvatar(*claims.User,*updatePayload)
	return drivers.TypeToResponse(ctx, u, err)
}