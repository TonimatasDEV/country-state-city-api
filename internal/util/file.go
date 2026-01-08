package util

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	file, err := os.Create(filepath)

	if err != nil {
		_ = resp.Body.Close()
		return err
	}

	_, err = io.Copy(file, resp.Body)
	_ = resp.Body.Close()
	_ = file.Close()

	if err != nil {
		return err
	}

	return nil
}
