package internalContracts

type IPayload interface {
	Validate(validator IValidator) error
}