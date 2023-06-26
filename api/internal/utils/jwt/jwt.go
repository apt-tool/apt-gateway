package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type authenticator struct {
	key    string
	expire int
}

func (a *authenticator) GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	expireTime := time.Now().Add(time.Duration(a.expire) * time.Minute)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expireTime.Unix()
	claims["email"] = email

	tokenString, err := token.SignedString([]byte(a.key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authenticator) ParseToken(token string) (string, error) {
	tokenString, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrSigningMethod
		}

		return []byte(a.key), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := tokenString.Claims.(jwt.MapClaims); ok && tokenString.Valid {
		email := claims["email"].(string)

		return email, nil
	}

	return "", ErrInvalidToken
}
