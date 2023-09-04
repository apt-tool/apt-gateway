package config

import (
	"github.com/apt-tool/apt-gateway/internal/sql"
	"github.com/apt-tool/apt-gateway/internal/utils/jwt"
)

func Default() Config {
	return Config{
		HTTP: HTTPConfig{
			Port:       8080,
			Core:       "",
			CoreSecret: "",
		},
		FTP: FTPConfig{
			Host:   "",
			Secret: "",
			Access: "",
		},
		JWT: jwt.Config{
			PrivateKey: "private",
			ExpireTime: 60,
		},
		MySQL: sql.Config{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Pass:     "",
			Database: "automated-pen-testing",
			Migrate:  false,
		},
	}
}
