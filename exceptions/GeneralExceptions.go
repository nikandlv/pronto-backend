package exceptions

import "nikan.dev/pronto/internals/exception"

var (
	NotAllowed = exception.Exception{
		Message: "access denied",
		Status:  exception.AccessDenied,
	}
	ServerError = exception.Exception{
		Message: "something went wrong",
		Status:  exception.ServerError,
	}
	InvalidInput = exception.Exception{
		Message: "invalid input",
		Status:  exception.InvalidInput,
	}
)
