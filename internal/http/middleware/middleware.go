package middleware

import (
	"github.com/apt-tool/apt-gateway/internal/http/controller/handler"
	"github.com/apt-tool/apt-gateway/internal/storage/redis"
	"github.com/apt-tool/apt-gateway/internal/utils/jwt"
	"github.com/apt-tool/apt-gateway/pkg/models"
)

type Middleware struct {
	JWTAuthenticator jwt.Authenticator
	Models           *models.Interface
	RedisConnector   redis.Connector
	ErrHandler       handler.ErrorHandler
}
