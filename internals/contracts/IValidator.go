package internalContracts


type IValidator interface {
	String(key string, val interface{}) error
	Email(key string, val interface{}) error
	Boolean(key string, val interface{}) error
	Number(key string, val interface{}) error
	Validate(payload IPayload) error
}
