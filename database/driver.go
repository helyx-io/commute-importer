package database

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "github.com/helyx-io/gtfs-importer/config"
    "github.com/jinzhu/gorm"
    "fmt"
    "log"
    "github.com/helyx-io/gtfs-importer/data"
    "github.com/helyx-io/gtfs-importer/utils"
    "regexp"
    "database/sql"
    "text/template"
    "bytes"
    "reflect"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type Driver struct {
    DB *gorm.DB
    ConnectInfos *config.DBConnectInfos
}

func NewDriver(DB *gorm.DB, ConnectInfos *config.DBConnectInfos) *Driver {
    return &Driver{DB, ConnectInfos}
}

func (d *Driver) Close() {
    if d.DB != nil {
        defer d.DB.Close()
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func (d *Driver) Open() (*sql.DB, error) {
    return sql.Open(d.ConnectInfos.Dialect, d.ConnectInfos.URL)
}

func (d *Driver) Raw(query string) *gorm.DB {
    return d.DB.Raw(query)
}

func (d *Driver) ExecQuery(query string, params ...string) error {
    log.Printf("Exec Stmt: '%s' - Params: %v", query, params)

    return d.DB.Exec(query, params).Error
}

func (d *Driver) Exec(filename string, params ...string) error {
    filePath := fmt.Sprintf("resources/ddl/%s/%s.sql", d.ConnectInfos.Dialect, filename)
    log.Printf("Executing query from file path: '%s'", filePath)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for exec", filePath))
    execStmt := fmt.Sprintf(string(dml))
    log.Printf("Exec Stmt: '%s' - Params: %v", execStmt, params)

    return d.DB.Exec(execStmt, params).Error
}

func (d *Driver) DropTable(schema, tableName string) error {
    log.Printf("d: %v", d)
    log.Printf("d.ConnectInfos: %v", d.ConnectInfos)
    filePath := fmt.Sprintf("resources/ddl/%s/drop-table.sql", d.ConnectInfos.Dialect)
    log.Printf("Dropping table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for drop of table '%s.%s'", filePath, schema, tableName))
    dropStmt := fmt.Sprintf(string(dml), schema, tableName)
    log.Printf("Drop statement: %s", dropStmt)

    return d.DB.Exec(dropStmt).Error
}

func (d *Driver) CreateTable(schema, tableName string, params map[string]interface{}, dropIfExists bool) error {

    if dropIfExists {
        d.DropTable(schema, tableName)
    }

    filePath := fmt.Sprintf("resources/ddl/%s/create-table-%s.sql", d.ConnectInfos.Dialect, tableName)
    log.Printf("Creating table with name: '%s.%s' with query from file path: '%s' - Params: %v", schema, tableName, filePath, params)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for create of table '%s.%s'", filePath, schema, tableName))


    log.Printf("query name: %s", fmt.Sprintf("%s_%s", schema, tableName))
    log.Printf("query content: %s", fmt.Sprintf(string(dml), schema))
    log.Printf("query params: %v", params)

    funcMap := template.FuncMap {
        "length": func(value interface{}) interface{} {

            log.Printf("=== Value: %v", value)
            log.Printf("=== Type Of: %v", reflect.TypeOf( value))

            if value == nil {
                return 1
            } else if reflect.ValueOf(value).Kind() == reflect.Int && value.(int) == 0 {
                return 1
            } else {
                return value
            }
        },
    }

    t, err := template.New(fmt.Sprintf("%s_%s", schema, tableName)).Funcs(funcMap).Parse(fmt.Sprintf(string(dml), schema))

    if err != nil {
        log.Printf("error: %v", err)
    }

    var doc bytes.Buffer
    t.Execute(&doc, params)
    createStmt :=  doc.String()

    log.Printf("Create statement: %s", createStmt)
    return d.DB.Exec(createStmt).Error
}

func (d *Driver) CreateIndex(schema, tableName, indexName string) error {
    return d.CreateIndexWithScript("create-index", schema, tableName, indexName)

}

func (d *Driver) CreateSpatialIndex(schema, tableName, indexName string) error {
    return d.CreateIndexWithScript("create-spatial-index", schema, tableName, indexName)
}


func (d *Driver) CreateIndexWithScript(script, schema, tableName, indexName string) error {
    filePath := fmt.Sprintf("resources/ddl/%s/%s.sql", d.ConnectInfos.Dialect, script)
    log.Printf("Creating index with for table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for create index of table '%s.%s'", filePath, schema, tableName))

    var createIndexStmt string
    if d.ConnectInfos.Dialect == "postgres" {
        createIndexStmt = fmt.Sprintf(string(dml), tableName, indexName, schema, tableName, indexName)
    } else {
        createIndexStmt = fmt.Sprintf(string(dml), schema, tableName, tableName, indexName, indexName)
    }
    log.Printf("Create statement for index: %s", createIndexStmt)
    return  d.DB.Exec(createIndexStmt).Error
}


func (d *Driver) PopulateTable(schema string, tableName string) {

    filePath := fmt.Sprintf("resources/ddl/%s/insert-%s.sql", d.ConnectInfos.Dialect, tableName)

    log.Printf("Inserting data into table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

    ddl, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table '%s.%s'", filePath, schema, tableName))

    re := regexp.MustCompile("%s")
    insertStmt := re.ReplaceAllString(string(ddl), schema)

    log.Printf("Insert statement: %s", insertStmt)
    err = d.DB.Exec(insertStmt).Error
    utils.FailOnError(err, fmt.Sprintf("Could not insert into '%s' table", tableName))
}
