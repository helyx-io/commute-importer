package controller

import (
	"archive/zip"
	"io"
	"os"
	"fmt"
	"bufio"
	"path/filepath"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/csv"
	"bytes"
)

type ImportController struct { }

func (importController *ImportController) Init(r *mux.Router) {
	r.HandleFunc("/", importController.Import)
}

func (ac *ImportController) Import(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("Importing agencies ..."))
	w.Write([]byte("<br/>"))

	url := "http://localhost/data/gtfs_paris_20140502.zip"
	zipFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502.zip"
	folderFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502"

	w.Write([]byte(fmt.Sprintf(" - Downloading '%v' to '%v' ...", url, zipFilename)))
	w.Write([]byte("<br/>"))

	out, err := os.Create(zipFilename)
	defer out.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	writtenBytes, err := io.Copy(out, resp.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(fmt.Sprintf(" - Downloaded zip file: '%v' - %v bytes", zipFilename, writtenBytes)))
	w.Write([]byte("<br/>"))

	w.Write([]byte(fmt.Sprintf(" - Unzipping file: '%v' to directory: '%v' ...", zipFilename, folderFilename)))
	w.Write([]byte("<br/>"))

	err = Unzip(zipFilename, folderFilename)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(fmt.Sprintf(" - Unzipped file: '%v' to directory: '%v'", zipFilename, folderFilename)))
	w.Write([]byte("<br/>"))

	stopsFilename := fmt.Sprintf("%v/stop_times.txt", folderFilename)

	w.Write([]byte(fmt.Sprintf(" - Reading file: '%v'", stopsFilename)))
	w.Write([]byte("<br/>"))

	err = ReadFile(stopsFilename)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(fmt.Sprintf(" - 	Read file: '%v'", stopsFilename)))
	w.Write([]byte("<br/>"))
}


func ReadFile(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	ScanNthLine := func(linesCount int) bufio.SplitFunc {

		// ScanLines is a split function for a Scanner that returns each line of
		// text, stripped of any trailing end-of-line marker. The returned line may
		// be empty. The end-of-line marker is one optional carriage return followed
		// by one mandatory newline. In regular expression notation, it is `\r?\n`.
		// The last non-empty line of input will be returned even if it has no
		// newline.
		return func(data []byte, atEOF bool) (advance int, token []byte, err error) {

			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}

			offset := linesCount
			lastIndex := 0
			for {
				offset--
				if i := bytes.IndexByte(data[lastIndex:], '\n'); i >= 0 {

					if offset <= 0 {
						// We have a full newline-terminated line.
						return lastIndex + i + 1, data[0:lastIndex + i], nil
					} else {
						lastIndex = lastIndex + i + 1
					}
				} else {
					break;
				}
			}

			// If we're at EOF, we have a final, non-terminated line. Return it.
			if atEOF {
				return len(data), data, nil
			}

			// Request more data.
			return 0, nil, nil
		}
	}

	scanner := bufio.NewScanner(reader)

	scanner.Split(ScanNthLine(100))

	i := 0
	for scanner.Scan() {
		i++
		text := scanner.Text()
		fmt.Println(i, " - ", len(text))
	}
	if err {
		fmt.Println("Scanner Err:", scanner.Err())
	}

	return nil
}


func ReadCsvFile(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Println(record)
	}

	return nil
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
