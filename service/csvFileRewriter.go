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
    xxhash "bitbucket.org/StephaneBunel/xxhash-go"
    "github.com/helyx-io/gtfs-importer/data"
    "strings"
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
    serviceIndexes, _ := cfr.getIndexes(schema, "trips.txt", 1)
    tripIndexes, _ := cfr.getIndexes(schema, "trips.txt", 2)
    stopIndexes, _ := cfr.getIndexes(schema, "stops.txt", 0)
    routeIndexes, _ := cfr.getIndexes(schema, "routes.txt", 0)

    cfr.writeIndexes(schema, "routes.indexes.txt", outFolderName, routeIndexes);
    cfr.writeIndexes(schema, "trip.indexes.txt", outFolderName, tripIndexes);
    cfr.writeIndexes(schema, "stop.indexes.txt", outFolderName, stopIndexes);

    folderFilename := cfr.tmpDir + "/" + schema
    outFolderFilename := path.Join(folderFilename, outFolderName)
    outTmpFolderFilename := path.Join(outFolderFilename, "tmp")

    if err := os.MkdirAll(outFolderFilename, 0755) ; err != nil {
        utils.FailOnError(err, fmt.Sprintf("Unable to create directory: '%s'!", outFolderFilename))
    }

    if err := os.MkdirAll(outTmpFolderFilename, 0755) ; err != nil {
        utils.FailOnError(err, fmt.Sprintf("Unable to create directory: '%s'!", outTmpFolderFilename))
    }

    stopIndexes, _ = cfr.dedupCsvFile(schema, path.Join(folderFilename, "stops.txt"), path.Join(outTmpFolderFilename, "stops.txt"), stopIndexes, 0)
    cfr.writeIndexes(schema, "stop.dedup.indexes.txt", outFolderName, stopIndexes);

    cfr.fixRouteColors(schema, path.Join(folderFilename, "routes.txt"), path.Join(outTmpFolderFilename, "routes.txt"))

    cfr.rewriteCsvFile(schema, path.Join(outTmpFolderFilename, "stops.txt"), path.Join(outFolderFilename, "stops.txt"), map[int]map[string]string{ 0: stopIndexes })
    cfr.rewriteCsvFile(schema, path.Join(folderFilename, "stop_times.txt"), path.Join(outFolderFilename, "stop_times.txt"), map[int]map[string]string{ 0: tripIndexes, 3: stopIndexes })
    cfr.rewriteCsvFile(schema, path.Join(outTmpFolderFilename, "routes.txt"), path.Join(outFolderFilename, "routes.txt"), map[int]map[string]string{ 0: routeIndexes, 1: agencyIndexes })
    cfr.rewriteCsvFile(schema, path.Join(folderFilename, "agency.txt"), path.Join(outFolderFilename, "agency.txt"), map[int]map[string]string{})
    cfr.rewriteCsvFile(schema, path.Join(folderFilename, "trips.txt"), path.Join(outFolderFilename, "trips.txt"), map[int]map[string]string{ 0: routeIndexes, 1: serviceIndexes, 2: tripIndexes })
    cfr.rewriteCsvFile(schema, path.Join(folderFilename, "calendar.txt"), path.Join(outFolderFilename, "calendar.txt"), map[int]map[string]string{ 0: serviceIndexes })
    cfr.rewriteCsvFile(schema, path.Join(folderFilename, "calendar_dates.txt"), path.Join(outFolderFilename, "calendar_dates.txt"), map[int]map[string]string{ 0: serviceIndexes })
    cfr.rewriteCsvFile(schema, path.Join(folderFilename, "transfers.txt"), path.Join(outFolderFilename, "transfers.txt"), map[int]map[string]string{})


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


func (cfr *CsvFileRewriter) rewriteCsvFile(schema, filePath, outFilePath string, indexes map[int]map[string]string) error {

    gtfsFile := models.GTFSFile{filePath}

    outFile, err := os.Create(outFilePath)
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

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            records, _ := models.ParseCsvAsStringArrays(lines)

            for _, record := range *records {
                for i, _ := range indexes {
                    record[i] = indexes[i][record[i]]
                }
            }

            resultChan <- *records
        }

        close(resultChan)

    }()

    for results := range resultChan {
        writer.WriteAll(results)
    }

    writer.Flush()
    err = outFile.Close()

    return err
}


func (cfr *CsvFileRewriter) dedupCsvFile(schema, filePath, outFilePath string, indexes map[string]string, pkeyField int) (map[string]string, error) {
    gtfsFile := models.GTFSFile{filePath}

    outFile, err := os.Create(outFilePath)
    if err != nil {
        log.Printf("Error: %v", err.Error())
        return nil, err
    }

    log.Printf("Writing to file: '%v'", outFile.Name())

    writer := csv.NewWriter(outFile)

    headers, err := utils.ReadCsvFileHeader(gtfsFile.Filename, ",")
    log.Printf("headers: '%v'", headers)

    if err != nil {
        log.Printf("Error: '%v'", err.Error())
        return nil, err
    }

    writer.Write(headers)

    resultChan := make(chan [][]string)

    recordHashes := make(map[uint32]string)

    go func() {

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            records, _ := models.ParseCsvAsStringArrays(lines)

            filteredRecords := make([][]string, 0, len(*records))

            for _, record := range *records {
                if pkeyField >= 0 {
                    digest := xxhash.GoNew32()
                    for i, recordField := range record {
                        if pkeyField != i {
//                            log.Printf("recordField[%d][%d]:%s", r, i, recordField)
                            digest.Write([]byte(recordField))
                        }
                    }

                    hash := digest.Sum32()
//                    log.Printf("recordField[%d]:%d", r, hash)
                    if pkey, ok := recordHashes[hash]; !ok {
                        recordHashes[hash] = record[pkeyField]

                        filteredRecords = append(filteredRecords, record)
                    } else {
                        indexes[record[pkeyField]] = indexes[pkey]
                        continue;
                    }
                }
            }

            resultChan <- filteredRecords
        }

        close(resultChan)

    }()

    for results := range resultChan {
        writer.WriteAll(results)
    }

    writer.Flush()
    err = outFile.Close()
    if err != nil {
        return nil, err
    }

    return indexes, nil
}


func (cfr *CsvFileRewriter) loadColorsByRoute(agencyKey string) (map[string]string, error) {

    asset, err := data.Asset(fmt.Sprintf("resources/gtfs/%s/route-colors.csv", agencyKey))
    if err != nil {
        log.Printf("[loadColorsByRoute] err: %v", asset, err)
        return nil, err
    }

    records, _ := models.ParseCsvAsStringArrays(asset)

    colorsByRouteMap := make(map[string]string)

    for _, record := range *records {
        colorsByRouteMap[record[1]] = strings.ToUpper(record[6])
    }

    log.Printf("[loadColorsByRoute] Colors loaded: %v", colorsByRouteMap)

    return colorsByRouteMap, nil
}


func (cfr *CsvFileRewriter) fixRouteColors(agencyKey, filePath, outFilePath string) error {
    log.Printf("[fixRouteColors] Fixing route colors ...")
    colorsByRoute, err := cfr.loadColorsByRoute(agencyKey)

    gtfsFile := models.GTFSFile{filePath}

    outFile, err := os.Create(outFilePath)
    if err != nil {
        log.Printf("[fixRouteColors] Error: %v", err.Error())
        return err
    }

    log.Printf("[fixRouteColors] Writing to file: '%v'", outFile.Name())

    writer := csv.NewWriter(outFile)

    headers, err := utils.ReadCsvFileHeader(gtfsFile.Filename, ",")
    log.Printf("[fixRouteColors] headers: '%v'", headers)

    if err != nil {
        log.Printf("[fixRouteColors] Error: '%v'", err.Error())
        return err
    }

    writer.Write(headers)

    resultChan := make(chan [][]string)

    go func() {

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            records, _ := models.ParseCsvAsStringArrays(lines)

            if colorsByRoute != nil {
                for _, record := range *records {

                    if routeColor, ok := colorsByRoute[record[2]]; !ok {
                        record[7] = "000000"
                    } else {
                        record[7] = routeColor
                    }

                    record[8] = "FFFFFF"
                }
            }

            resultChan <- *records
        }

        close(resultChan)

    }()

    for results := range resultChan {
        writer.WriteAll(results)
    }

    writer.Flush()
    err = outFile.Close()
    if err != nil {
        return err
    }

    return nil
}



func (cfr *CsvFileRewriter) getIndexes(schema, filename string, index int) (map[string]string, error) {

    filePath := path.Join(cfr.tmpDir, schema, filename)

    gtfsFile := models.GTFSFile{filePath}

    resultChan := make(chan []string)

    go func() {

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            records, _ := models.ParseCsvAsStringArrays(lines)

            keys := []string{}
            for _, record := range *records {
                keys = append(keys, record[index])
            }

            resultChan <- keys
        }

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
