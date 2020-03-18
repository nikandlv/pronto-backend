package internalContracts

type IGateway interface {
	Boot(config IConfiguration, endpoints ...IEndpoint)
}
