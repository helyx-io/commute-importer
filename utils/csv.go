package utils

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io"
	"os"
	"log"
	"bufio"
	"strings"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Helper Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func ReadCsvFile(src string, maxLength int, channel chan []byte) {

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

		if len(b) >= maxLength {
			chunk++
            log.Printf("Chunk Index: %d - Number of lines : %d", chunk, i)

			channel <- fixUTF8BomIfNecessary(b)

			i = 0
			b = []byte{}
		}

	}

	if len(b) > 0 {
		chunk++
		log.Printf("Chunk Index: %d - Number of lines : %d", chunk, i)

		channel <- fixUTF8BomIfNecessary(b)
	}

}

func fixUTF8BomIfNecessary(data []byte) []byte {
	if len(data) >= 3 && data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
		return data[3:]
	} else if len(data) >= 6 && data[0] == 0xc3 && data[1] == 0xaf && data[2] == 0xc2 && data[3] == 0xbb && data[4] == 0xc2 && data[5] == 0xbf {
		return data[6:]
	} else {
		return data
	}
}


func ReadCsvFileHeader(src string, separator string) ([]string, error) {

	file, err := os.Open(src)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	headers, err := bufio.NewReader(file).ReadBytes('\n')

	if err != nil {
		return nil, err
	}

    headersFixed := fixUTF8BomIfNecessary(headers)
	headerStr := string(headersFixed[0:len(headersFixed) - 1])
	return strings.Split(headerStr, ","), nil
}
