package models

import (
	"io"
	"fmt"
	"log"
	"bytes"
	"encoding/csv"
)

type Records [][]string

type RecordsInserter interface {
	InsertStopTimes(sts *StopTimes) (err error)
	InsertStops(sts *Stops) (err error)
}


func ParseCsv(b *[]byte) (*Records, error) {
	r := bytes.NewReader(*b)
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
