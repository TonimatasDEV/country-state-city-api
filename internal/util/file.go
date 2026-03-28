package util

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	contentDisposition := resp.Header.Get("Content-Disposition")
	filename := strings.Split(contentDisposition, "=")[1]
	file, err := os.Create(filename)

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

	return gunzip(filename, filepath)
}

func gunzip(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer fileClose(file)

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fileClose(out)

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	defer gzipReaderClose(gzReader)

	_, err = io.Copy(out, gzReader)
	return err
}
