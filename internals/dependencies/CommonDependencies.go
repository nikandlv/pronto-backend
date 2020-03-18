package dependencies

import (
	contracts2 "nikan.dev/pronto/internals/contracts"
)

type CommonDependencies struct {
	Configuration contracts2.IConfiguration
	Validator     contracts2.IValidator
}