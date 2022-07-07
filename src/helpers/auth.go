package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	Id string `json:"_id"`
	// Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var expirationTime = time.Now().Add(time.Hour * 3)

func NewToken(id, email, role string) *claims {
	return &claims{
		Id:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
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
