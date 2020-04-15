package payloads

import internalContracts "nikan.dev/pronto/internals/contracts"

type SettingSetPayload struct {
	Name string
	Value string
}

func (i SettingSetPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Name).Name("Name").Require().String(),
		validator.Validatable().Field(i.Name).Name("Value").Require().String(),
	}
}

type SettingGetPayload struct {
	Name string
}

func (i SettingGetPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Name).Name("Name").Require().String(),
	}
}

type SettingSetBatchPayload struct {
	Setting []SettingSetPayload
}

func (i SettingSetBatchPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Setting).Name("Setting").Require(),
	}
}