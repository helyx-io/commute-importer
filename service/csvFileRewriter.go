package service

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "os"
    "fmt"
    "log"
    "path"
    "encoding/csv"
    "github.com/helyx-io/gtfs-importer/utils"
    "github.com/helyx-io/gtfs-importer/models"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type CsvFileRewriter struct {
    tmpDir string
}

func NewCsvFileRewriter(tmpDir string) *CsvFileRewriter {
    return &CsvFileRewriter{tmpDir}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func (cfr *CsvFileRewriter) RewriteCsvFiles(schema, outFolderName string) error {

    agencyIndexes, err := cfr.getIndexes(schema, "agency.txt", 0)
    servcfreIndexes, err := cfr.getIndexes(schema, "trips.txt", 1)
    tripIndexes, err := cfr.getIndexes(schema, "trips.txt", 2)
    stopIndexes, err := cfr.getIndexes(schema, "stops.txt", 0)
    routeIndexes, err := cfr.getIndexes(schema, "routes.txt", 0)

    cfr.writeIndexes(schema, "routes.indexes.txt", outFolderName, routeIndexes);
    cfr.writeIndexes(schema, "trip.indexes.txt", outFolderName, tripIndexes);
    cfr.writeIndexes(schema, "stop.indexes.txt", outFolderName, stopIndexes);

    folderFilename := cfr.tmpDir + "/" + schema
    outFolderFilename := path.Join(folderFilename, outFolderName)

    if os.MkdirAll(outFolderFilename, 0755) != nil {
        panic("Unable to create directory for tagfile!")
    }

    indexes := map[int](map[string]string){ 0: stopIndexes }
    cfr.rewriteCsvFile(schema, "stops.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: tripIndexes, 3: stopIndexes }
    cfr.rewriteCsvFile(schema, "stop_times.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: routeIndexes, 1: agencyIndexes }
    cfr.rewriteCsvFile(schema, "routes.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: agencyIndexes }
    cfr.rewriteCsvFile(schema, "agency.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: routeIndexes, 1: servcfreIndexes, 2: tripIndexes }
    cfr.rewriteCsvFile(schema, "trips.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: servcfreIndexes }
    cfr.rewriteCsvFile(schema, "calendar.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: servcfreIndexes }
    cfr.rewriteCsvFile(schema, "calendar_dates.txt", outFolderName, indexes)

    return err
}


func (cfr *CsvFileRewriter) writeIndexes(schema, filename, outFolderName string, indexes map[string]string) error {

    folderName := path.Join(cfr.tmpDir, schema)

    outFile, err := os.Create(path.Join(folderName, path.Join(outFolderName, filename)))
    if err != nil {
        log.Printf("Error: '%v'", err.Error())
        return err
    }

    log.Printf("Writing to file: '%v'", outFile.Name())

    writer := csv.NewWriter(outFile)

    for key, value := range indexes {
        writer.Write([]string{key, value})
    }

    writer.Flush()
    err = outFile.Close()

    return err
}


func (cfr *CsvFileRewriter) rewriteCsvFile(schema, filename, outFolderName string, indexes map[int](map[string]string)) error {

    folderName := path.Join(cfr.tmpDir, schema)
    filePath := path.Join(folderName, filename)

    gtfsFile := models.GTFSFile{filePath}

    outFile, err := os.Create(path.Join(folderName, path.Join(outFolderName, filename)))
    if err != nil {
        log.Printf("Error: %v", err.Error())
        return err
    }

    log.Printf("Writing to file: '%v'", outFile.Name())

    writer := csv.NewWriter(outFile)
    headers, err := utils.ReadCsvFileHeader(gtfsFile.Filename, ",")
    log.Printf("headers: '%v'", headers)

    if err != nil {
        log.Printf("Error: '%v'", err.Error())
        return err
    }

    writer.Write(headers)

    resultChan := make(chan [][]string)

    go func() {

        //        sem := make(chan bool, 8)

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            //            sem <- true
            //            go func() {
            //                defer func() { <-sem }()
            records, _ := models.ParseCsvAsStringArrays(lines)

            for _, record := range *records {
                for i, _ := range indexes {
                    record[i] = indexes[i][record[i]]
                }
            }

            resultChan <- *records
            //            }()
        }

        //        for i := 0; i < cap(sem); i++ {
        //            sem <- true
        //        }

        close(resultChan)

    }()

    offset := 0

    for results := range resultChan {
        offset++
        log.Printf("[%s][%d] Records write", filePath, offset)

        writer.WriteAll(results)
    }

    writer.Flush()
    err = outFile.Close()

    return err
}

func (cfr *CsvFileRewriter) getIndexes(schema, filename string, index int) (map[string]string, error) {

    folderName := path.Join(cfr.tmpDir, schema)
    filePath := path.Join(folderName, filename)

    gtfsFile := models.GTFSFile{filePath}

    resultChan := make(chan []string)

    go func() {

        //        sem := make(chan bool, 8)

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            //            sem <- true
            //            go func() {
            //                defer func() { <-sem }()
            records, _ := models.ParseCsvAsStringArrays(lines)

            keys := []string{}
            for _, record := range *records {
                keys = append(keys, record[index])
            }

            resultChan <- keys
            //            }()
        }

        //        for i := 0; i < cap(sem); i++ {
        //            sem <- true
        //        }

        close(resultChan)

    }()

    offset := 0
    result := []string{}
    indexes := map[string]string{}
    increment := 0

    for results := range resultChan {
        offset++
        log.Printf("[%s][%d] Records read", filePath, offset)
        for _, key := range results {
            if _, ok := indexes[key]; !ok {
                increment++
                result = append(result, key)
                index :=  fmt.Sprintf("%d", increment)
                indexes[key] = index
            }
        }
    }

    return indexes, nil
}
