package app

import (
	"LinkService/config"
	"LinkService/internal/controllers"
	"LinkService/internal/exptype"
	"fmt"
	"net/http"
	"strings"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	cfg     *config.Config
	e       *echo.Echo
	links   controllers.LinkController
	entries controllers.EntryController
}

func New(cfg *config.Config, links controllers.LinkController, entries controllers.EntryController) (*App, error) {
	e := echo.New()

	app := &App{
		cfg:     cfg,
		e:       e,
		links:   links,
		entries: entries,
	}

	//TODO: Work on logger.
	app.e.Use(middleware.CORS())
	app.e.Use(middleware.Logger())
	app.e.Use(middleware.Recover())
	app.e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:    cfg.SigningKey,
		SigningMethod: echojwt.AlgorithmHS256,
		Skipper: func(c echo.Context) bool {
			if c.Request().Method == http.MethodGet || c.Request().Method == http.MethodHead || strings.Contains(c.Request().URL.Path, "/entries") {
				return true
			}

			return false
		},
	}))

	app.e.GET("/links/:id", app.links.UseLink)
	app.e.HEAD("/links/:id", app.links.UseLink)
	app.e.POST("/links", app.links.CreateLink)
	app.e.POST("/links/:id", app.links.UploadFile)
	app.e.PUT("/links/:id", app.links.LoadChunk)

	app.e.POST("/entries", app.entries.CreateEntry)

	app.e.GET("/types", app.GetExpireTypes)

	return app, nil
}

func (a *App) GetExpireTypes(c echo.Context) error {
	return c.JSON(http.StatusOK, exptype.GetTypes())
}

func (a *App) Serve() {
	addr := fmt.Sprintf(":%s", a.cfg.Port)
	a.e.Start(addr)
}
