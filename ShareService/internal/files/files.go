package files

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	tmpDir  = "../tmp/"
	dirName = "snapshare_share-service"
)

func InitFiles() {
	_ = os.Mkdir(tmpDir, 0755)
}

func GetDir() string {
	return fmt.Sprintf("%s%s", tmpDir, dirName)
}

func FileExists(fileName string) bool {
	_, err := os.Open(fmt.Sprintf("%s%s", GetDir(), fileName))
	return err == nil
}

func setFileTimer(ctx context.Context, fileName string, dur time.Duration) {
	time.Sleep(dur)
	os.Remove(fmt.Sprintf("%s%s", GetDir(), fileName))
}

func createFile(fileName string, data []byte) error {
	if FileExists(fileName) {
		return errors.New("file exists")
	}

	file, err := os.Create(fmt.Sprintf("%s%s", GetDir(), fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprint(file, data)
	if err != nil {
		return err
	}

	return nil
}

func ReadTempFile(fileName string) ([]byte, error) {
	file, err := os.Open(fmt.Sprintf("%s%s", GetDir(), fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CreateTempFile(fileName string, data []byte, duration time.Duration) error {
	err := createFile(fileName, data)
	if err != nil {
		return err
	}

	go setFileTimer(context.Background(), fileName, duration)

	return nil
}
