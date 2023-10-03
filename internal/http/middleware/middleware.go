package middleware

import (
	"github.com/ptaas-tool/gateway/internal/http/controller/handler"
	"github.com/ptaas-tool/gateway/internal/utils/jwt"
)

type Middleware struct {
	JWTAuthenticator jwt.Authenticator
	ErrHandler       handler.ErrorHandler
}
