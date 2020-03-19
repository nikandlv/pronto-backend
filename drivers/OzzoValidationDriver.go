package drivers

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"nikan.dev/pronto/exceptions"
	contracts2 "nikan.dev/pronto/internals/contracts"
)
import Is "github.com/go-ozzo/ozzo-validation/is"

type ozzoValidator struct {
}

func NewOzzoValidator() ozzoValidator {
	return ozzoValidator{}
}

func (o ozzoValidator) String(key string, val interface{}) error {
	return generateResult(validation.Validate(val, validation.Required), key)
}

func (o ozzoValidator) Email(key string, val interface{}) error {
	return generateResult(validation.Validate(val, validation.Required, Is.Email), key)
}

func (o ozzoValidator) Boolean(key string, val interface{}) error {
	return generateResult(validation.Validate(
		val, validation.Required,
		Is.Int, validation.Length(0, 1)), key)
}

func (o ozzoValidator) Number(key string, val interface{}) error {
	return generateResult(validation.Validate(val, validation.Required, Is.Int), key)
}

func generateResult(err error, key string) error {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%v %v", key, err))
}

func (o ozzoValidator) Validate(payload contracts2.IPayload) error {
	items := payload.Validation(o)
	for err := range items {
		if items[err] != nil {
			return exceptions.InvalidInput.WithMessage(items[err].Error())
		}
	}
	return nil
}
