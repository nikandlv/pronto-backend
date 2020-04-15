package drivers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type echoGateway struct{}

func NewEchoGateway() echoGateway {
	return echoGateway{}
}

func RequestToPayload(ctx echo.Context, payload internalContracts.IPayload) (internalContracts.IPayload,error) {
	if err := ctx.Bind(payload); err != nil {
		return payload, err
	}
	return payload, nil
}

func PayloadToResponse(ctx echo.Context, payload internalContracts.IPayload, err error) error {
	if err != nil {
		return err
	}
	return ctx.JSON(200,payload)
}
func TypeToResponse(ctx echo.Context, payload interface{}, err error) error {
	if err != nil {
		return err
	}
	return ctx.JSON(200,payload)
}

func (gateway echoGateway) Boot(config internalContracts.IConfiguration, endpoints ...internalContracts.IEndpoint) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.Static("/static"))
	e.HTTPErrorHandler = EchoExceptionDriver

	group := e.Group("/api")
	for index := range endpoints {
		endpoints[index].Boot(group)
	}
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
		host, err := config.Get("Host")
	if err != nil {
		panic(err)
	}
	e.Logger.Fatal(e.Start(host.(string)))

}

func ProtectedGroup(group *echo.Group) (*echo.Group) {
	panic("implement me")
}
