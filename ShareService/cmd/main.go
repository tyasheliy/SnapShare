package main

import (
	"LinkService/config"
	"LinkService/internal/app"
	"LinkService/internal/cache"
	"LinkService/internal/controllers"
	"os"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	os.Mkdir("../tmp", 0755)

	redis, err := cache.NewRedisCache(cfg.CacheConfig.Address, cfg.CacheConfig.Password)
	if err != nil {
		panic(err)
	}

	links := controllers.NewLinkController(redis)
	entries := controllers.NewEntryController(redis)

	app, err := app.New(cfg, links, entries)
	if err != nil {
		panic(err)
	}

	app.Serve()
}
