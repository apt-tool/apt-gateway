package controller

import (
	"github.com/apt-tool/apt-gateway/internal/config"
	"github.com/apt-tool/apt-gateway/internal/http/controller/handler"
	"github.com/apt-tool/apt-gateway/internal/storage/redis"
	"github.com/apt-tool/apt-gateway/internal/utils/jwt"
	"github.com/apt-tool/apt-gateway/pkg/client"

	"github.com/apt-tool/apt-core/pkg/models"
)

type Controller struct {
	Config           config.Config
	JWTAuthenticator jwt.Authenticator
	Models           *models.Interface
	RedisConnector   redis.Connector
	ErrHandler       handler.ErrorHandler
	Client           client.HTTPClient
}
