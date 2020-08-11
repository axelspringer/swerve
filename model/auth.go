package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// CheckPassword checks pwd hash on db against entered plain pwd, returns nil if it is correct
func (c *Model) CheckPassword(username string, plainPwd string) error {
	pwd, err := c.DB.GetPwdHash(username)
	if err != nil {
		return err
	}

	byteHash := []byte(pwd)
	bytePlain := []byte(plainPwd)

	err = bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	// err = bcrypt.CompareHashAndPassword(bytePlain, byteHash)
	if err != nil {
		return errors.WithMessage(err, ErrPwdInvalid)
	}

	return nil
}

// CreateJWT creates a new JWT with the given username and secret
func (c *Model) CreateJWT(username string, secret string) (string, int64, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, errors.WithMessage(err, ErrJWTSign)
	}

	return tokenString, expirationTime.Unix(), nil
}

// CheckToken checks if the given token is valid
func (c *Model) CheckToken(token string, secret string) bool {
	claims := &claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return err == nil && tkn.Valid
}
