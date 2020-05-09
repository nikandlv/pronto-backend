package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/exception"
)

type UserIDOnlyPayload struct {
	ID uint
}

func (payload UserIDOnlyPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ID(payload.ID); err != nil {
		return err.(exception.Exception).WithPrefix("ID: ")
	}
	return nil
}


type UserEmailOnlyPayload struct {
	Email string
}

func (payload UserEmailOnlyPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Email(payload.Email); err != nil {
		return err.(exception.Exception).WithPrefix("Email: ")
	}
	return nil
}

type UserCredentialsPayload struct {
	Email string
	Password string
}

func (payload UserCredentialsPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Email(payload.Email); err != nil {
		return err.(exception.Exception).WithPrefix("Email: ")
	}
	if err := validator.Text(payload.Password); err != nil {
		return err.(exception.Exception).WithPrefix("Password: ")
	}
	return nil
}


type UserRegisterPayload struct {
	FirstName string
	LastName string
	Password string
	Email string
}

func (payload UserRegisterPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(payload.FirstName); err != nil {
		return err.(exception.Exception).WithPrefix("FirstName: ")
	}
	if err := validator.ShortText(payload.LastName); err != nil {
		return err.(exception.Exception).WithPrefix("LastName: ")
	}
	if err := validator.Email(payload.Email); err != nil {
		return err.(exception.Exception).WithPrefix("Email: ")
	}
	if err := validator.Text(payload.Password); err != nil {
		return err.(exception.Exception).WithPrefix("Email: ")
	}
	return nil
}


type JWTPayload struct {
	AccessToken string
	RefreshToken string
	Expire int64
}

func (payload JWTPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Text(payload.AccessToken); err != nil {
		return err.(exception.Exception).WithPrefix("AccessToken: ")
	}
	if err := validator.Text(payload.RefreshToken); err != nil {
		return err.(exception.Exception).WithPrefix("RefreshToken: ")
	}
	if err := validator.Timestamp(payload.Expire); err != nil {
		return err.(exception.Exception).WithPrefix("Expire: ")
	}
	return nil
}

type JWTRefreshPayload struct {
	RefreshToken string
}

func (payload JWTRefreshPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Text(payload.RefreshToken); err != nil {
		return err.(exception.Exception).WithPrefix("RefreshToken: ")
	}
	return nil
}