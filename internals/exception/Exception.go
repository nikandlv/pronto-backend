package exception

type Status int

const (
	NotFound Status = iota
	AccessDenied
	ServerError
	AlreadyExists
	InvalidInput
)

type Exception struct {
	Message string `json:"message"`
	Status  Status
}

func (e Exception) Error() string {
	return e.Message
}

func (e Exception) WithMessage(Message string) Exception {
	e.Message = Message
	return e
}

func (e Exception) WithStatus(Status Status) Exception {
	e.Status = Status
	return e
}
