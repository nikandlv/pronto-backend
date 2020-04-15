package exceptions

import "nikan.dev/pronto/internals/exception"

var (
	UserNotFound = exception.Exception{
		Message: "user not found",
		Status:  exception.NotFound,
	}
	UserAlreadyExists = exception.Exception{
		Message: "user already exists",
		Status:  exception.AlreadyExists,
	}

	InvalidCredentials = exception.Exception{
		Message: "invalid credentials",
		Status:  exception.InvalidInput,
	}
)
