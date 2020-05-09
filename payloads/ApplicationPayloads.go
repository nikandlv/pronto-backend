package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type ApplicationInfoPayload struct {
	Version string
}

func (a ApplicationInfoPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(a.Version, "Version"); err != nil {
		return err
	}
	return nil
}

type ApplicationPingPayload struct {
	Ping string
}

func (i ApplicationPingPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(i.Ping, "Ping"); err != nil {
		return err
	}
	return nil
}


type MessagePayload struct {
	Message string
}

func (m MessagePayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(m.Message, "Message"); err != nil {
		return err
	}
	return nil}
