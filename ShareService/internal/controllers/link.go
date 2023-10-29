package controllers

import (
	"LinkService/internal/cache"
	"LinkService/internal/exptype"
	"LinkService/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type linkController struct {
	cache cache.Cache
}

type LinkController interface {
	UseLink(ctx echo.Context) error
	CreateLink(ctx echo.Context) error
	UploadFile(ctx echo.Context) error
	LoadChunk(ctx echo.Context) error
}

func NewLinkController(cache cache.Cache) LinkController {
	return &linkController{
		cache: cache,
	}
}

func (c *linkController) UseLink(ctx echo.Context) error {
	// Getting link's data from cache.
	errCh := make(chan error)
	cacheCh := make(chan string)

	go c.cache.Get(ctx.Param("id"), cacheCh, errCh)

	err := <-errCh
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Link is not found"})
	}

	// Parsing link's data.
	var link models.Link
	err = json.Unmarshal([]byte(<-cacheCh), &link)
	if err != nil {
		return err
	}

	// If link has password protection.
	if link.Password != "" {
		// If entry is not provided.
		if ctx.QueryParam("entry") == "" {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "This link is password protected"})
		}

		// Getting entry from cache.
		go c.cache.Get(ctx.QueryParam("entry"), cacheCh, errCh)

		err = <-errCh
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Entry is not found"})
		}

		// Parsing entry.
		var entry models.Entry
		err = json.Unmarshal([]byte(<-cacheCh), &entry)
		if err != nil {
			return err
		}

		// Delete entry if request
		if ctx.Request().Method == http.MethodGet {
			go c.cache.Delete(entry.ID, errCh)

			err = <-errCh
			if err != nil {
				return err
			}
		}

		if link.Password != entry.Password {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Invalid password"})
		}
	}

	// Checking if file exists.
	_, err = os.Open(fmt.Sprintf("../tmp/%s/%s", link.ID, link.FileName))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "File is not found"})
	}

	// Send response with file as attachment.
	return ctx.Attachment(fmt.Sprintf("../tmp/%s/%s", link.ID, link.FileName), link.FileName)
}

func (c *linkController) CreateLink(ctx echo.Context) error {
	// Getting token.
	user := ctx.Get("user").(*jwt.Token)

	// Reading claims.
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrBadRequest
	}

	if claims["userId"] == nil {
		return echo.ErrBadRequest
	}

	userId := claims["userId"].(string)

	// Reading body.
	body := ctx.Request().Body
	defer body.Close()

	bodyData, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	// Parsing body.
	var bodyMap map[string]interface{}
	err = json.Unmarshal(bodyData, &bodyMap)
	if err != nil {
		return err
	}

	// Checking body for required data.
	if bodyMap["fileName"] == nil || bodyMap["size"] == nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Not all necessary data is given"})
	}

	// Set empty password if has not got one.
	if bodyMap["password"] == nil {
		bodyMap["password"] = ""
	}

	// Getting data from parsed body.
	size := int64(bodyMap["size"].(float64))
	fileName := bodyMap["fileName"].(string)
	password := bodyMap["password"].(string)
	expireType := bodyMap["expireType"].(string)

	// Creating new link.
	link, err := models.NewLink(userId, fileName, size, exptype.ExpireType(expireType), password)
	if err != nil {
		return err
	}

	// Encoding new link in JSON.
	json, err := json.Marshal(link)
	if err != nil {
		return err
	}

	// Caching new link.
	errCh := make(chan error)

	go c.cache.Cache(link.ID, string(json), link.GetDuration(), errCh)

	err = <-errCh
	if err != nil {
		return err
	}

	// Return response with the link ID.
	return ctx.JSON(http.StatusOK, echo.Map{"id": link.ID})
}

func (c *linkController) UploadFile(ctx echo.Context) error {
	// Reading token.
	user := ctx.Get("user").(*jwt.Token)
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrBadRequest
	}

	if claims["userId"] == nil {
		return echo.ErrBadRequest
	}

	userId := claims["userId"].(string)
	linkId := ctx.Param("id")

	// Checking if link exists.
	cacheCh := make(chan string)
	errCh := make(chan error)

	go c.cache.Get(linkId, cacheCh, errCh)

	err := <-errCh
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Link is not found"})
	}

	// Reading link data.
	linkData := []byte(<-cacheCh)
	var link models.Link

	// Parsing link.
	err = json.Unmarshal(linkData, &link)
	if err != nil {
		return err
	}

	// If link's owner is not user.
	if link.UserID != userId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "You do not have persmission to do this"})
	}

	// If the link is too big for single file uploading.
	if link.IsChunked() {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "This link is chunked"})
	}

	// Checking if link's directory already has a file.
	entries, err := os.ReadDir(fmt.Sprintf("../tmp/%s", link.ID))
	if err != nil {
		return err
	}

	if len(entries) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "This link already has a file"})
	}

	// Getting request file.
	header, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	// Checking if file has valid filename and size.
	if header.Filename != link.FileName {
		return echo.ErrBadRequest
	}

	if header.Size != link.Size {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Link has other size"})
	}

	// Uploading file.
	err = link.UploadFile(header)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *linkController) LoadChunk(ctx echo.Context) error {
	// Reading token.
	user := ctx.Get("user").(*jwt.Token)
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrBadRequest
	}

	if claims["userId"] == nil {
		return echo.ErrBadRequest
	}

	// Getting the link.
	userId := claims["userId"].(string)
	linkId := ctx.Param("id")

	cacheCh := make(chan string)
	errCh := make(chan error)

	go c.cache.Get(linkId, cacheCh, errCh)

	err := <-errCh
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Link is not found"})
	}

	// Reading link.
	linkData := []byte(<-cacheCh)
	var link models.Link

	// Parsing link data.
	err = json.Unmarshal(linkData, &link)
	if err != nil {
		return err
	}

	// Checking if user is link's owner.
	if link.UserID != userId {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "You do not have persmission to do this"})
	}

	// If link supposed to contain a small file.
	if !link.IsChunked() {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "This link is not chunked"})
	}

	// Getting request chunk.
	chunkHeader, err := ctx.FormFile("data")
	if err != nil {
		return err
	}

	chunk, err := chunkHeader.Open()
	if err != nil {
		return err
	}
	defer chunk.Close()

	// Getting chunk's number in sequence.
	number, err := strconv.Atoi(ctx.FormValue("number"))
	if err != nil {
		return err
	}

	// Reading request chunk.
	data, err := io.ReadAll(chunk)
	if err != nil {
		return err
	}

	// Loading chunk in link's file.
	err = link.UploadChunk(number, data)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
