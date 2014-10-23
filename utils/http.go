package utils

import (
	"io"
	"os"
	"net/http"
)

func DownloadFileFromURL(url, destPath string) (int64, error) {
	out, err := os.Create(destPath)
	defer out.Close()

	if err != nil {
		return -1, err
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return -1, err
	}

	writtenBytes, err := io.Copy(out, resp.Body)

	if err != nil {
		return -1, err
	}

	return writtenBytes, err
}
