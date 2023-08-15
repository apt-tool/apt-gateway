package config

import (
	"github.com/apt-tool/apt-gateway/internal/config/ftp"
	"github.com/apt-tool/apt-gateway/internal/config/http"
	"github.com/apt-tool/apt-gateway/internal/config/migration"
	"github.com/apt-tool/apt-gateway/internal/storage/redis"
	"github.com/apt-tool/apt-gateway/internal/storage/sql"
	"github.com/apt-tool/apt-gateway/internal/utils/jwt"
)

func Default() Config {
	return Config{
		HTTP: http.Config{
			Port: 8080,
			Core: "",
		},
		JWT: jwt.Config{
			PrivateKey: "private",
			ExpireTime: 60,
		},
		Redis: redis.Config{
			Host: "localhost:6379",
			Pass: "",
		},
		MySQL: sql.Config{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Pass:     "",
			Database: "automated-pen-testing",
			Migrate:  false,
		},
		Migrate: migration.Config{
			Enable: false,
		},
		FTP: ftp.Config{
			Host:   "",
			Secret: "",
			Access: "",
		},
	}
}
