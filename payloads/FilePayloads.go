package payloads

import internalContracts "nikan.dev/pronto/internals/contracts"

type StoreFilePayload struct {
	Scope string
}

func (payload StoreFilePayload) Validate(validator internalContracts.IValidator) error {
	return validator.Text(payload.Scope, "Scope")
}
type DeleteFilePayload struct {
	ID uint
}

func (payload DeleteFilePayload) Validate(validator internalContracts.IValidator) error {
	return validator.ID(payload.ID, "ID")
}