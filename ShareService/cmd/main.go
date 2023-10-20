package main

import (
	"LinkService/config"
	"LinkService/internal/app"
	"LinkService/internal/cache"
	"LinkService/internal/files"
	"LinkService/internal/logger"
)

func main() {
	l := logger.NewLogger("[CMD]")
	al := logger.NewLogger("[APP]")

	cfg, err := config.InitConfig()
	if err != nil {
		l.Error(err)
	}

	files.InitFiles()

	cache, err := cache.NewCache(cfg.Cache["addr"], cfg.Cache["pass"])
	if err != nil {
		l.Error(err)
	}

	app, err := app.New(cfg, cache, al)
	if err != nil {
		l.Error(err)
	}

	app.Serve()
}
