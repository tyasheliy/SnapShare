package controllers

import (
	"LinkService/internal/cache"
	"LinkService/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type entryController struct {
	cache cache.Cache
}

type EntryController interface {
	CreateEntry(ctx echo.Context) error
}

func NewEntryController(cache cache.Cache) *entryController {
	return &entryController{
		cache: cache,
	}
}

func (c *entryController) CreateEntry(ctx echo.Context) error {
	// Reading request body.
	body := ctx.Request().Body
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

	// Creating a new entry with given password.
	entry := models.NewEntry(bodyMap["password"].(string))
	json, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	errCh := make(chan error)

	// Caching new entry.
	go c.cache.Cache(entry.ID, string(json), time.Minute, errCh)

	err = <-errCh
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{"id": entry.ID})
}
