package models

import (
	"io"
	"fmt"
	"log"
	"bytes"
    "errors"
    "encoding/csv"
    "github.com/helyx-io/gtfs-importer/csv/length"
)

type Records [][]string


func ParseCsvAsIntArrays(b []byte) ([]int, error) {

    r := bytes.NewReader(b)
    reader := length.NewReader(r)
    records := make([][]int, 0)

    for {
        record, err := reader.Read()

        if err == io.EOF {
            break
        } else if err, ok := err.(*csv.ParseError); ok {
            if err.Err != csv.ErrFieldCount {
                fmt.Println(fmt.Sprintf("%#v", err))
                log.Println("2 - Error on line read:", err, "line:", record)
                panic(err)
            }
        } else if err != nil {
            fmt.Println(fmt.Sprintf("%#v", err))
            log.Println("3 - Error on line read:", err, "line:", record)
            break;
        }

        records = append(records, record)
    }

    if len(records) > 0 {
        lengthRecord := make([]int, len(records[0]))
        for _, record := range records {
            for i, field := range record {
                if field > lengthRecord[i] {
                    lengthRecord[i] = field
                }
            }
        }

        return lengthRecord, nil
    } else {
        return nil, errors.New("No lines to count field lengths")
    }
}




func ParseCsvAsStringArrays(b []byte) (*[][]string, error) {
    r := bytes.NewReader(b)
    reader := csv.NewReader(r)
    records := make([][]string, 0)

    var err error

    for {
        record, err := reader.Read()

        if err == io.EOF {
            break
        } else if err, ok := err.(*csv.ParseError); ok {
            if err.Err != csv.ErrFieldCount {
                fmt.Println(fmt.Sprintf("%#v", err))
                log.Println("2 - Error on line read:", err, "line:", record)
                panic(err)
            }
        } else if err != nil {
            fmt.Println(fmt.Sprintf("%#v", err))
            log.Println("3 - Error on line read:", err, "line:", record)
            break;
        }

        records = append(records, record)
    }

    return &records, err
}

func ParseCsv(b []byte) (*Records, error) {
	r := bytes.NewReader(b)
	reader := csv.NewReader(r)
	records := make(Records, 0)

	var err error

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err, ok := err.(*csv.ParseError); ok {
			if err.Err != csv.ErrFieldCount {
				fmt.Println(fmt.Sprintf("%#v", err))
				log.Println("2 - Error on line read:", err, "line:", record)
				panic(err)
			}
		} else if err != nil {
			fmt.Println(fmt.Sprintf("%#v", err))
			log.Println("3 - Error on line read:", err, "line:", record)
			break;
		}

		records = append(records, record)
	}

	return &records, err
}
