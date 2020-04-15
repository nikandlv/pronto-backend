package hash

import (
	"golang.org/x/crypto/scrypt"
	"nikan.dev/pronto/internals/contracts"
)

func Generate(config internalContracts.IConfiguration, val string) ([]byte, error) {
	configSalt, err := config.Get("Salt")
	if err != nil {
		return nil,err
	}
	salt := []byte(configSalt.(string))

	dk, err := scrypt.Key([]byte(val), salt, 1<<15, 8, 1, 32)
	if err != nil {
		return nil,err
	}
	return dk, nil
}
