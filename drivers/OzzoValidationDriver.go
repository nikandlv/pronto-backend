package drivers

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"nikan.dev/pronto/exceptions"
	internalContracts "nikan.dev/pronto/internals/contracts"
)
import Is "github.com/go-ozzo/ozzo-validation/is"

type ozzoValidator struct {
}

func (o ozzoValidator) Validatable() internalContracts.IValidatable {
	return newValidatable()
}

func NewOzzoValidator() ozzoValidator {
	return ozzoValidator{}
}



type validatable struct {
	Key string
	Value interface{}
	Type  internalContracts.ValidationType
	Required bool
}

func (v validatable) Field(value interface{}) internalContracts.IValidatable {
	v.Value = value
	return v
}

func (v validatable) Name(key string) internalContracts.IValidatable {
	v.Key = key
	return v
}

func newValidatable() validatable {
	return validatable{}
}

func (v validatable) String() internalContracts.IValidatable {
	v.Type = internalContracts.StringValidation
	return v
}

func (v validatable) Email() internalContracts.IValidatable {
	v.Type = internalContracts.EmailValidation
	return v
}

func (v validatable) Boolean() internalContracts.IValidatable {
	v.Type = internalContracts.BooleanValidation
	return v
}

func (v validatable) Number() internalContracts.IValidatable {
	v.Type = internalContracts.NumberValidation
	return v
}

func (v validatable) Require() internalContracts.IValidatable {
	v.Required = true
	return v
}


func (o ozzoValidator) String(rule validatable) error {
	if val, ok := rule.Value.(string); ok == true {
		if !rule.Required && val == "" {
			return nil
		}
	}
	return generateResult(validation.Validate(rule.Value, validation.Required), rule.Key)
}

func (o ozzoValidator) Email(rule validatable) error {
	if val, ok := rule.Value.(string); ok == true {
		if !rule.Required && val == "" {
			return nil
		}
	}
	return generateResult(validation.Validate(rule.Value, validation.Required, Is.Email), rule.Key)
}

func (o ozzoValidator) Boolean(rule validatable) error {
	return generateResult(validation.Validate(
		rule.Value, validation.Required,
		Is.Int, validation.Length(0, 1)), rule.Key)
}

func (o ozzoValidator) Number(rule validatable) error {
	return generateResult(validation.Validate(fmt.Sprintf("%v", rule.Value),  Is.Int), rule.Key)
}

func generateResult(err error, key string) error {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%v %v", key, err))
}

func (o ozzoValidator) Validate(payload internalContracts.IPayload) error {
	items := payload.Validation(o)
	for i := range items {
		rule := items[i].(validatable)
		var err error
		switch rule.Type {
			case internalContracts.StringValidation:
				err = o.String(rule)
				break;
			case internalContracts.EmailValidation:
				err = o.Email(rule)
				break;
			case internalContracts.NumberValidation:
				err = o.Number(rule)
				break;
			case internalContracts.BooleanValidation:
				err = o.Boolean(rule)
				break;
		}
		if err != nil {
			return exceptions.InvalidInput.WithMessage(err.Error())
		}
	}
	return nil
}
