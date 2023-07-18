package controller

import (
	"github.com/automated-pen-testing/api/internal/storage/redis"
	"github.com/automated-pen-testing/api/internal/utils/jwt"
	"github.com/automated-pen-testing/api/pkg/models"
)

type Controller struct {
	JWTAuthenticator jwt.Authenticator
	Models           *models.Interface
	RedisConnector   redis.Connector
}
