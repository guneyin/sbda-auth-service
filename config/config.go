package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const defaultRpcPort = 3000

type Config struct {
	RpcPort                 int    `env:"RPC_PORT"`
	DiscoverySvcAddr        string `env:"DISCOVERY_SVC_ADDR"`
	GoogleOauthClientId     string `env:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleOauthClientSecret string `env:"GOOGLE_OAUTH_CLIENT_SECRET"`
}

func GetConfig() *Config {
	var config Config

	err := cleanenv.ReadConfig(".env", &config)
	if err != nil {
		return nil
	}

	if config.RpcPort == 0 {
		config.RpcPort = defaultRpcPort
	}

	return &config
}
