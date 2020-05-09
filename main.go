package main

import (
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/endpoints/http"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/repositories/gorm"
	"nikan.dev/pronto/repositories/storage"
	"nikan.dev/pronto/services"
)


func main() {

	gateway := drivers.NewEchoGateway()
	config := drivers.NewViperConfiguration()
	validator := drivers.NewOzzoValidator()
	pool := drivers.NewGormDriver().Boot(config,entities.User{}, entities.Category{}, entities.Post{}, entities.Setting{}, entities.File{})

	deps := dependencies.CommonDependencies{Configuration: config, Validator: validator}

	storageDeps := dependencies.StorageDependencies{
		storage.NewLocalFileStorage(deps),
		gorm.NewFileRepository(deps),
	}

	settingsRepository := gorm.NewSettingsRepository(deps,pool)

	applicationRepository := gorm.NewApplicationRepository(deps, pool)
	applicationService := services.NewApplicationService(deps, applicationRepository, settingsRepository)
	applicationEndpoint := http.NewApplicationEndpoint(deps,applicationService)

	userRepository := gorm.NewUserRepository(deps, pool)
	userService := services.NewUserService(deps,storageDeps,userRepository)
	userEndpoint := http.NewUserEndpoint(deps,userService)

	categoryRepository := gorm.NewCategoryRepository(deps, pool)
	categoryService := services.NewCategoryService(deps,categoryRepository)
	categoryEndpoint := http.NewCategoryEndpoint(deps,categoryService)

	postRepository := gorm.NewPostRepository(deps, pool)
	postService := services.NewPostService(deps, postRepository)
	postEndpoint := http.NewPostEndpoint(deps,postService)

	gateway.Boot(config, applicationEndpoint, userEndpoint, categoryEndpoint, postEndpoint)

}
