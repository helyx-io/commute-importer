package utils

import (
	"io"
	"os"
	"fmt"
	"log"
	"bytes"
	"bufio"
	"encoding/csv"
	"github.com/akinsella/go-playground/models"
)


func ParseCsv(b []byte) (models.Records, error) {
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

	return models.Records{ records }, err
}

func ReadCsvFile(src string, channel chan []byte) {

	file, err := os.Open(src)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	b := []byte{}
	chunk := 0
	i := 0

	for {
		l, err := r.ReadBytes('\n')

		if err == io.EOF {
			break;
		}

		if err != nil {
			panic(err)
		}

		if len(l) == 0 {
			break;
		}

		if (chunk != 0 || i != 0) {
			b = append(b, l...)
			b = append(b, '\n')
		}

		i++

		if len(b) >= 512000 {
			chunk++
			//			fmt.Println("Chunk Index: ", chunk, "Number of lines :", i)
			channel <- b
			i = 0
			b = []byte{}
		}

	}

	if len(b) > 0 {
		chunk++
		//		fmt.Println("Chunk Index: ", chunk, "Number of lines :", i)
		channel <- b
	}

}
