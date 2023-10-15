package models

import (
	"LinkService/internal/exptype"
	"LinkService/internal/identifiers"
	"io"
	"mime/multipart"
)

type Link struct {
	ID                 string `json: "id"`
	UserID             string `json: "userId"`
	exptype.ExpireType `json: "expireType"`
	Password           string `json: "password"`
	FileName           string `json: "fileName"`
	FileData           []byte `json: "fileData"`
}

func NewLink(userId string, header *multipart.FileHeader, expireType exptype.ExpireType, password string) (Link, error) {
	f, err := header.Open()
	if err != nil {
		return Link{}, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return Link{}, err
	}

	return Link{
		ID:         identifiers.GenerateLinkID(),
		UserID:     userId,
		ExpireType: expireType,
		Password:   password,
		FileName:   header.Filename,
		FileData:   data,
	}, nil
}
