package utils

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io"
	"os"
	"encoding/json"
	"net/http"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Helper Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func SendJSON(w http.ResponseWriter, data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if  err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

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
