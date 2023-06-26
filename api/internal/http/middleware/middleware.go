package middleware

import (
	"github.com/automated-pen-testing/api/internal/utils/jwt"
)

type Middleware struct {
	JWTAuthenticator jwt.Authenticator
}
