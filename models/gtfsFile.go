package models

import (
	"github.com/helyx-io/gtfs-playground/utils"
)

type GTFSFile struct {
	Filename string
}

func (gf GTFSFile) LinesIterator(maxLength int) <- chan []byte {
	channel := make(chan []byte)
	go func() {
		utils.ReadCsvFile(gf.Filename, maxLength, channel)
		defer close(channel)
	}()
	return channel
}
