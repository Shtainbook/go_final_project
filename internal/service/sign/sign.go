package sign

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const (
	UnAuthorized = "authentication required"
)

type Password struct {
	Password string `json:"password"`
}

type SignAction interface {
	jwtToken() (string, error)
	Auth(token string) error
	SignIn(pass Password) (string, error)
}

type SignService struct {
	initialPassHash string
	secretKey       []byte
}

var Service *SignService

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	h := hex.EncodeToString(sha[:])
	return h
}

func InitSignService(initialPass string, secretKey []byte) *SignService {
	return &SignService{
		initialPassHash: hash(initialPass),
		secretKey:       secretKey,
	}
}

func (service *SignService) jwtToken() (string, error) {
	claims := jwt.MapClaims{
		"pass": service.initialPassHash,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(service.secretKey)
}

func (service *SignService) Auth(token string) error {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return service.secretKey, nil
	})
	if err != nil {
		return err
	}

	if !jwtToken.Valid {
		return errors.New(UnAuthorized)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New(UnAuthorized)
	}

	passHash, ok := claims["pass"].(string)
	if !ok {
		return errors.New(UnAuthorized)
	}
	if passHash != service.initialPassHash {
		return errors.New(UnAuthorized)
	}
	return nil
}

func (service *SignService) SignIn(pass Password) (string, error) {
	if service.initialPassHash == hash(pass.Password) {
		return service.jwtToken()
	}
	return "", errors.New(UnAuthorized)
}
