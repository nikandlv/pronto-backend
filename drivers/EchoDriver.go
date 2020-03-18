package drivers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type echoGateway struct{}

func NewEchoGateway() echoGateway {
	return echoGateway{}
}

func PayloadToResponse(ctx echo.Context, payload internalContracts.IPayload, err error) error {
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

	host, err := config.Get("Host")
	if err != nil {
		panic(err)
	}
	e.Logger.Fatal(e.Start(host.(string)))

}
