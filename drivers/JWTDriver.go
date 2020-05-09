package drivers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/entities"
	internalContracts "nikan.dev/pronto/internals/contracts"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
	"time"
)

type jwtGrantClaims struct {
	user entities.User
	jwt.StandardClaims
}

type jwtRefreshClaims struct {
	ID uint
	jwt.StandardClaims
}

func GenerateJWT(config internalContracts.IConfiguration, user entities.User) (payloads.JWTPayload, error) {
	secret, err := config.Get("Secret")
	if err != nil {
		return payloads.JWTPayload{}, err
	}
	accessTokenExpire := time.Now().Add(time.Hour * 2).Unix()
	claims := &jwtGrantClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: accessTokenExpire,
		},
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenExpire := time.Now().Add(time.Hour * 72).Unix()
	refreshToken.Claims = &jwtRefreshClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: refreshTokenExpire,
		},
	}
	rt, err := refreshToken.SignedString([]byte(secret.(string)))
	if err != nil {
		return payloads.JWTPayload{}, err
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret.(string)))
	if err != nil {
		return payloads.JWTPayload{}, err
	}
	return  payloads.JWTPayload{
		AccessToken:  t,
		RefreshToken: rt,
		Expire:       accessTokenExpire,
	}, nil
}

func RefreshJWT(deps dependencies.CommonDependencies, userService contracts.IUserService, payload payloads.JWTRefreshPayload) (payloads.JWTPayload, error) {
	if err := payload.Validate(deps.Validator); err != nil {
		return payloads.JWTPayload{},err
	}
	secret, err := deps.Configuration.Get("Secret")
	if err != nil {
		return payloads.JWTPayload{}, err
	}
	token, err := jwt.ParseWithClaims(payload.RefreshToken, &jwtRefreshClaims{} ,func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret.(string)), nil
	})
	if err != nil {
		return payloads.JWTPayload{}, err
	}
	claims, ok := token.Claims.(*jwtRefreshClaims);
	if ok && token.Valid {
		user, err := userService.Get(payloads.UserIDOnlyPayload{claims.ID})
		if err != nil {
			return payloads.JWTPayload{}, err
		}
		return GenerateJWT(deps.Configuration, user )
	}

	return payloads.JWTPayload{}, errors.New("cant verify the refresh token")

}