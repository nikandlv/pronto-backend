package main

import (
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/endpoints/http"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/repositories/mysql"
	"nikan.dev/pronto/services"
)


func main() {

	gateway := drivers.NewEchoGateway()
	config := drivers.NewViperConfiguration()
	validator := drivers.NewOzzoValidator()
	pool := drivers.NewGormDriver()
	pool.Boot(config,entities.User{})

	deps := dependencies.CommonDependencies{config,validator,}

	applicationRepository := mysql.NewApplicationRepository(deps, pool)
	applicationService := services.NewApplicationService(deps, applicationRepository)
	applicationEndpoint := http.NewApplicationEndpoint(deps,applicationService)

	gateway.Boot(config, applicationEndpoint)

}
