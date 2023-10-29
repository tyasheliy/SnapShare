package models

import (
	"LinkService/internal/exptype"
	"LinkService/internal/identifiers"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type Link struct {
	ID                 string `json: "id"`
	UserID             string `json: "userId"`
	exptype.ExpireType `json: "expireType"`
	Size               int64  `json: "size"`
	Password           string `json: "password"`
	FileName           string `json: "fileName"`
}

const (
	chunkSize              = 1000000
	minFileSizeForChunking = 5 * 1000000
)

func NewLink(userId string, fileName string, size int64, expireType exptype.ExpireType, password string) (*Link, error) {
	id := identifiers.GenerateLinkID()

	err := os.Mkdir(fmt.Sprintf("../tmp/%s", id), 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(fmt.Sprintf("../tmp/%s/%s", id, fileName))
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(size-1, 0)
	if err != nil {
		return nil, err
	}

	_, err = file.Write([]byte{0})
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	link := &Link{
		ID:         id,
		UserID:     userId,
		ExpireType: expireType,
		Password:   password,
		FileName:   fileName,
		Size:       size,
	}

	err = os.Remove(file.Name())
	if err != nil {
		return nil, err
	}

	go func(dur time.Duration) {
		time.Sleep(dur)
		link.RemoveLink()
	}(expireType.GetDuration())

	return link, nil
}

func (l *Link) IsChunked() bool {
	return l.Size >= minFileSizeForChunking
}

func (l *Link) RemoveLink() error {
	err := os.RemoveAll(fmt.Sprintf("../tmp/%s", l.ID))
	return err
}

func (l *Link) UploadFile(header *multipart.FileHeader) error {
	f, err := header.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	file, err := os.Create(fmt.Sprintf("../tmp/%s/%s", l.ID, header.Filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, f)
	return err
}

func (l *Link) UploadChunk(chunkNumber int, data []byte) error {
	if chunkNumber > int(l.Size)/chunkSize {
		return errors.New("Chunk number is too high")
	}

	if len(data) > chunkSize {
		return errors.New("Invalid chunk size")
	}

	file, err := os.OpenFile(fmt.Sprintf("../tmp/%s/%s", l.ID, l.FileName), os.O_WRONLY, os.ModeAppend)
	if err != nil {
		file, err = os.Create(fmt.Sprintf("../tmp/%s/%s", l.ID, l.FileName))
		if err != nil {
			return nil
		}
	}
	defer file.Close()

	_, err = file.Seek(int64(chunkNumber*chunkSize), io.SeekCurrent)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
