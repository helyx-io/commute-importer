package service

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"fmt"
	"log"
	"path"
	"github.com/fatih/stopwatch"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

func NewGTFSArchiveFile(fi os.FileInfo) *GTFSArchiveFile {
	return &GTFSArchiveFile{fi}
}

type GTFSArchiveFile struct {
	fi os.FileInfo
}

func (gaf *GTFSArchiveFile) Name() string {
	return gaf.fi.Name()
}

func (gaf *GTFSArchiveFile) Size() int64 {
	return gaf.fi.Size()
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Import Service
////////////////////////////////////////////////////////////////////////////////////////////////


func (gaf *GTFSArchiveFile) ImportGTFSArchiveFile(agencyKey string, folderFilename string, gtfsModelRepository database.GTFSModelRepository, maxLength int, workPool *workpool.WorkPool) {

	log.Println(fmt.Sprintf("Filename '%v' is available in map - Reading File with size: %d bytes ...", gaf.Name(), gaf.Size()))

	sw := stopwatch.Start(0)

	offset := 0

	log.Println(fmt.Sprintf(" - Removing entries from repository related to file with name: '%v' ...", gaf.Name()))

	err := gtfsModelRepository.RemoveAllByAgencyKey(agencyKey)
	utils.FailOnError(err, fmt.Sprintf("Could not remove table for file with name: '%v'", gaf.Name()))

	err = gtfsModelRepository.CreateTableByAgencyKey(agencyKey)
	utils.FailOnError(err, fmt.Sprintf("Could not create table for file with name: '%v'", gaf.Name()))

	log.Println(fmt.Sprintf(" - Removed entries from repository related to file with name: '%v'", gaf.Name()))

	gtfsFile := models.GTFSFile{path.Join(folderFilename, gaf.Name())}

	headers, err := utils.ReadCsvFileHeader(gtfsFile.Filename, ",")
	utils.FailOnError(err, fmt.Sprintf("Could not read csv headers for file with name: '%v'", gaf.Name()))


	// Init WorkPool

	doneChan := make(chan error, 32)

	for lines := range gtfsFile.LinesIterator(maxLength) {

		offset++

		log.Println(fmt.Sprintf(" - Inserting chunk of data with offset: '%d' related to file with name: '%v'", offset, gaf.Name()))

		taskName := fmt.Sprintf("ChunkImport-%d for file with name: '%v'", offset, gaf.Name())
		task := gtfsModelRepository.CreateImportTask(taskName, offset, gaf.Name(), agencyKey, headers, lines, workPool, doneChan)

		err := workPool.PostWork("import", task)

		utils.FailOnError(err, fmt.Sprintf("Could not post work with offset: %d", offset))
	}
	log.Println(fmt.Sprintf(" - 	Read file: '%v' - Duration: %v", gaf.Name(), sw.ElapsedTime()))

	doneCount := 0
	for err := range doneChan {
		if err != nil {
			log.Println(fmt.Sprintf("Received event on done chan with error: %s", err))
		} else {
			doneCount += 1
			if offset == doneCount {
//				log.Println(fmt.Sprintf("offset (%d) = done (%d)", offset, doneCount))
//				log.Println(fmt.Sprintf("Closing done chan"))
				close(doneChan)
			} else {
//				log.Println(fmt.Sprintf("Received event on done chan."))
			}
		}
	}

	if err != nil {
		log.Printf("[ERROR] %s", err)
	}

	log.Println(fmt.Sprintf("Adding indexes for file: '%v'", gaf.Name()))

	err = gtfsModelRepository.AddIndexesByAgencyKey(agencyKey)
	utils.FailOnError(err, fmt.Sprintf("Could not add indexes for file: '%v'", gaf.Name()))

	log.Println(fmt.Sprintf("Indexes created for file: '%v'", gaf.Name()))

	log.Println(fmt.Sprintf("All done in for file: '%v' - Duration: %v", gaf.Name(), sw.ElapsedTime()))
}
