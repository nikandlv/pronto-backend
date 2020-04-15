package payloads

import internalContracts "nikan.dev/pronto/internals/contracts"

type UserIDOnlyPayload struct {
	ID uint
}

func (i UserIDOnlyPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.ID).Name("ID").Require().Number(),
	}
}


type UserEmailOnlyPayload struct {
	Email string
}

func (i UserEmailOnlyPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Email).Name("Email").Require().Email(),
	};
}

type UserCredentialsPayload struct {
	Email string
	Password string
}

func (i UserCredentialsPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Email).Name("Email").Require().Email(),
		validator.Validatable().Field(i.Password).Name("Password").Require().String(),
	};
}


type UserRegisterPayload struct {
	FirstName string
	LastName string
	Password string
	Email string
}

func (i UserRegisterPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.FirstName).Name("FirstName").Require().String(),
		validator.Validatable().Field(i.LastName).Name("LastName").Require().String(),
		validator.Validatable().Field(i.Password).Name("Password").Require().String(),
		validator.Validatable().Field(i.Email).Name("Email").Require().Email(),
	}
}

type JWTPayload struct {
	AccessToken string
	RefreshToken string
	Expire int64
}

func (i JWTPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.AccessToken).Name("AccessToken").Require().String(),
		validator.Validatable().Field(i.RefreshToken).Name("RefreshToken").Require().String(),
		validator.Validatable().Field(i.Expire).Name("Expire").Require().Number(),
	};
}

type JWTRefreshPayload struct {
	RefreshToken string
}

func (i JWTRefreshPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.RefreshToken).Name("RefreshToken").Require().String(),
	};
}
