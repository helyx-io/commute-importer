package utils

import (
	"io"
	"os"
	"bufio"
)

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

		if len(b) >= 256000 {
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
