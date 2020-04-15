package internalContracts

type ValidationType int

const (
	StringValidation  ValidationType = iota + 1
	NumberValidation
	EmailValidation
	BooleanValidation
)

type IValidatable interface {
	Field(value interface{}) IValidatable
	Name(key string) IValidatable
	String() IValidatable
	Email() IValidatable
	Boolean() IValidatable
	Number() IValidatable
	Require() IValidatable
}

type IValidator interface {
	Validate(payload IPayload) error
	Validatable() IValidatable
}
