package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type CategoryCreatePayload struct {
	Title string
	Slug string
	ParentID uint
}

func (payload CategoryCreatePayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(payload.Title, "Title"); err != nil {
		return err
	}
	if err := validator.ShortText(payload.Slug, "Slug"); err != nil {
		return err
	}
	return nil
}
