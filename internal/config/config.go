package config

import (
	"encoding/json"
	"log"

	"github.com/ptaas-tool/gateway/internal/sql"
	"github.com/ptaas-tool/gateway/internal/utils/jwt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tidwall/pretty"
)

type (
	Config struct {
		HTTP  HTTPConfig `koanf:"http"`
		FTP   FTPConfig  `koanf:"ftp"`
		JWT   jwt.Config `koanf:"jwt"`
		MySQL sql.Config `koanf:"mysql"`
	}

	HTTPConfig struct {
		Port       int    `koanf:"port"`
		Core       string `koanf:"core"`
		CoreSecret string `koanf:"core_secret"`
		DevMode    bool   `koanf:"dev_mode"`
	}

	FTPConfig struct {
		Host   string `koanf:"host"`
		Secret string `koanf:"secret"`
		Access string `koanf:"access"`
	}
)

func Load(path string) Config {
	var instance Config

	k := koanf.New(".")

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	indent, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		log.Fatalf("error marshaling config to json: %s", err)
	}

	indent = pretty.Color(indent, nil)
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return instance
}
