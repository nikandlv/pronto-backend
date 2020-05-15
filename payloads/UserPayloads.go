package payloads

import (
	internalContracts "nikan.dev/pronto/internals/contracts"
)

type UserIDOnlyPayload struct {
	ID uint
}

func (payload UserIDOnlyPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ID(payload.ID, "ID"); err != nil {
		return err
	}
	return nil
}


type UserEmailOnlyPayload struct {
	Email string
}

func (payload UserEmailOnlyPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Email(payload.Email, "Email"); err != nil {
		return err
	}
	return nil
}

type UserCredentialsPayload struct {
	Email string
	Password string
}

func (payload UserCredentialsPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Email(payload.Email, "Email"); err != nil {
		return err
	}
	if err := validator.Text(payload.Password, "Password"); err != nil {
		return err
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
	if err := validator.ShortText(payload.FirstName, "FirstName"); err != nil {
		return err
	}
	if err := validator.ShortText(payload.LastName, "LastName"); err != nil {
		return err
	}
	if err := validator.Email(payload.Email, "Email"); err != nil {
		return err
	}
	if err := validator.Text(payload.Password, "Password"); err != nil {
		return err
	}
	return nil
}


type JWTPayload struct {
	AccessToken string
	RefreshToken string
	Expire int64
}

func (payload JWTPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Text(payload.AccessToken, "AccessToken"); err != nil {
		return err
	}
	if err := validator.Text(payload.RefreshToken, "RefreshToken"); err != nil {
		return err
	}
	if err := validator.Timestamp(payload.Expire, "Expire"); err != nil {
		return err
	}
	return nil
}

type JWTRefreshPayload struct {
	RefreshToken string
}

func (payload JWTRefreshPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.Text(payload.RefreshToken, "RefreshToken"); err != nil {
		return err
	}
	return nil
}

type UserUpdatePayload struct {
	FirstName string
	LastName string
	Email string
}

func (payload UserUpdatePayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(payload.FirstName, "FirstName"); err != nil {
		return err
	}
	if err := validator.ShortText(payload.LastName, "LastName"); err != nil {
		return err
	}
	if err := validator.Email(payload.Email, "Email"); err != nil {
		return err
	}
	return nil
}

type UserUpdatePredefinedAvatarPayload struct {
	Avatar string
}

func (payload UserUpdatePredefinedAvatarPayload) Validate(validator internalContracts.IValidator) error {
	if err := validator.ShortText(payload.Avatar, "Avatar"); err != nil {
		return err
	}
	return nil
}