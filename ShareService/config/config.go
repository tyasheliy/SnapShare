package config

import (
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Build string
	Port  string
}

type TokenConfig struct {
	SigningKey []byte
}

type CacheConfig struct {
	Address  string
	Password string
}

type Config struct {
	AppConfig
	TokenConfig
	CacheConfig
}

func InitConfig() (*Config, error) {
	vp := viper.New()

	vp.AddConfigPath("./config")
	vp.AddConfigPath("../config")
	vp.AddConfigPath(".")

	vp.SetConfigName("config")
	vp.SetConfigType("json")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	app := AppConfig{
		Build: vp.GetString("app.build"),
		Port:  vp.GetString("app.port"),
	}

	t := TokenConfig{
		SigningKey: []byte(os.Getenv("TOKEN_SECRET_KEY")),
	}

	c := CacheConfig{
		Address:  "host.docker.internal:4002",
		Password: os.Getenv("SHARE_REDIS_PASSWORD"),
	}

	return &Config{
		AppConfig:   app,
		TokenConfig: t,
		CacheConfig: c,
	}, nil
}
