package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/exception"
)

type CategoryCreatePayload struct {
	Title string
	Slug string
	ParentID uint
}

func (payload CategoryCreatePayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(payload.Title); err != nil {
		return err.(exception.Exception).WithPrefix("Title: ")
	}
	if err := validator.ShortText(payload.Slug); err != nil {
		return err.(exception.Exception).WithPrefix("Slug: ")
	}
	if err := validator.ID(payload.ParentID); err != nil {
		return err.(exception.Exception).WithPrefix("ParentID: ")
	}
	return nil
}
