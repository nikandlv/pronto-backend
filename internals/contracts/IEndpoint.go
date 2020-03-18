package internalContracts

type IEndpoint interface {
	Boot(transport interface{})
}
