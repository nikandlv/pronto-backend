package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type ApplicationInfoPayload struct {
	Version string
}

func (i ApplicationInfoPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Version).Name("Version").Require().String(),
	};
}

type ApplicationPingPayload struct {
	Ping string
}

func (i ApplicationPingPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Ping).Name("Ping").Require().String(),
	};
}

type MessagePayload struct {
	Message string
}

func (i MessagePayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Message).Name("Message").Require().String(),
	};
}