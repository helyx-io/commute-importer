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

func (gaf *GTFSArchiveFile) ImportGTFSArchiveFile(folderFilename string, gtfsModelRepository database.GTFSModelRepository, workPool *workpool.WorkPool) {

	log.Println(fmt.Sprintf("Filename '%v' is available in map - Reading File with size: %d bytes ...", gaf.Name(), gaf.Size()))

	sw := stopwatch.Start(0)

	offset := 0

	log.Println(fmt.Sprintf(" - Removing entries from repository related to file with name: '%v' ...", gaf.Name()))
	gtfsModelRepository.RemoveAllByAgencyKey("RATP")
	log.Println(fmt.Sprintf(" - Removed entries from repository related to file with name: '%v'", gaf.Name()))

	gtfsFile := models.GTFSFile{path.Join(folderFilename, gaf.Name())}

	for lines := range gtfsFile.LinesIterator() {

		offset++

		log.Println(fmt.Sprintf(" - Inserting chunk of data with offset: '%d' related to file with name: '%v'", offset, gaf.Name()))

		taskName := fmt.Sprintf("ChunkImport-%d", offset)
		task := gtfsModelRepository.CreateImportTask(taskName, "RATP", lines, workPool)

		err := workPool.PostWork("import", task)

		utils.FailOnError(err, fmt.Sprintf("Could not post work with offset: %d", offset))
	}

	log.Println(fmt.Sprintf(" - 	Read file: '%v' - Duration: %v", gaf.Name(), sw.ElapsedTime()))
}
