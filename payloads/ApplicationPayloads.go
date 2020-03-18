package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type ApplicationInfoPayload struct {
	Version string
}

func (i ApplicationInfoPayload) Validation(validator internalContracts.IValidator) []error {
	return []error {
		validator.String("Version", i.Version),
	};
}

type ApplicationPingPayload struct {
	Ping string
}

func (i ApplicationPingPayload) Validation(validator internalContracts.IValidator) []error {
	return []error {
		validator.String("Ping", i.Ping),
	};
}