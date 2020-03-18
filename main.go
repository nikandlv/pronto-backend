package main

import (
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/endpoints/http"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/repositories/mysql"
	"nikan.dev/pronto/services"
)


func main() {

	gateway := drivers.NewEchoGateway()
	config := drivers.NewViperConfiguration()
	pool := drivers.NewGormDriver()
	validator := drivers.NewOzzoValidator()

	deps := dependencies.CommonDependencies{config,validator,}

	applicationRepository := mysql.NewApplicationRepository(deps, pool)
	applicationService := services.NewApplicationService(deps, applicationRepository)
	applicationEndpoint := http.NewApplicationEndpoint(applicationService)

	gateway.Boot(config, applicationEndpoint)

}
