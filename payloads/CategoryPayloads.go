package payloads

import internalContracts "nikan.dev/pronto/internals/contracts"

type CategoryCreatePayload struct {
	Title string
	Slug string
	ParentID uint
}

func (i CategoryCreatePayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Title).Name("Title").Require().String(),
		validator.Validatable().Field(i.Slug).Name("Slug").Require().String(),
		validator.Validatable().Field(i.ParentID).Name("ParentID").Number(),
	};
}

