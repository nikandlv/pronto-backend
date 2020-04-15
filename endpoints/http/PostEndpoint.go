package http

import (
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/internals/dependencies"
)

type postEndpoint struct {
	deps dependencies.CommonDependencies
	service contracts.IPostService
}

func (endpoint postEndpoint) Boot(transport interface{}) {
	//t := transport.(*echo.Group)
	//group := t.Group("/application")
	//group.GET("/info", endpoint.info)
	//group.GET("/ping", endpoint.ping)
	//group.GET("/echo", endpoint.echo)
}

func NewPostEndpoint(deps dependencies.CommonDependencies, service contracts.IPostService) postEndpoint {
	return postEndpoint{deps,service }
}
