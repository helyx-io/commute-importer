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
    "errors"
    "github.com/fatih/stopwatch"
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

func (cfr *CsvFileRewriter) RewriteCsvFiles(schema, outFolderName string) (map[string]map[string]int, error) {

    agencyIndexes, _ := cfr.getIndexes(schema, "agency.txt", 0)
    servcfreIndexes, _ := cfr.getIndexes(schema, "trips.txt", 1)
    tripIndexes, _ := cfr.getIndexes(schema, "trips.txt", 2)
    stopIndexes, _ := cfr.getIndexes(schema, "stops.txt", 0)
    routeIndexes, _ := cfr.getIndexes(schema, "routes.txt", 0)

    cfr.writeIndexes(schema, "routes.indexes.txt", outFolderName, routeIndexes);
    cfr.writeIndexes(schema, "trip.indexes.txt", outFolderName, tripIndexes);
    cfr.writeIndexes(schema, "stop.indexes.txt", outFolderName, stopIndexes);

    folderFilename := cfr.tmpDir + "/" + schema
    outFolderFilename := path.Join(folderFilename, outFolderName)

    if os.MkdirAll(outFolderFilename, 0755) != nil {
        panic("Unable to create directory for tagfile!")
    }


    indexesByFiles := make(map[string]map[int]map[string]string)
    indexesByFiles["stops.txt"] = map[int]map[string]string{ 0: stopIndexes }
    indexesByFiles["stop_times.txt"] = map[int]map[string]string{ 0: tripIndexes, 3: stopIndexes }
    indexesByFiles["routes.txt"] = map[int]map[string]string{ 0: routeIndexes, 1: agencyIndexes }
    indexesByFiles["agency.txt"] = map[int]map[string]string{ 0: stopIndexes }
    indexesByFiles["trips.txt"] = map[int]map[string]string{ 0: routeIndexes, 1: servcfreIndexes, 2: tripIndexes }
    indexesByFiles["calendar.txt"] = map[int]map[string]string{ 0: servcfreIndexes }
    indexesByFiles["calendar_dates.txt"] = map[int]map[string]string{ 0: servcfreIndexes }
    indexesByFiles["transfers.txt"] = map[int]map[string]string{}

    for filename, indexesByFile := range indexesByFiles {
        cfr.rewriteCsvFile(schema, filename, outFolderName, indexesByFile)
    }


    columLengthsMap := make(map[string]map[string]int)

    files := map[string]string {
        "agency.txt": "agencies",
        "calendar_dates.txt": "calendar_dates",
        "calendar.txt": "calendars",
        "routes.txt": "routes",
        "stop_times.txt": "stop_times",
        "stops.txt": "stops",
        "transfers.txt": "transfers",
        "trips.txt": "trips",
    }

    for file, tableName := range files {
        columLengths, err := cfr.getColumnsLength(schema, tableName, outFolderName, file)

        log.Printf("[%s][%s] columLengths: %v", tableName, file, columLengths)

        if err != nil {
            log.Printf("[%s][%s] Error: %v", tableName, file, err)
        } else {
            columLengthsMap[tableName] = columLengths
        }
    }

    log.Printf("columLengthsMap: %v", columLengthsMap)
    return columLengthsMap, nil
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

func (cfr *CsvFileRewriter) getColumnsLength(schema, tableName, outFolderName, filename string) (map[string]int, error) {

    sw := stopwatch.Start(0)

    filePath := path.Join(cfr.tmpDir, schema, outFolderName, filename)

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        log.Printf("File does not exists: %v", filePath)
        return nil, err
    }

    log.Printf("Processing file: %v", filePath)

    gtfsFile := models.GTFSFile{filePath}

    resultChan := make(chan []int)

    go func() {
        for lines := range gtfsFile.LinesIterator(1024 * 1024) {

            record, err := models.ParseCsvAsIntArrays(lines)
            if err == nil {
                resultChan <- record
            }
        }

        close(resultChan)
    }()

    records := make([][]int, 0)
    for record := range resultChan {
        records = append(records, record)
    }

    log.Printf("Records: %v", records)

    if len(records) > 0 {
        lengthRecord := make([]int, len(records[0]))
        for _, record := range records {
            for i, field := range record {
                if field > lengthRecord[i] {
                    lengthRecord[i] = field
                }
            }
        }

        log.Printf("[%s] FieldMaxLengths: %d", filename, lengthRecord)
        log.Printf("[%s] ElapsedTime: %v", filename, sw.ElapsedTime())

        fields, err := cfr.readCsvFileHeaders(schema, outFolderName, filename)

        for _, field := range fields {
            log.Printf("Header: %v", field)
        }

        log.Printf("filename: %s, headers: %v", filename, fields)

        if err != nil {
            log.Printf("Error: %v", err)
            return nil, err
        } else {
            lengthByHeader := make(map[string]int)
            log.Printf("headers: %v - lengthRecord: %v", fields, lengthRecord)
            for i, field := range fields {
                lengthByHeader[field] = lengthRecord[i]
            }
            return lengthByHeader, nil
        }

    } else {
        return nil, errors.New("No lines to count field lengths")
    }
}

func (cfr *CsvFileRewriter) readCsvFileHeaders(schema, outFolderName, filename string) ([]string, error) {
    filePath := path.Join(cfr.tmpDir, schema, outFolderName, filename)

    log.Printf("Reading csv file headers for file: %v", filePath)

    gtfsFile := models.GTFSFile{filePath}

    return utils.ReadCsvFileHeader(gtfsFile.Filename, ",")
}
