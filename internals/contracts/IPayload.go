package internalContracts

type IPayload interface {
	Validation(validator IValidator) []error
}