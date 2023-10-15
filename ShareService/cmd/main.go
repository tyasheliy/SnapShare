package main

import (
	"LinkService/config"
	"LinkService/internal/app"
	"LinkService/internal/cache"
	"LinkService/internal/logger"
	"os"
)

func main() {
	defer os.Exit(1)
	l := logger.NewLogger("[CMD]")
	al := logger.NewLogger("[APP]")

	cfg, err := config.InitConfig()
	if err != nil {
		l.Error(err)
	}

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
