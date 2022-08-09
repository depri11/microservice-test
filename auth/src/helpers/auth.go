package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func NewToken(id string) *claims {
	return &claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func CheckToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokens.Claims.(*claims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	return claims, nil
}
