package models

import (
	"github.com/helyx-io/gtfs-importer/utils"
)

type GTFSFile struct {
	Filename string
}

func (gf GTFSFile) LinesIterator(maxLength int) <- chan []byte {
	channel := make(chan []byte)
	go func() {
		err := utils.ReadCsvFile(gf.Filename, maxLength, channel)
        if err != nil {
            panic (err.Error())
        }

		defer close(channel)
	}()
	return channel
}
