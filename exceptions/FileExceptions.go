package exceptions

import "nikan.dev/pronto/internals/exception"

var (
	FileNotFound = exception.Exception{
		Message: "file not found",
		Status:  exception.NotFound,
	}
)
