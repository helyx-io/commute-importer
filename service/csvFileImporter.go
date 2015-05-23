package service

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "fmt"
    "log"
    "sort"
    "github.com/helyx-io/commute-importer/utils"
    "github.com/helyx-io/commute-importer/database"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type CsvFileImporter struct {
    driver *database.Driver
    gtfs database.GTFSRepository
    repositoryByFilenameMap map[string]database.GTFSCreatedModelRepository
}

func NewCsvFileImporter(driver *database.Driver, gtfs database.GTFSRepository) *CsvFileImporter {
    repositories := initRepositoryMap(gtfs)
    return &CsvFileImporter{driver, gtfs, repositories}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Helper functions
////////////////////////////////////////////////////////////////////////////////////////////////

func initRepositoryMap(gtfs database.GTFSRepository) map[string]database.GTFSCreatedModelRepository {
    repositoryByFilenameMap := make(map[string]database.GTFSCreatedModelRepository)

    repositoryByFilenameMap["agency.txt"] = gtfs.Agencies()
    repositoryByFilenameMap["calendar_dates.txt"] = gtfs.CalendarDates()
    repositoryByFilenameMap["calendar.txt"] = gtfs.Calendars()
    repositoryByFilenameMap["routes.txt"] = gtfs.Routes()
    repositoryByFilenameMap["stops.txt"] = gtfs.Stops()
    repositoryByFilenameMap["stop_times.txt"] = gtfs.StopTimes()
    repositoryByFilenameMap["transfers.txt"] = gtfs.Transfers()
    repositoryByFilenameMap["trips.txt"] = gtfs.Trips()

    return repositoryByFilenameMap

}

////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func (cfi *CsvFileImporter) ImportCsvFiles(agencyKey, outFolderFilename string, columnLengthsByFiles map[string]map[string]int) {

    fis := utils.ReadDirectoryFileInfos(outFolderFilename)
    sort.Sort(utils.FileInfosBySize(fis))

    err := cfi.gtfs.CreateSchema(agencyKey)
    utils.FailOnError(err, fmt.Sprintf("Could not create schema for key: '%s'", agencyKey))

    for _, fi := range fis {
        if fi.Mode().IsRegular() {
            gtfsModelRepository := cfi.repositoryByFilenameMap[fi.Name()]

            if gtfsModelRepository == nil {
                log.Printf("Filename '%v' is not available in map", fi.Name())
                continue;
            }

            gaf := NewGTFSArchiveFile(fi)

            log.Printf("columnLengthsByFiles:  %v - gaf.Name(): %v", columnLengthsByFiles, gaf.Name())

            columnLengthsByFilesTmp := make(map[string]interface{})
            for key, value := range columnLengthsByFiles {
                columnLengthsByFilesTmp[key] = value
            }

            log.Printf("ImportGTFSArchiveFileWithTableCreation:  %v", columnLengthsByFilesTmp)

            err := gaf.ImportGTFSArchiveFileWithTableCreation(agencyKey, outFolderFilename, gtfsModelRepository, columnLengthsByFilesTmp, 512 * 1000)
            utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive with table creation for key: '%s'", agencyKey, fi.Name()))

            if fi.Name() == "agency.txt" {
                log.Println("Importing agencies in GTFS agencies table ...")

                gtfsAgencyModelRepository := cfi.gtfs.GtfsAgencies()
                gaf := NewGTFSArchiveFile(fi)

                err:= gaf.ImportGTFSArchiveFileWithoutTableCreation(agencyKey, outFolderFilename, gtfsAgencyModelRepository, 512 * 1000)
                utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive without table creation for key: '%s'", agencyKey, fi.Name()))
            }

        }
    }

}
