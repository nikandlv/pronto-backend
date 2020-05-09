package dependencies

import (
	"nikan.dev/pronto/contracts"
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type CommonDependencies struct {
	Configuration internalContracts.IConfiguration
	Validator     internalContracts.IValidator
}

type StorageDependencies struct {
	Storage contracts.IFileStorage
	Repository contracts.IFileRepository
}