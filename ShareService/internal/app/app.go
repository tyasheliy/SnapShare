package app

import (
	"LinkService/config"
	"LinkService/internal/cache"
	"LinkService/internal/exptype"
	"LinkService/internal/files"
	"LinkService/internal/logger"
	"LinkService/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	cfg    *config.Config
	logger *logger.Logger
	cache  *cache.Cache
	e      *echo.Echo
}

func New(cfg *config.Config, cache *cache.Cache, logger *logger.Logger) (*App, error) {
	e := echo.New()

	//TODO: Work on logger.
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:    cfg.SigningKey,
		SigningMethod: echojwt.AlgorithmHS256,
		Skipper: func(c echo.Context) bool {
			allowUnauth := [...]string{fmt.Sprintf("/links/%s", c.Param("id")), "/entries"}
			for _, v := range allowUnauth {
				if strings.Contains(c.Request().URL.Path, v) {
					return true
				}
			}

			return false
		},
	}))

	return &App{
		cfg:    cfg,
		logger: logger,
		e:      e,
		cache:  cache,
	}, nil
}

func (a *App) createLink(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrBadRequest
	}

	if claims["userId"] == nil {
		return echo.ErrBadRequest
	}

	userId := claims["userId"].(string)

	header, err := c.FormFile("file")
	if err != nil {
		return err
	}

	password := c.FormValue("password")
	expireType := c.FormValue("expireType")

	link, err := models.NewLink(userId, header, exptype.ExpireType(expireType), password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Link with the given data already exists"})
	}

	json, err := json.Marshal(link)
	if err != nil {
		return err
	}

	errCh := make(chan error)

	go a.cache.Cache(link.ID, string(json), link.GetDuration(), errCh)

	err = <-errCh
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"id": link.ID})
}

func (a *App) consumeLink(c echo.Context) error {
	errCh := make(chan error)
	cacheCh := make(chan string)

	go a.cache.Get(c.Param("id"), cacheCh, errCh)

	err := <-errCh
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Link is not found"})
	}

	var link models.Link
	err = json.Unmarshal([]byte(<-cacheCh), &link)
	if err != nil {
		return err
	}

	if link.Password != "" {
		if c.QueryParam("entry") == "" {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "This link is password protected"})
		}

		go a.cache.Get(c.QueryParam("entry"), cacheCh, errCh)

		err = <-errCh
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Entry is not found"})
		}

		var entry models.Entry
		err = json.Unmarshal([]byte(<-cacheCh), &entry)
		if err != nil {
			return err
		}

		go a.cache.Delete(entry.ID, errCh)

		err = <-errCh
		if err != nil {
			a.e.Logger.Error(err)
		}

		if link.Password != entry.Password {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Invalid password"})
		}
	}

	if !files.FileExists(link.FileName) {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "File is not found"})
	}

	return c.Attachment(fmt.Sprintf("%s%s", files.GetDir(), link.FileName), link.FileName)
}

func (a *App) createEntry(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	bodyData, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	var bodyMap map[string]interface{}
	err = json.Unmarshal(bodyData, &bodyMap)
	if err != nil {
		return err
	}

	if bodyMap["password"] == nil {
		return echo.ErrBadRequest
	}

	entry := models.NewEntry(bodyMap["password"].(string))
	json, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	errCh := make(chan error)

	go a.cache.Cache(entry.ID, string(json), time.Minute, errCh)

	err = <-errCh
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"id": entry.ID})
}

func (a *App) getExpireTypes(c echo.Context) error {
	return c.JSON(http.StatusOK, exptype.GetTypes())
}

func (a *App) Serve() {
	a.e.GET("/links/:id", a.consumeLink)
	a.e.POST("/links", a.createLink)

	a.e.POST("/entries", a.createEntry)

	a.e.GET("/types", a.getExpireTypes)

	addr := fmt.Sprintf(":%s", a.cfg.Port)
	a.e.Start(addr)
}
