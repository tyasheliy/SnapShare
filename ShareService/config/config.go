package config

import "github.com/spf13/viper"

type AppConfig struct {
	Build string
	Port  string
}

type TokenConfig struct {
	SigningKey []byte
	Issuer     string
	Audience   string
}

type Config struct {
	AppConfig
	TokenConfig
	Cache       map[string]string
	ServiceURLs map[string]string
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
		SigningKey: []byte(vp.GetString("authToken.secretKey")),
		Issuer:     vp.GetString("authToken.issuer"),
		Audience:   vp.GetString("authToken.audience"),
	}

	return &Config{
		AppConfig:   app,
		Cache:       vp.GetStringMapString("conns.cache"),
		ServiceURLs: vp.GetStringMapString("serviceUrls"),
		TokenConfig: t,
	}, nil
}
