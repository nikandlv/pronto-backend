package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/exception"
)

type ApplicationInfoPayload struct {
	Version string
}

func (a ApplicationInfoPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(a.Version); err != nil {
		return err.(exception.Exception).WithPrefix("Version: ")
	}
	return nil
}

type ApplicationPingPayload struct {
	Ping string
}

func (i ApplicationPingPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(i.Ping); err != nil {
		return err.(exception.Exception).WithPrefix("Ping: ")
	}
	return nil
}


type MessagePayload struct {
	Message string
}

func (m MessagePayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(m.Message); err != nil {
		return err.(exception.Exception).WithPrefix("Message: ")
	}
	return nil}
