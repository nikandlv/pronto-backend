package drivers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"nikan.dev/pronto/exceptions"
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/entity"
	"strings"
)

type echoGateway struct{
}

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
	e.Use(middleware.Static("storage"))
	e.Static("/storage", "storage")
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

func FileFromRequest(ctx echo.Context, name string) (entity.FileEntity, error) {
	file, err := ctx.FormFile(name)
	if err != nil {
		return entity.FileEntity{}, err
	}
	src, err := file.Open()
	if err != nil {
		return entity.FileEntity{},err
	}
	//// Destination
	//target := fmt.Sprintf("%v/%v-%v", directory, uuid.New().String(),file.Filename)
	//dst, err := os.Create(target)
	//if err != nil {
	//	return &multipart.FileHeader{}, "",err
	//}
	//defer dst.Close()
	//
	//// Store
	//if _, err = io.Store(dst, src); err != nil {
	//	return &multipart.FileHeader{}, "",err
	//}
	return entity.FileEntity{
		Reader: src,
		Name:   file.Filename,
		Size:   file.Size,
		Mime:   fmt.Sprintf("%v", file.Header.Get("Content-Type")),
	}, nil
}

func MimeFileFromRequest(ctx echo.Context, name string, mime string) (entity.FileEntity, error) {
	file, err := ctx.FormFile(name)
	if err != nil {
		return entity.FileEntity{}, err
	}
	src, err := file.Open()
	if err != nil {
		return entity.FileEntity{},err
	}
	//// Destination
	//target := fmt.Sprintf("%v/%v-%v", directory, uuid.New().String(),file.Filename)
	//dst, err := os.Create(target)
	//if err != nil {
	//	return &multipart.FileHeader{}, "",err
	//}
	//defer dst.Close()
	//
	//// Store
	//if _, err = io.Store(dst, src); err != nil {
	//	return &multipart.FileHeader{}, "",err
	//}

	if !strings.Contains(fmt.Sprintf("%v", file.Header), mime) {
		return entity.FileEntity{}, exceptions.InvalidInput.WithMessage(fmt.Sprintf("%v should be of type %v",name, mime))
	}


	return entity.FileEntity{
		Reader: src,
		Name:   file.Filename,
		Size:   file.Size,
		Mime:   fmt.Sprintf("%v", file.Header.Get("Content-Type")),
	}, nil

}