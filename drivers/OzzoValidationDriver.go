package drivers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"nikan.dev/pronto/exceptions"
)

type ozzoValidator struct {
}

func (o ozzoValidator) ShortText(s string) error {
	if len(s) == 0 {
		return exceptions.InvalidInput.WithMessage("cannot be empty")
	}
	if len(s) > 255 {
		return exceptions.InvalidInput.WithMessage("cannot be longer then 255 characters")
	}
	return nil
}

func (o ozzoValidator) Text(s string) error {
	if len(s) == 0 {
		return exceptions.InvalidInput.WithMessage("cannot be empty")
	}
	return nil
}

func (o ozzoValidator) ID(u uint) error {
	if u == 0 {
		return exceptions.InvalidInput.WithMessage("cannot be 0")
	}
	return nil
}

func (o ozzoValidator) Timestamp(i int64) error {
	err := validation.Validate(i,
		validation.Required,
		is.Int)
	if err != nil {
		return exceptions.InvalidInput.WithMessage(err.Error())
	}
	return nil
}

func (o ozzoValidator) Email(s string) error {
	err := validation.Validate(s,
		validation.Required,
		is.Email)
	if err != nil {
		return exceptions.InvalidInput.WithMessage(err.Error())
	}
	return nil
}

func NewOzzoValidator() ozzoValidator {
	return ozzoValidator{}
}
