package exceptions

import "nikan.dev/pronto/internals/exception"

var (
	CategoryAlreadyExists = exception.Exception{
		Message: "category already exists",
		Status:  exception.AlreadyExists,
	}
	CategoryParentNotFound = exception.Exception{
		Message: "parent does not exist",
		Status:  exception.NotFound,
	}
)
