package payloads

import (
	"fmt"
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/exception"
)

type SettingSetPayload struct {
	Name string
	Value string
}

func (payload SettingSetPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(payload.Name); err != nil {
		return err.(exception.Exception).WithPrefix("Name: ")
	}
	if err := validator.Text(payload.Value); err != nil {
		return err.(exception.Exception).WithPrefix("Value: ")
	}
	return nil
}


type SettingGetPayload struct {
	Name string

}

func (payload SettingGetPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(payload.Name); err != nil {
		return err.(exception.Exception).WithPrefix("Name: ")
	}
	return nil
}

type SettingSetBatchPayload struct {
	Settings []SettingSetPayload
}

func (payload SettingSetBatchPayload) Validate(validator internalContracts.IValidator) error {
	for index,item  := range payload.Settings {
		if err := item.Validate(validator); err != nil {
			return err.(exception.Exception).WithPrefix(fmt.Sprintf("Settings[%v]:", index))
		}
	}
	return nil
}
