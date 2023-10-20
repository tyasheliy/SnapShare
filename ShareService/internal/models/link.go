package models

import (
	"LinkService/internal/exptype"
	"LinkService/internal/files"
	"LinkService/internal/identifiers"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
)

type Link struct {
	ID                 string `json: "id"`
	UserID             string `json: "userId"`
	exptype.ExpireType `json: "expireType"`
	Password           string `json: "password"`
	FileName           string `json: "fileName"`
}

func NewLink(userId string, header *multipart.FileHeader, expireType exptype.ExpireType, password string) (Link, error) {
	file, err := header.Open()
	if err != nil {
		return Link{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Link{}, err
	}

	var fileName string

	if files.FileExists(header.Filename) {
		fdata, err := files.ReadTempFile(header.Filename)
		if err != nil {
			return Link{}, err
		}

		if string(data) == string(fdata) {
			return Link{}, errors.New("Link with this data already exists")
		}

		fileName = fmt.Sprintf("%s%s", "another_", header.Filename)
	} else {
		fileName = header.Filename
	}

	err = files.CreateTempFile(fileName, data, expireType.GetDuration())
	if err != nil {
		return Link{}, err
	}

	return Link{
		ID:         identifiers.GenerateLinkID(),
		UserID:     userId,
		ExpireType: expireType,
		Password:   password,
		FileName:   fileName,
	}, nil
}
