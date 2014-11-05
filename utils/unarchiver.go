package utils

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"github.com/fatih/stopwatch"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Private Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func UnzipArchive(zipFilename, folderFilename string) {

	log.Println(fmt.Sprintf(" - Unzipping file: '%v' to directory: '%v' ...", zipFilename, folderFilename))

	sw := stopwatch.Start(0)

	err := Unzip(zipFilename, folderFilename)
	FailOnError(err, fmt.Sprintf("Could unzip filename: '%v' to folder: '%v'", zipFilename, folderFilename))

	log.Println(fmt.Sprintf(" - Unzipped file: '%v' to directory: '%v' - Duration: %v", zipFilename, folderFilename, sw.ElapsedTime()))
}
