package utils

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io"
	"os"
    "log"
	"encoding/json"
	"net/http"
    "runtime/debug"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Helper Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func RecoverFromError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err, _ := r.(error)
        log.Println("Err:", err.Error());
        debug.PrintStack()
		http.Error(w, err.Error(), 500)
		return
	}
}

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
