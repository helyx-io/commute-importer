package utils

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"fmt"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type FileInfosBySize []os.FileInfo

func (fis FileInfosBySize) Len() int {
	return len(fis)
}
func (fis FileInfosBySize) Less(i, j int) bool {
	return fis[i].Size() < fis[j].Size()
}
func (fis FileInfosBySize) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Private Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func ReadDirectoryFileInfos(folderFilename string) []os.FileInfo {

	d, err := os.Open(folderFilename)
	FailOnError(err, fmt.Sprintf("Could not open directory '%v' for read", folderFilename))
	defer d.Close()

	fisArr, err := d.Readdir(-1)
//	fis := FileInfos(fisArr)
	FailOnError(err, fmt.Sprintf("Could not read directory '%v' content", folderFilename))

	return fisArr
}
