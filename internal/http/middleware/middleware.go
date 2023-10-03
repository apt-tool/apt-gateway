package middleware

import (
	"github.com/ptaas-tool/gateway/internal/http/controller/handler"
	"github.com/ptaas-tool/gateway/internal/utils/jwt"

	"github.com/ptaas-tool/base-api/pkg/models"
)

type Middleware struct {
	JWTAuthenticator jwt.Authenticator
	Models           *models.Interface
	ErrHandler       handler.ErrorHandler
}
