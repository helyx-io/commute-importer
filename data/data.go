package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _resources_ddl_mysql_create_index_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x50\x2d\x4e\xd0\x03\x11\x0a\x8e\x2e\x2e\x0a\x9e\x7e\x2e\xae\x11\x20\xb1\x78\x20\xca\x4c\xa9\x48\x50\xd0\x80\xc8\x05\x3b\x6b\x5a\x03\x02\x00\x00\xff\xff\x5f\x0f\x2c\x41\x37\x00\x00\x00")

func resources_ddl_mysql_create_index_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_index_sql,
		"resources/ddl/mysql/create-index.sql",
	)
}

func resources_ddl_mysql_create_index_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_index_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-index.sql", size: 55, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_schema_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x76\xf6\x70\xf5\x75\x54\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x50\x2d\x4e\x50\x70\x71\x75\x73\x0c\xf5\x09\x51\x70\xf6\x70\x0c\x72\x74\x0e\x71\x0d\x52\x08\x76\x0d\x51\x28\x2d\x49\xb3\xc8\x4d\x32\xb1\x06\x04\x00\x00\xff\xff\x1e\x51\x0a\x2a\x3f\x00\x00\x00")

func resources_ddl_mysql_create_schema_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_schema_sql,
		"resources/ddl/mysql/create-schema.sql",
	)
}

func resources_ddl_mysql_create_schema_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_schema_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-schema.sql", size: 63, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_spatial_index_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x50\x2d\x4e\xd0\x03\x11\x0a\x8e\x2e\x2e\x0a\xc1\x01\x8e\x21\x9e\x8e\x3e\x0a\x9e\x7e\x2e\xae\x11\x20\xb9\x78\x20\xca\x4c\xa9\x48\x50\xd0\x00\xa9\xd1\xb4\xe6\x02\x04\x00\x00\xff\xff\x8c\xc6\x85\xbd\x3c\x00\x00\x00")

func resources_ddl_mysql_create_spatial_index_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_spatial_index_sql,
		"resources/ddl/mysql/create-spatial-index.sql",
	)
}

func resources_ddl_mysql_create_spatial_index_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_spatial_index_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-spatial-index.sql", size: 60, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_agencies_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xd0\xc1\x4a\xc4\x30\x10\x80\xe1\x7b\x9f\x62\x2e\x42\x02\x22\x2c\xac\x20\xc8\x1e\xb2\xbb\xa3\x16\x63\x94\x98\x3d\xec\xc9\xc4\x9a\xd6\x40\x33\x81\x36\x15\xf5\xe9\x55\x04\x29\x48\x21\xc7\x61\xbe\xf9\x0f\xb3\xd3\x28\x0c\x82\x11\x5b\x89\x60\x4f\x46\x7b\x66\x5d\xe7\xa9\x09\x7e\xb4\xc0\x2a\x80\xdf\xf1\xe3\x29\xbc\x58\x08\x94\xd9\x6a\xc5\x41\xdd\x1b\x50\x07\x29\x4f\x67\x7b\x72\xd1\x5b\x78\x73\x43\xf3\xea\x06\xb6\x3e\xe7\xb0\xc7\x2b\x71\x90\xff\xe5\x34\xf4\x65\x30\x87\xe8\x3f\x13\x15\x66\x7b\x47\x5d\x99\x8c\x81\xbe\x75\xb6\xd0\xf6\xc9\xe5\x65\xe6\xde\x8b\xd8\x4f\x2d\x51\x51\x6d\x99\x3d\xe8\xfa\x4e\xe8\x23\xdc\xe2\x11\xd8\xec\xe7\xbc\xe2\x80\xea\xba\x56\xb8\xa9\x89\xd2\x7e\xfb\x77\xb9\xbb\x11\xfa\x11\xcd\x66\xca\xed\x45\x7c\x5e\x5f\x56\x5f\x01\x00\x00\xff\xff\x5c\xc4\x45\xf8\xcb\x01\x00\x00")

func resources_ddl_mysql_create_table_agencies_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_agencies_sql,
		"resources/ddl/mysql/create-table-agencies.sql",
	)
}

func resources_ddl_mysql_create_table_agencies_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_agencies_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-agencies.sql", size: 459, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_calendar_dates_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x48\x50\x2d\x4e\xd0\x4b\x48\x4e\xcc\x49\xcd\x4b\x49\x2c\x8a\x4f\x49\x2c\x49\x2d\x4e\x50\xd0\xe0\x52\x50\x48\x28\x4e\x2d\x2a\xcb\x4c\x4e\x8d\xcf\x4c\x49\x50\xc8\xcc\x2b\xd1\x30\x34\xd4\x54\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\xd1\x01\x29\x00\x29\x4e\x50\x00\x91\xa8\xe2\xa9\x15\xc9\xa9\x05\x25\x99\xf9\x79\xf1\x25\x95\x05\xa9\x08\xcd\x2e\xae\x6e\x8e\xa1\x3e\x08\x85\x01\x41\x9e\xbe\x8e\x41\x91\x0a\xde\xae\x91\x0a\x1a\xc8\xd6\xe9\x40\xcd\xd6\xe4\xd2\x54\x70\xf5\x73\xf7\xf4\x73\xb5\xf5\xcc\xcb\xcb\x77\x71\x82\x1b\xe1\xec\xe1\x18\x14\xec\x1a\x62\x5b\x5a\x92\x66\x91\x9b\x64\x62\xcd\x05\x08\x00\x00\xff\xff\x95\xc2\x9e\x75\xd5\x00\x00\x00")

func resources_ddl_mysql_create_table_calendar_dates_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_calendar_dates_sql,
		"resources/ddl/mysql/create-table-calendar_dates.sql",
	)
}

func resources_ddl_mysql_create_table_calendar_dates_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_calendar_dates_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-calendar_dates.sql", size: 213, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_calendars_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xd0\x31\x4b\x43\x31\x10\xc0\xf1\xfd\x7d\x8a\x5b\x84\x17\x10\xa1\xe0\x20\x48\x87\xb4\x3d\xf5\x61\x8c\x12\xd3\xa1\x53\x2f\x36\x29\x06\x6c\x0a\xc9\x45\xe9\xb7\xb7\x75\x90\x07\x82\xe9\x72\xcb\xfd\x6e\xb8\xff\xdc\xa0\xb4\x08\x56\xce\x14\x02\x5d\x14\xba\xa2\x8d\xfb\x08\xc9\xbb\x5c\x08\xfa\x0e\x80\x4a\xc8\x9f\x71\x13\xd6\xd1\x13\xc4\xc4\xfd\x64\x22\x40\x3f\x5b\xd0\x4b\xa5\x2e\x4f\x60\xb7\x3f\xf2\x03\x01\xc7\x74\xf8\x01\x02\x16\x78\x27\x97\x6a\x64\xb8\x86\xd2\x44\x5f\xc1\xa7\x33\x18\xbf\xd7\xdc\x56\xdb\x1c\x9b\xa6\x38\xae\xb9\xad\x6a\xfb\xbf\xc2\x2e\xf3\xda\x3b\x0e\x04\xa7\xf9\x57\x1c\x9b\xfe\xb7\x7f\x31\xc3\x93\x34\x2b\x78\xc4\x15\xf4\xe3\xe6\xa2\x13\x80\xfa\x7e\xd0\x38\x1d\x52\xda\x2f\x66\xbf\xa7\xf3\x07\x69\x5e\xd1\x4e\x2b\x6f\x6f\x76\x6f\xd7\xb7\xdd\x77\x00\x00\x00\xff\xff\xf6\x18\xf0\xcb\xcd\x01\x00\x00")

func resources_ddl_mysql_create_table_calendars_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_calendars_sql,
		"resources/ddl/mysql/create-table-calendars.sql",
	)
}

func resources_ddl_mysql_create_table_calendars_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_calendars_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-calendars.sql", size: 461, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_line_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x48\x50\x2d\x4e\xd0\x4b\xc8\xc9\xcc\x4b\x8d\x2f\x2e\xc9\x2f\x28\x4e\x50\xd0\xe0\x52\x50\x80\x08\x64\xa6\x24\x28\x64\xe6\x95\x68\x18\x1a\x6a\x2a\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\xe8\x80\x64\x41\x2a\x09\xc8\x26\xe7\xa7\xa4\x26\x28\x94\x25\x16\x25\x67\x24\x16\x69\x98\x98\xa2\xaa\x09\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xd0\x80\x5b\xa6\x03\x37\x58\x93\x4b\x53\xc1\xd5\xcf\xdd\xd3\xcf\xd5\xd6\x33\x2f\x2f\xdf\xc5\x49\xc1\xc5\xd5\xcd\x31\xd4\x27\x44\xc1\xd9\xc3\x31\x28\xd8\x35\xc4\xb6\xb4\x24\xcd\x22\x37\xc9\xc4\x9a\x0b\x10\x00\x00\xff\xff\x40\x4c\x0d\xa7\xce\x00\x00\x00")

func resources_ddl_mysql_create_table_line_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_line_stops_sql,
		"resources/ddl/mysql/create-table-line_stops.sql",
	)
}

func resources_ddl_mysql_create_table_line_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_line_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-line_stops.sql", size: 206, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8c\xc1\x0b\x82\x30\x14\x87\xef\xfe\x15\xef\x12\x6c\x10\x81\x60\x10\x84\x87\xa9\xaf\x1a\xcd\x19\x6b\x1e\x3c\x39\x2b\x23\x21\x17\xa8\x45\x7f\x7e\x2a\xa4\xb7\xc7\xef\x7d\xdf\x17\x2a\x64\x1a\x41\xb3\x40\x20\x98\x45\x6b\x56\xe6\x59\xd9\xb2\x35\x40\x1c\x80\xf1\xce\xab\x9b\x81\xca\x76\xc4\x75\x29\xc8\x44\x83\x4c\x85\x00\x96\xea\x24\xe7\xb2\xf7\x63\x94\x7a\x39\xc1\xb6\xa8\x4b\x03\x9f\xa2\xb9\x3e\x8a\x86\x78\xeb\x59\x19\x98\x93\xe2\x31\x53\x19\x1c\x31\x03\x32\xd5\xe9\xf0\x1a\xa6\x39\xd1\xcf\x5f\xf3\x47\xc6\x26\x75\x28\xa0\xdc\x73\x89\x3e\xb7\xf6\x15\x05\x10\xe1\x8e\xa5\x42\x43\x78\x60\xea\x8c\xda\x7f\x77\xf7\x4d\x7d\xf1\xb6\xce\x2f\x00\x00\xff\xff\xe3\xbe\x63\xc3\xd5\x00\x00\x00")

func resources_ddl_mysql_create_table_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_lines_sql,
		"resources/ddl/mysql/create-table-lines.sql",
	)
}

func resources_ddl_mysql_create_table_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-lines.sql", size: 213, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_route_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x48\x50\x2d\x4e\xd0\x4b\x28\xca\x2f\x2d\x49\x8d\x2f\x2e\xc9\x2f\x28\x4e\x50\xd0\xe0\x52\x50\x80\x8a\x64\xa6\x24\x28\x64\xe6\x95\x68\x18\x1a\x6a\x2a\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\xe8\x80\xa4\x41\x4a\x09\xc8\x26\xe7\xa7\xa4\x26\x28\x94\x25\x16\x25\x67\x24\x16\x69\x98\x98\xa2\xaa\x09\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xd0\x40\xd8\xa6\x03\x37\x59\x93\x4b\x53\xc1\xd5\xcf\xdd\xd3\xcf\xd5\xd6\x33\x2f\x2f\xdf\xc5\x49\xc1\xc5\xd5\xcd\x31\xd4\x27\x44\xc1\xd9\xc3\x31\x28\xd8\x35\xc4\xb6\xb4\x24\xcd\x22\x37\xc9\xc4\x9a\x0b\x10\x00\x00\xff\xff\xba\x29\x70\x22\xd1\x00\x00\x00")

func resources_ddl_mysql_create_table_route_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_route_stops_sql,
		"resources/ddl/mysql/create-table-route_stops.sql",
	)
}

func resources_ddl_mysql_create_table_route_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_route_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-route_stops.sql", size: 209, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_routes_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xd0\x41\x4b\x03\x31\x10\x05\xe0\xfb\xfe\x8a\xb9\x08\x09\x88\xb0\xb2\x4a\x41\x7a\x48\xdb\x51\x17\x63\x94\x98\x1e\x7a\x6a\x62\x1a\xdb\xc2\x36\x91\x24\x2b\xf6\xdf\x8b\x2c\xb8\xa0\x04\x7b\x0c\xf3\x3d\xde\x23\x73\x89\x4c\x21\x28\x36\xe3\x08\xfa\x2c\xe9\x0b\x1d\x43\x9f\x5d\xd2\x40\x2a\x80\xe1\xb1\xde\x6f\x34\xec\x7d\x26\x75\x4d\x41\x3c\x29\x10\x4b\xce\xcf\xbf\xcf\x66\xeb\xbc\x3d\x96\xef\x43\x3c\xed\x42\xcc\x6b\x6f\x0e\x4e\xc3\x87\x89\x76\x67\x22\x69\xae\x28\x2c\xf0\x96\x2d\xf9\x1f\xde\x05\xbf\xfd\xa5\xeb\xcb\x49\x91\x6f\x5c\xb2\xa3\xbc\x6e\x8a\x30\x1f\xdf\xdd\xb8\xb3\x80\xfa\xd8\x9d\x34\xd2\x86\x2e\x44\x0d\x43\x67\xb9\xd2\x7d\xe6\xff\xe8\xb3\x6c\x1f\x99\x5c\xc1\x03\xae\x80\x8c\x1f\x4e\x2b\x0a\x28\xee\x5a\x81\xd3\xd6\xfb\xb0\x98\xfd\x04\xe7\xf7\x4c\xbe\xa0\x9a\xf6\xf9\x6d\x72\x78\x6d\x6e\xaa\xaf\x00\x00\x00\xff\xff\x07\x8f\x8b\x27\xc5\x01\x00\x00")

func resources_ddl_mysql_create_table_routes_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_routes_sql,
		"resources/ddl/mysql/create-table-routes.sql",
	)
}

func resources_ddl_mysql_create_table_routes_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_routes_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-routes.sql", size: 453, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_station_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x48\x50\x2d\x4e\xd0\x4b\x28\x2e\x49\x2c\xc9\xcc\xcf\x8b\xcf\xc9\xcc\x4b\x2d\x4e\x50\xd0\xe0\x52\x50\x80\x8b\x65\xa6\x24\x28\x64\xe6\x95\x68\x18\x1a\x6a\x2a\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\xe8\x80\x14\x80\x14\xe3\x94\x0d\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xd0\x40\x36\x4b\x07\xa1\x51\x93\x4b\x53\xc1\xd5\xcf\xdd\xd3\xcf\xd5\xd6\x33\x2f\x2f\xdf\xc5\x49\xc1\xc5\xd5\xcd\x31\xd4\x27\x44\xc1\xd9\xc3\x31\x28\xd8\x35\xc4\xb6\xb4\x24\xcd\x22\x37\xc9\xc4\x9a\x0b\x10\x00\x00\xff\xff\x7e\x55\xf7\xb0\xb4\x00\x00\x00")

func resources_ddl_mysql_create_table_station_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_station_lines_sql,
		"resources/ddl/mysql/create-table-station_lines.sql",
	)
}

func resources_ddl_mysql_create_table_station_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_station_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-station_lines.sql", size: 180, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_station_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x48\x50\x2d\x4e\xd0\x4b\x28\x2e\x49\x2c\xc9\xcc\xcf\x8b\x2f\x2e\xc9\x2f\x28\x4e\x50\xd0\xe0\x52\x50\x80\x8b\x65\xa6\x24\x28\x64\xe6\x95\x68\x18\x1a\x6a\x2a\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\xe8\x40\x14\xe4\x17\xe0\x94\x0d\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xd0\x40\x36\x4b\x07\xae\x4f\x93\x4b\x53\xc1\xd5\xcf\xdd\xd3\xcf\xd5\xd6\x33\x2f\x2f\xdf\xc5\x49\xc1\xc5\xd5\xcd\x31\xd4\x27\x44\xc1\xd9\xc3\x31\x28\xd8\x35\xc4\xb6\xb4\x24\xcd\x22\x37\xc9\xc4\x9a\x0b\x10\x00\x00\xff\xff\x0a\x99\xa0\xa0\xb3\x00\x00\x00")

func resources_ddl_mysql_create_table_station_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_station_stops_sql,
		"resources/ddl/mysql/create-table-station_stops.sql",
	)
}

func resources_ddl_mysql_create_table_station_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_station_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-station_stops.sql", size: 179, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_stations_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xcf\x41\x4b\xc4\x30\x10\x05\xe0\x7b\x7f\xc5\xbb\x08\x29\x88\xb0\xb0\x88\x20\x7b\xc8\xee\x8e\x1a\xcc\xa6\x12\xd3\x43\x4f\x9d\xa8\xad\x16\xda\x04\xda\x28\xf8\xef\xed\x45\x6d\x41\xf0\x32\x30\xf0\xcd\x1b\xde\xc1\x92\x74\x04\x27\xf7\x9a\xc0\x67\x13\x5f\xf0\x94\x7c\xea\x62\x98\x18\x22\x03\xbe\xd7\xba\x7b\x61\x74\x21\x89\xcd\x26\x87\x29\x1c\x4c\xa9\x35\x64\xe9\x8a\x5a\x99\x39\xe5\x44\xc6\x9d\x2f\x7d\xf0\x43\xc3\xf8\xf0\xe3\xf3\x9b\x1f\xc5\xe5\xf6\xf7\x6a\xc5\x7a\x9f\x18\x6d\x1f\x7d\xc2\x91\x6e\x64\xa9\xff\x32\x31\xfc\x6b\x5e\x9b\xc8\x98\xc7\xd0\xa4\xf1\x73\xf5\xea\xc1\xaa\x93\xb4\x15\xee\xa9\x82\x58\xd6\xc9\xb3\x1c\x64\x6e\x95\xa1\x9d\x0a\x21\x1e\xf7\x3f\xe9\x87\x3b\x69\x1f\xc9\xed\xde\x53\x7b\x35\x3c\x6d\xaf\xb3\xaf\x00\x00\x00\xff\xff\x36\xd3\x78\x66\x27\x01\x00\x00")

func resources_ddl_mysql_create_table_stations_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_stations_sql,
		"resources/ddl/mysql/create-table-stations.sql",
	)
}

func resources_ddl_mysql_create_table_stations_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_stations_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-stations.sql", size: 295, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_stop_times_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x8f\x3f\x4b\xc5\x30\x14\x47\xf7\x7e\x8a\xbb\x08\x7d\x20\xc2\x03\x87\x07\xf2\x86\xbc\x36\x6a\x21\x44\xa8\xe9\xdc\x1b\x9b\xd4\x06\x6d\x1a\xf3\x47\xf0\xdb\xdb\x76\x50\x87\xda\x25\xcb\x39\x39\x3f\x6e\x51\x53\x22\x28\x08\x72\x61\x14\xf0\x2a\xe0\x0d\x86\x38\xb9\x36\x9a\x51\x07\x84\x3c\x03\xc0\xe8\x8d\x6b\x8d\x42\x30\x36\xe6\xc7\xe3\x01\xf8\x93\x00\xde\x30\x76\xbd\x50\xe9\xbd\xf9\x94\xef\xeb\x0f\x84\xe5\x85\x92\xde\x93\x86\xfd\x71\x94\x76\xd2\xc7\xe4\xf5\xae\xb5\x0e\xff\xbb\xb3\xd2\xa0\x3f\x92\xb6\x9d\xfe\x75\xb6\x2b\x83\x96\xaa\x0d\xe6\xd5\x22\x74\x83\xf4\xf9\x69\x43\x74\xa6\x7b\x4b\xf3\xa5\x5f\x6e\x2f\xa7\xfc\x9c\x9b\xfa\x7e\xc7\xcb\x0e\x40\xf9\x43\xc5\xe9\xb9\xb2\x76\x2a\x2f\x3f\xb0\x78\x24\xf5\x33\x15\xe7\x14\xfb\xd3\xf8\x72\x7b\x97\x7d\x07\x00\x00\xff\xff\xcc\x13\xd9\x40\x6e\x01\x00\x00")

func resources_ddl_mysql_create_table_stop_times_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_stop_times_sql,
		"resources/ddl/mysql/create-table-stop_times.sql",
	)
}

func resources_ddl_mysql_create_table_stop_times_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_stop_times_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-stop_times.sql", size: 366, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_stop_times_full_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x94\x4f\x8f\xda\x30\x10\xc5\xef\x7c\x0a\x5f\x2a\x81\xd4\xad\x44\x45\xab\x95\xaa\x3d\x64\x77\xdd\xed\x4a\x69\x56\xa2\xe1\xd0\x53\x6c\x9c\x49\x62\x61\xec\xd4\x76\x50\xe9\xa7\x6f\xb0\x05\xe1\x4f\x9c\xec\x85\x03\xfe\xf9\x79\xde\xcc\x9b\x3c\x2d\x71\x94\x62\x94\x46\x8f\x31\x46\xe4\x83\x21\x9f\x88\xb1\xaa\xce\x2c\xdf\x82\xc9\x8a\x46\x08\x82\xa6\x13\x84\xfc\xbf\x3c\x27\x88\x4b\x3b\x9d\xcf\x67\x28\x79\x4b\x51\xb2\x8a\xe3\x8f\x93\xbb\xbb\xe3\x39\x53\x39\x10\xb4\xa3\x9a\x55\x54\x4f\x17\x5f\x66\xe8\x19\x7f\x8f\x56\xf1\x91\x3c\x72\x92\x6e\xcf\xb8\xaf\x8b\x10\x97\x83\x61\x1d\x37\xff\x7c\x1f\x02\x05\xb5\x04\x15\x42\x51\x1b\x02\x94\xec\x07\xba\xe2\x4b\x50\x04\xbd\xe0\xb7\x9f\x38\x5d\xfe\xee\xa3\xfe\x29\x09\xae\x05\x61\x83\x9d\x5a\xa3\xc5\x48\x27\x84\x62\xd4\x72\x25\x33\xbb\xaf\xa1\xeb\x6b\x8f\x60\x4d\x35\x48\x9b\x19\xeb\xf8\x11\x59\xaa\x35\xdf\x51\xe1\x26\x48\xd0\xe1\xf7\x96\xc9\xa1\x95\xb4\x8d\x86\x41\xca\xf9\x30\xf0\xa7\x01\xc9\x86\x0b\x74\x64\x05\x34\xcf\x0c\x2f\xcf\x0a\xbc\x99\x97\xb7\xc3\xd9\xa6\xa9\xc7\x6d\xe7\xba\x55\x55\x45\x31\x4e\xd2\xb2\x2d\x71\x7f\x11\xcf\x1e\x4a\xab\xc6\xc2\x00\x74\x42\x4c\xa5\xb4\xbd\x0a\x69\x60\xd6\xfe\x42\x9b\xae\xf2\x8a\xef\x09\xeb\xd9\x85\xcb\x60\xf7\x2e\x80\x07\x47\xad\x7b\x6c\x3c\x6d\x9e\x63\x4a\x28\x4d\x90\x7f\x35\xfc\x28\xfc\xb5\xa3\xa8\xd5\x3c\xf4\x41\x38\x24\x02\xf4\x8e\xb3\xa1\x6e\xbb\xea\x9d\xc8\x21\x38\x97\xb9\xe9\x75\x90\x73\x0d\xcc\x2d\x4c\x48\xd4\x6b\xae\xdb\xc5\xda\xbc\x6b\x4f\x2b\x5a\x0f\x2f\xf4\x64\x86\x70\xf2\xf2\x9a\xe0\x87\x57\x29\xd5\xf3\xe3\xe9\xf0\xe9\x47\xb4\xfc\x85\xd3\x87\xc6\x16\xf7\xdb\xf5\xe2\xdb\xff\x00\x00\x00\xff\xff\x39\x1d\xb9\x15\x47\x05\x00\x00")

func resources_ddl_mysql_create_table_stop_times_full_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_stop_times_full_sql,
		"resources/ddl/mysql/create-table-stop_times_full.sql",
	)
}

func resources_ddl_mysql_create_table_stop_times_full_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_stop_times_full_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-stop_times_full.sql", size: 1351, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xd1\x51\x4b\xc3\x30\x10\x07\xf0\xf7\x7e\x8a\x7b\x11\x5a\x10\xa1\x52\x65\x20\x7b\xc8\xb6\xdb\x2c\x66\xad\xd4\xec\x61\x4f\x4d\x6c\xa3\x16\xba\xa4\x24\x37\x41\x3f\xbd\xd2\x87\x4e\x1c\x5d\xdf\x02\xf9\xdd\x71\xff\xbb\x65\x81\x4c\x20\x08\xb6\xe0\x08\xf2\xca\xcb\x1b\xe9\xc9\x76\x5e\x42\x18\x00\xf4\xef\xb2\xa9\x25\x34\x86\xc2\x38\x8e\x20\xcb\x05\x64\x3b\xce\xaf\x87\xdf\xca\xd6\x5a\xc2\xa7\x72\xd5\x87\x72\x61\x72\x17\xc1\x0a\xd7\x6c\xc7\xff\x3b\xa3\x0e\x7f\xdc\x7d\x32\xe6\x6a\xed\xab\x93\x8b\x6f\x67\x63\xb0\x55\x24\x61\xcd\x73\x26\xc6\x80\x35\x97\xc1\xbb\xb6\x12\x36\x98\x6f\x51\x14\xfb\x73\xf3\x6d\x8d\xee\xc3\x4f\x87\x3b\xba\x76\x82\xb5\xb6\x52\xd4\x58\x53\xd2\x57\xa7\x4f\xfb\x3c\x73\x9d\x72\xda\x50\xe9\xa9\xd7\x97\x9b\x3e\x17\xe9\x96\xfd\x4e\xfe\x84\x7b\x08\x87\x5b\x45\x41\x04\x98\x6d\xd2\x0c\xe7\xa9\x31\x76\xb5\x18\xea\x96\x8f\xac\x78\x41\x31\x3f\xd2\xdb\xec\xf0\x9a\x3c\x04\x3f\x01\x00\x00\xff\xff\x26\x41\x8e\xe9\xfe\x01\x00\x00")

func resources_ddl_mysql_create_table_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_stops_sql,
		"resources/ddl/mysql/create-table-stops.sql",
	)
}

func resources_ddl_mysql_create_table_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-stops.sql", size: 510, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_transfers_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xce\xc1\x4b\xc3\x30\x14\xc7\xf1\x7b\xfe\x8a\x77\x91\xa5\xd0\x88\x83\x09\x82\xec\x90\x6d\x4f\x2d\xc6\x28\x31\x3b\xec\x94\x54\x5d\x31\x87\x24\x25\x89\x82\xff\xbd\xb4\xa8\xad\x9e\xbc\x7f\x7e\xef\x7d\xb7\x0a\xb9\x46\xd0\x7c\x23\x10\xec\x49\xb6\xa7\xb6\xa4\x36\xe4\xee\x98\xb2\x05\x4a\x00\x6c\x97\xa2\x37\xb9\xc4\xde\xb8\x17\x0b\xef\x6d\x7a\x7e\x6d\x13\x5d\x9d\x57\x20\xef\x35\xc8\xbd\x10\xf5\xc0\x4a\xfc\x0f\xfa\xba\x6d\xca\x47\x7f\xb4\xe0\x42\xa1\xcb\xe5\x64\x60\x87\x57\x7c\x2f\x34\x2c\xce\x16\xa3\xf7\x2e\x98\x69\xe3\xfc\x6c\xf3\x4d\xc7\x1d\x63\x03\x67\x0c\x1e\x54\x73\xc7\xd5\x01\x6e\xf1\x00\xf4\x77\x7a\x3d\x4f\xac\xff\xa4\x54\x84\x54\x80\xf2\xba\x91\xb8\x6e\x42\x88\xbb\xcd\xcf\x83\xed\x0d\x57\x8f\xa8\xd7\x6f\xa5\xbb\xf0\x4f\xab\x4b\xf2\x19\x00\x00\xff\xff\xde\x96\xc0\x3c\x34\x01\x00\x00")

func resources_ddl_mysql_create_table_transfers_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_transfers_sql,
		"resources/ddl/mysql/create-table-transfers.sql",
	)
}

func resources_ddl_mysql_create_table_transfers_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_transfers_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-transfers.sql", size: 308, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_create_table_trips_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xcf\xc1\x4b\xc3\x30\x14\xc7\xf1\x7b\xff\x8a\x77\x11\x12\x10\x61\x30\x41\x90\x1d\xb2\xed\xa9\xc5\x18\x25\x66\x87\x9d\x4c\x96\x46\x1b\xd4\xa4\x24\x69\xff\x7e\xa9\x87\xa2\x0c\xda\xf3\xef\xf3\xbe\xf0\x76\x12\x99\x42\x50\x6c\xcb\x11\xf4\x45\xd6\x57\xba\x24\xdf\x65\x0d\xa4\x02\xd0\x29\xf6\xc5\xbd\xf9\x46\x83\x0f\x85\xac\x56\x14\xc4\xb3\x02\x71\xe0\xfc\x72\x9c\xb3\x4b\x83\xb7\x33\x60\x6c\x2d\xac\xad\x33\x4d\xf6\x1f\x41\xc3\x60\x92\x6d\x4d\x22\xeb\x6b\x0a\x7b\xbc\x63\x07\xfe\xc7\x36\x3e\x39\x5b\x7c\x0c\xff\x72\x67\xec\xf4\x15\xed\xe7\x2f\x99\xad\xe5\xd6\x74\x6e\x91\xbd\xc8\xfa\x89\xc9\x23\x3c\xe2\x11\xc8\xf4\x0b\xad\x28\xa0\xb8\xaf\x05\x6e\xea\x10\xe2\x7e\x3b\xdd\xed\x1e\x98\x7c\x45\xb5\xe9\xcb\xfb\xcd\xf7\x69\x7d\x5b\xfd\x04\x00\x00\xff\xff\xc8\x64\xeb\x18\x5e\x01\x00\x00")

func resources_ddl_mysql_create_table_trips_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_create_table_trips_sql,
		"resources/ddl/mysql/create-table-trips.sql",
	)
}

func resources_ddl_mysql_create_table_trips_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_create_table_trips_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/create-table-trips.sql", size: 350, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_delete_agency_by_key_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x71\xf5\x71\x0d\x71\x55\x70\x0b\xf2\xf7\x55\x48\x2f\x49\x2b\xd6\x4b\x4c\x4f\xcd\x4b\xce\x4c\x2d\x56\x08\xf7\x70\x0d\x72\x55\x00\x73\x2b\xe3\xb3\x53\x2b\x6d\xed\xad\x01\x01\x00\x00\xff\xff\x87\x7b\xb2\x17\x2d\x00\x00\x00")

func resources_ddl_mysql_delete_agency_by_key_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_delete_agency_by_key_sql,
		"resources/ddl/mysql/delete-agency-by-key.sql",
	)
}

func resources_ddl_mysql_delete_agency_by_key_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_delete_agency_by_key_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/delete-agency-by-key.sql", size: 45, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_drop_table_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x50\x2d\x4e\xd0\x03\x11\x80\x00\x00\x00\xff\xff\x32\x81\xd6\x0f\x1e\x00\x00\x00")

func resources_ddl_mysql_drop_table_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_drop_table_sql,
		"resources/ddl/mysql/drop-table.sql",
	)
}

func resources_ddl_mysql_drop_table_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_drop_table_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/drop-table.sql", size: 30, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_insert_line_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xd0\x41\x4e\xc4\x30\x0c\x05\xd0\xf5\xe4\x14\xde\x20\x81\x84\x72\x01\x34\x67\x69\xa1\x31\xc2\x28\xb5\x91\x6d\x16\xdc\x9e\x38\x93\x69\x36\x9d\xae\xbe\xdb\xbc\x7c\xd5\xc4\x86\xea\x40\xec\x02\xeb\x93\xad\x79\xad\xc4\xb8\x98\xcb\x8f\xad\xf0\xdc\x07\x2a\xaf\x10\x2f\x66\xd8\xa4\xe0\x0b\x18\x56\xdc\x1c\x0a\x99\x13\x6f\x9e\x2e\x50\xf3\x04\x79\x92\x7c\xa0\xf4\xa9\xb2\xa7\xcb\xad\x69\x94\x58\x6b\x67\x54\xf8\x16\xe2\x04\xed\x99\x5f\x17\xa7\x1d\xe3\x88\x83\xf0\xbc\xf2\x6a\x7e\x8f\x70\x8e\x5d\x29\xae\xee\xcc\x73\x4c\x83\x8d\xf8\x80\xa9\xfc\x7a\xf4\xe9\xcd\xf5\x31\xa0\x1e\xf1\xdc\xc5\x5f\x37\x56\x83\xdd\xcf\xda\x97\xa8\x2f\xfc\xbe\xe3\x75\xac\x25\x72\x12\x2d\xcd\x7f\xfc\x75\x7d\xb6\xaf\xb7\xf4\x1f\x00\x00\xff\xff\xfb\xee\x37\x1b\x94\x01\x00\x00")

func resources_ddl_mysql_insert_line_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_insert_line_stops_sql,
		"resources/ddl/mysql/insert-line_stops.sql",
	)
}

func resources_ddl_mysql_insert_line_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_insert_line_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/insert-line_stops.sql", size: 404, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_insert_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcc\x31\x0e\xc2\x40\x0c\x44\xd1\x7e\x4f\x31\x0d\x12\x34\xb9\x00\x87\x89\x43\x62\x84\xa5\xc4\x96\xc6\xa6\xe0\xf6\x2c\x8b\x44\x87\x1b\x37\x7f\x9e\x79\x2a\x0b\xe6\x15\x90\x53\xca\x24\xbb\xb9\xa6\xe0\xfc\xf9\xb3\x2f\x87\x5e\x90\xba\xeb\x5a\xd8\x2c\xcb\x7c\xad\x86\x7e\x9c\x18\xcf\xd2\x39\x1f\xc1\x1a\x1d\x96\xc4\x6f\xd4\xee\x8c\x63\x84\x5f\x75\xc4\x9d\x65\x0b\x6e\x4a\xdc\x5e\x7f\x94\x6b\x7b\x07\x00\x00\xff\xff\x7f\x86\x90\x71\x93\x00\x00\x00")

func resources_ddl_mysql_insert_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_insert_lines_sql,
		"resources/ddl/mysql/insert-lines.sql",
	)
}

func resources_ddl_mysql_insert_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_insert_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/insert-lines.sql", size: 147, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_insert_route_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\x51\x0a\xc2\x30\x0c\x86\xdf\x7b\x8a\xbc\x08\x0a\xd2\x0b\xc8\xce\xb2\x42\x1b\x21\xe2\x12\x49\xe2\xfd\x6d\x57\xb7\x8a\xac\x0f\x25\x69\xf3\xfd\x7f\x7e\x62\x43\x75\x20\x76\x81\x74\xb2\x14\x93\xca\xdb\x71\x36\x97\x97\x25\x38\xf7\x8e\xca\x15\xda\xcb\x28\xb2\x14\xbc\x80\xe1\x13\xb3\x43\x21\x73\xe2\xec\x01\xea\xd1\xf8\xc3\xc4\x41\xc5\x9d\x0b\x77\x95\x65\x9d\xed\x8e\x5f\x2f\xab\x5b\x30\x2a\x3c\x84\xf8\xef\x77\x76\x5a\xb0\x8d\x38\x08\xd7\x7b\x93\x85\x69\x58\x1c\xd3\xae\xd4\xb4\x37\xae\xb5\x9d\x1b\xf5\x21\xb7\x66\xa8\xa0\x36\xd0\xf7\x48\xd3\x48\x77\x0b\x9f\x00\x00\x00\xff\xff\x49\xec\x45\xe9\x3c\x01\x00\x00")

func resources_ddl_mysql_insert_route_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_insert_route_stops_sql,
		"resources/ddl/mysql/insert-route_stops.sql",
	)
}

func resources_ddl_mysql_insert_route_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_insert_route_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/insert-route_stops.sql", size: 316, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_insert_station_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8f\xd1\xca\xc2\x30\x0c\x85\xef\xfb\x14\xb9\xf9\xe1\x17\xa4\x2f\x20\x7b\x96\x55\xb6\x0a\x91\x9a\x8c\x9e\xbc\x3f\x36\x73\x53\x8b\xda\xab\xa4\x9c\xef\x3b\x84\x05\xb9\x1a\xb1\x98\x52\xfa\x43\x8a\x09\x76\x36\x56\x19\x0b\x4b\x46\xa2\xff\x7d\xe7\xf9\x48\xfe\xd7\x86\x03\x21\x97\x3c\x19\xcd\x0c\x63\x99\x2c\x50\x7b\x88\x5d\x14\x71\x4b\x87\x4b\xd5\xdb\x9a\xe8\x0a\x9a\x1b\xad\x57\x72\xa5\xab\xb2\x7c\x06\x46\x98\x2e\x9e\x02\xa9\x74\xf6\x01\xef\xdb\x77\xc9\x5a\xbe\x19\xca\xc3\xe0\x90\x2e\xce\x97\xe7\xf8\x1b\x76\xce\xb1\xb2\xdf\x31\xbc\x4e\x3a\x85\x7b\x00\x00\x00\xff\xff\x3b\x15\x92\xb5\x38\x01\x00\x00")

func resources_ddl_mysql_insert_station_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_insert_station_lines_sql,
		"resources/ddl/mysql/insert-station_lines.sql",
	)
}

func resources_ddl_mysql_insert_station_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_insert_station_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/insert-station_lines.sql", size: 312, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_insert_station_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8d\x51\x0a\x02\x31\x0c\x44\xff\x7b\x8a\xf9\x11\x14\xa4\x17\x10\xcf\x62\x97\xdd\x08\x11\x37\x91\x66\xee\x8f\x6d\xb1\xe0\xe6\x2b\x99\x3c\xe6\xa9\x85\x54\x42\x8d\x8e\x72\x8a\x92\x4b\x70\xa1\xba\x3d\x82\xfe\x89\x82\xf3\xbc\x75\xbb\xa2\x67\x6d\xb9\x20\xe4\x2d\x2b\xb1\x69\x50\x6d\x65\x42\x9b\x60\x3e\xb0\xf9\x47\xa7\x67\xf5\x7d\x10\x53\x30\x8a\x03\xcd\x6a\x52\xf1\x72\xb5\xc3\x7b\x74\x74\x82\x70\x9b\x3d\xb6\xec\x82\xfb\xbf\xa4\x27\xb7\xf4\x0d\x00\x00\xff\xff\x2a\x51\x48\x3d\xc1\x00\x00\x00")

func resources_ddl_mysql_insert_station_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_insert_station_stops_sql,
		"resources/ddl/mysql/insert-station_stops.sql",
	)
}

func resources_ddl_mysql_insert_station_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_insert_station_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/insert-station_stops.sql", size: 193, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_insert_stations_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8f\xb1\x8e\xc2\x30\x0c\x86\xf7\x3c\x85\x97\x53\x12\xa9\xea\x0b\xdc\x7e\xa7\x5b\x0e\x86\xee\x24\x04\x53\x45\x6a\xed\xaa\x36\x08\xde\x9e\x50\x5a\x51\xd1\x4c\xf9\xbe\xe1\x93\xff\x4c\x82\xa3\x42\x26\x65\x08\x5f\x12\xea\x20\x1a\x35\x33\x49\x00\x37\x7f\x0f\x14\x7b\xac\x60\xa1\x2e\xea\x0a\x98\xde\xd0\x22\x7b\x10\xec\x30\x29\x9c\xb2\x68\xa6\xa4\x06\xca\x93\x5a\x94\x87\x57\x67\x12\xf1\xda\xba\x59\x96\x9c\xdf\x4a\xa6\x59\xfe\x22\xf7\x3f\x23\xf7\x0d\xde\xd4\x25\xa6\x14\xd5\xd9\xfd\xee\xef\xbf\x71\xb6\xda\x74\xc0\xc2\x87\x7d\x86\xc0\x7a\xeb\xbd\x39\x97\xcc\xd4\x5c\x86\xf2\x50\x56\x8a\x69\x47\xbe\x0c\x70\xbc\xaf\xcf\xfc\x36\x8f\x00\x00\x00\xff\xff\xa4\x41\x1c\xc9\x1a\x01\x00\x00")

func resources_ddl_mysql_insert_stations_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_insert_stations_sql,
		"resources/ddl/mysql/insert-stations.sql",
	)
}

func resources_ddl_mysql_insert_stations_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_insert_stations_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/insert-stations.sql", size: 282, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_insert_stop_times_full_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\x4b\x96\xa3\x30\x0c\x5c\xc3\x29\xd8\xe4\x65\x93\xf8\x06\x39\x0b\x78\x6c\x01\x9e\x38\x16\x23\x8b\x4c\xa7\x4f\xdf\xe2\x17\xe3\x34\xac\x70\x49\x55\xb6\x4a\xe5\x42\x04\xe2\xca\x05\xc6\xaa\x39\xc5\x46\x35\x91\x71\xa8\xd9\x3d\x20\xd6\xed\xe8\x7d\x53\x46\xf0\x60\xb8\x2a\x2b\xf9\xa2\x9a\xcb\xce\x5e\xca\xeb\x75\x07\x18\xb4\x70\xd9\x77\x04\xfd\xc8\x01\x0b\xd1\x64\x80\xd7\x9c\x9f\x31\x7c\x88\x76\x80\x09\xf9\xc6\x00\xbf\xef\x1d\xc9\x6f\x22\x1e\x8d\x66\x87\xa1\xe6\xd7\x00\xa9\x6d\xd0\x04\x81\xeb\xc8\x73\xf1\x52\x16\x91\x95\x26\x72\x4f\xed\xe7\x29\x57\x3a\x2b\x0b\xd2\xca\x23\x41\x0e\xcf\xd7\x44\xf8\x37\x42\x30\x49\x76\xc5\x7b\xd0\xb6\x8e\xae\x0b\xbb\xc2\xe0\xcc\x7d\x1c\xf2\x57\x88\x3a\x49\x3b\xb6\x6d\x86\x93\xd2\x9d\xc8\xbe\x76\x73\x91\x22\x1c\x79\x99\xb4\xd8\x0e\xb1\x47\xe2\xd5\xd2\xbc\x4d\x4c\xeb\x0e\x0b\xc9\xee\x0d\xc9\x2f\x5e\xb0\xb7\x7d\x1b\x60\xd0\x23\x7d\xf0\xe0\x8b\xf7\x38\x2b\x26\xb7\x44\x60\x39\x4a\x80\x9e\xce\x6c\xcb\x29\xd6\xfa\x64\xcd\xe2\xcc\xd2\x65\x1d\x49\x8a\xa6\xfd\x38\xbb\xbe\x82\xd5\x1f\x59\xda\x7d\x37\xbd\x88\xf5\x7a\x98\xa4\xca\x96\xf0\x51\x16\x29\x92\xb1\xa9\xa2\xc4\x34\x00\x55\x7f\xd1\x85\x59\xf5\x33\xb0\xd2\xc2\x15\x86\x94\xd2\xdb\xb6\x28\x67\x8f\xb9\xd3\x53\x85\xb6\xb0\xde\x93\xdd\xde\x7f\xc7\xac\xd9\x18\xa1\xd1\x44\x4b\x1b\x13\xda\xf6\x5b\xfe\xef\x81\xe0\x68\x83\xb7\xf3\xe9\x79\xfe\x09\x00\x00\xff\xff\x4e\x2c\x60\x2c\x76\x03\x00\x00")

func resources_ddl_mysql_insert_stop_times_full_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_insert_stop_times_full_sql,
		"resources/ddl/mysql/insert-stop_times_full.sql",
	)
}

func resources_ddl_mysql_insert_stop_times_full_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_insert_stop_times_full_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/insert-stop_times_full.sql", size: 886, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_select_agency_zone_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4e\xcd\x49\x4d\x2e\x51\xc8\xcd\xcc\xd3\x28\xd6\x4b\x28\x2e\xc9\x2f\x88\xcf\x49\x2c\x49\xd0\xd4\x51\xc8\x4d\xac\xc0\x10\x42\x56\x95\x9f\x87\xa1\x0a\x24\xa4\x90\x56\x94\x9f\xab\x90\xa0\x5a\x9c\x00\x11\x2d\x4e\x50\x28\x06\x04\x00\x00\xff\xff\xad\x63\x94\xbb\x65\x00\x00\x00")

func resources_ddl_mysql_select_agency_zone_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_select_agency_zone_sql,
		"resources/ddl/mysql/select-agency-zone.sql",
	)
}

func resources_ddl_mysql_select_agency_zone_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_select_agency_zone_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/select-agency-zone.sql", size: 101, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_select_trip_stop_times_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8e\x5d\xaa\xc2\x40\x0c\x46\xb7\x92\x97\xd2\x97\x4b\x77\x70\xd7\xd2\x19\x3b\x9f\x18\xe9\x64\x6a\x92\x2a\xee\xde\xa6\xe2\x0f\xf8\x16\xce\xe1\x3b\xc4\x30\x63\x72\x32\x1f\xb2\x2a\x5f\xf3\x3c\x3a\x57\xfc\x05\x28\x58\xb2\xfa\xaa\xf8\x20\xf3\xb6\x8c\x86\xcb\x0a\x99\x82\x3c\x81\xe4\x0a\x3a\x6a\xab\x94\x3a\x4b\x43\xda\x61\x6c\x2c\x6d\x23\x62\x11\x28\x9d\x1b\xcb\x97\x0f\x45\x4d\xde\x51\x2e\xff\xf6\xba\xe8\x76\x82\x22\x94\x2b\xef\xaa\xef\xac\xa7\xa6\x65\xeb\x1c\xee\x3f\x8f\x3c\x02\x00\x00\xff\xff\x88\x54\xd1\x98\xc4\x00\x00\x00")

func resources_ddl_mysql_select_trip_stop_times_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_select_trip_stop_times_sql,
		"resources/ddl/mysql/select_trip_stop_times.sql",
	)
}

func resources_ddl_mysql_select_trip_stop_times_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_select_trip_stop_times_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/select_trip_stop_times.sql", size: 196, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_select_trips_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4e\xcd\x49\x4d\x2e\x51\x28\x29\xca\x2c\x88\xcf\x4c\x51\x48\x2b\xca\xcf\x55\x48\x50\x2d\x4e\xd0\x4b\x00\x09\x15\x27\x28\xe4\x17\xa5\xa4\x16\x29\x24\x55\xc2\x94\x00\x02\x00\x00\xff\xff\x54\x32\xa7\x83\x31\x00\x00\x00")

func resources_ddl_mysql_select_trips_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_select_trips_sql,
		"resources/ddl/mysql/select_trips.sql",
	)
}

func resources_ddl_mysql_select_trips_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_select_trips_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/select_trips.sql", size: 49, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_update_agency_zone_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x2d\x48\x49\x2c\x49\x55\x48\x50\x2d\x4e\xd0\x4b\x48\x4c\x4f\xcd\x4b\xce\x4c\x2d\x4e\x50\x28\x4e\x2d\x51\x80\x70\x2b\xe3\x73\x33\xf3\xe2\x73\x12\x4b\x12\x6c\xed\x75\x10\x62\x89\x15\x98\x62\x20\x75\xf9\x79\x98\xea\xc0\x62\x5c\x80\x00\x00\x00\xff\xff\x72\x6b\x2a\xe4\x6a\x00\x00\x00")

func resources_ddl_mysql_update_agency_zone_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_update_agency_zone_sql,
		"resources/ddl/mysql/update-agency-zone.sql",
	)
}

func resources_ddl_mysql_update_agency_zone_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_update_agency_zone_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/update-agency-zone.sql", size: 106, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_mysql_update_gtfs_agency_zone_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x2d\x48\x49\x2c\x49\x55\x48\x48\x2f\x49\x2b\x4e\xd0\x4b\x48\x4c\x4f\xcd\x4b\xce\x4c\x2d\x4e\x50\x28\x4e\x2d\x51\x80\x70\x2b\xe3\x73\x33\xf3\xe2\x73\x12\x4b\x12\x6c\xed\x75\x10\x62\x89\x15\x98\x62\x20\x75\xf9\x79\x98\xea\xc0\x62\x0a\xe5\x19\xa9\x45\xa9\x70\x89\xec\xd4\x4a\xa0\x20\x17\x20\x00\x00\xff\xff\x86\xf4\xa1\xf0\x81\x00\x00\x00")

func resources_ddl_mysql_update_gtfs_agency_zone_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_mysql_update_gtfs_agency_zone_sql,
		"resources/ddl/mysql/update-gtfs-agency-zone.sql",
	)
}

func resources_ddl_mysql_update_gtfs_agency_zone_sql() (*asset, error) {
	bytes, err := resources_ddl_mysql_update_gtfs_agency_zone_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/mysql/update-gtfs-agency-zone.sql", size: 129, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_index_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\xf0\xf4\x73\x71\x8d\x50\x50\x2d\x8e\x07\xa2\xcc\x94\x0a\x05\x7f\x3f\x20\x47\x4f\xb5\x58\x41\x43\xb5\x58\xd3\x1a\x10\x00\x00\xff\xff\x26\x15\x66\xe5\x25\x00\x00\x00")

func resources_ddl_postgres_create_index_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_index_sql,
		"resources/ddl/postgres/create-index.sql",
	)
}

func resources_ddl_postgres_create_index_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_index_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-index.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_schema_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x76\xf6\x70\xf5\x75\x54\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x50\x2d\xb6\x06\x04\x00\x00\xff\xff\x89\x07\x73\x65\x1f\x00\x00\x00")

func resources_ddl_postgres_create_schema_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_schema_sql,
		"resources/ddl/postgres/create-schema.sql",
	)
}

func resources_ddl_postgres_create_schema_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_schema_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-schema.sql", size: 31, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_spatial_index_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\xf0\xf4\x73\x71\x8d\x50\x50\x2d\x8e\x07\xa2\xcc\x94\x0a\x05\x7f\x3f\x20\x47\x4f\xb5\x58\x21\x34\xd8\xd3\xcf\x5d\xc1\xdd\x33\x38\x44\x41\x43\xb5\x58\xd3\x1a\x10\x00\x00\xff\xff\x57\x9d\x38\xbf\x30\x00\x00\x00")

func resources_ddl_postgres_create_spatial_index_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_spatial_index_sql,
		"resources/ddl/postgres/create-spatial-index.sql",
	)
}

func resources_ddl_postgres_create_spatial_index_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_spatial_index_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-spatial-index.sql", size: 48, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_agencies_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xd1\xdd\x0a\x82\x30\x14\xc0\xf1\x7b\x9f\xe2\xdc\x04\x0a\xe1\x0b\x74\xb5\x6a\x41\xb4\x2c\x44\x2f\xbc\x92\x61\x6b\x0e\xe6\x31\x74\x41\x26\xbe\x7b\x1a\x64\x41\xf8\x71\x33\x18\xe7\xcf\x8f\x03\x67\xe3\x53\x12\x50\x08\xc8\x9a\x51\x58\x94\x2e\x97\x02\x13\x25\x4a\xb0\x2d\x80\xf7\xa7\x8a\xd5\x05\x14\x1a\x21\x45\x01\xde\x29\x00\x2f\x64\x6c\xf9\x9d\x22\xcf\x04\x24\x29\x2f\xec\xba\xd6\x02\xa5\x49\xa1\x57\xdc\x9f\xa6\x69\x1c\xd8\xd2\x1d\x09\xd9\x1f\x71\x2f\xf4\x94\xd0\x26\x23\x80\x51\x99\x78\xe6\x38\xb9\xc7\xa7\x1b\xa1\x34\x47\x39\xc5\x74\x4d\x6c\xaa\xdb\x98\x93\x29\x6c\x3b\x03\x57\x9d\xb7\xef\x50\xc4\x1f\x33\xa2\x4e\xca\x71\x86\x34\x18\x9d\xfd\xfd\x91\xf8\x11\x1c\x68\x04\x76\x7f\x55\xc7\x72\x56\xd6\x2b\x00\x00\xff\xff\xfc\xca\x27\xfc\x01\x02\x00\x00")

func resources_ddl_postgres_create_table_agencies_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_agencies_sql,
		"resources/ddl/postgres/create-table-agencies.sql",
	)
}

func resources_ddl_postgres_create_table_agencies_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_agencies_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-agencies.sql", size: 513, mode: os.FileMode(420), modTime: time.Unix(1429220662, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_calendar_dates_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\x4b\x4e\xcc\x49\xcd\x4b\x49\x2c\x8a\x4f\x49\x2c\x49\x2d\x56\xd0\xe0\x52\x50\x28\x4e\x2d\x2a\xcb\x4c\x4e\x8d\xcf\x4c\x51\xc8\xcc\x2b\x49\x4d\x4f\x2d\x52\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\xd1\x01\x4a\x83\x14\x42\x08\x64\xd1\xd4\x8a\xe4\xd4\x82\x92\xcc\xfc\xbc\xf8\x92\xca\x82\x54\xb8\x46\x17\x57\x37\xc7\x50\x1f\x84\xb2\x80\x20\x4f\x5f\xc7\xa0\x48\x05\x6f\xd7\x48\x05\x0d\x84\x45\x3a\x60\x03\x35\xb9\x34\xad\xb9\x00\x01\x00\x00\xff\xff\x41\x4b\x1b\xbd\xa1\x00\x00\x00")

func resources_ddl_postgres_create_table_calendar_dates_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_calendar_dates_sql,
		"resources/ddl/postgres/create-table-calendar_dates.sql",
	)
}

func resources_ddl_postgres_create_table_calendar_dates_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_calendar_dates_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-calendar_dates.sql", size: 161, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_calendars_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\xd0\xb1\x8a\xc2\x40\x10\xc6\xf1\x7e\x9f\x62\x9a\x83\x04\x8e\x7b\x81\xab\x56\x5d\x41\x5c\xa3\x84\xa4\x48\x15\xd6\xec\xa8\x0b\x71\x03\xb3\xb3\x8a\x6f\x6f\xb0\x09\x58\x4c\x33\xcd\xef\xdf\xcc\xb7\xae\x8d\x6e\x0c\x34\x7a\x65\x0d\xfc\xa4\xbf\xc1\x8d\x18\xbd\xa3\x04\x85\x02\x48\x48\x8f\x30\x60\x1f\x3c\x84\xc8\x78\x45\x82\xea\xd8\x40\xd5\x5a\xfb\x3b\xf3\x7d\x9a\xd3\x17\x9c\xa7\x69\x44\x17\x61\x63\xb6\xba\xb5\x0b\x73\xc6\x24\xf9\x13\x7d\x94\x0b\xbe\x65\x12\x83\x0b\x05\x89\x93\xe3\x4c\x62\x90\xc5\x0f\x12\x3b\xe2\xde\x3b\x46\xf8\x9c\x6f\x9f\xa7\x12\xf4\x54\xef\x0e\xba\xee\x60\x6f\x3a\x28\x96\x29\x4b\x55\xfe\x2b\xf5\x0e\x00\x00\xff\xff\xf8\x8d\x5c\xc5\x79\x01\x00\x00")

func resources_ddl_postgres_create_table_calendars_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_calendars_sql,
		"resources/ddl/postgres/create-table-calendars.sql",
	)
}

func resources_ddl_postgres_create_table_calendars_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_calendars_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-calendars.sql", size: 377, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_line_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\xcb\xc9\xcc\x4b\x8d\x2f\x2e\xc9\x2f\x28\x56\xd0\xe0\x52\x50\x00\x73\x33\x53\x14\x32\xf3\x4a\x52\xd3\x53\x8b\x14\xfc\xfc\x43\x14\xfc\x42\x7d\x7c\x74\x80\x72\x20\x55\x78\xe5\x92\xf3\x53\x52\x15\x92\x33\x12\x8b\x34\xaa\xab\x73\x52\xf3\xd2\x4b\x32\x14\xf4\xc0\x46\xeb\xc1\xa5\x6b\x6b\x35\x51\xf4\x05\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\x2a\x68\x40\x2d\xd7\x81\x5a\xa4\xc9\xa5\x69\xcd\x05\x08\x00\x00\xff\xff\x73\xa3\xda\x97\xb0\x00\x00\x00")

func resources_ddl_postgres_create_table_line_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_line_stops_sql,
		"resources/ddl/postgres/create-table-line_stops.sql",
	)
}

func resources_ddl_postgres_create_table_line_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_line_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-line_stops.sql", size: 176, mode: os.FileMode(420), modTime: time.Unix(1429223652, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\xcb\xc9\xcc\x4b\x2d\x56\xd0\xe0\x52\x50\x00\xb1\xe2\x33\x53\x14\x82\x5d\x83\x3c\x1d\x7d\x74\x60\x22\x79\x89\xb9\xa9\x0a\xc9\x19\x89\x45\x1a\xd5\xd5\x39\xa9\x79\xe9\x25\x19\x0a\x7a\x45\xf9\xa5\x25\xa9\xc5\x10\x2a\xbe\x38\x23\xbf\xa8\x04\xac\xac\xb6\x56\x53\xc1\xcf\x3f\x44\xc1\x2f\xd4\x07\xac\x3f\x20\xc8\xd3\xd7\x31\x28\x52\xc1\xdb\x35\x52\x41\x03\x6a\xbc\x26\x97\xa6\x35\x17\x20\x00\x00\xff\xff\x03\xaf\x4c\x2c\x85\x00\x00\x00")

func resources_ddl_postgres_create_table_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_lines_sql,
		"resources/ddl/postgres/create-table-lines.sql",
	)
}

func resources_ddl_postgres_create_table_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-lines.sql", size: 133, mode: os.FileMode(420), modTime: time.Unix(1429223652, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_route_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\x2b\xca\x2f\x2d\x49\x8d\x2f\x2e\xc9\x2f\x28\x56\xd0\xe0\x52\x50\x80\xf0\x33\x53\x14\x32\xf3\x4a\x52\xd3\x53\x8b\x14\xfc\xfc\x43\x14\xfc\x42\x7d\x7c\x74\x80\x92\x20\x65\x78\xe5\x92\xf3\x53\x52\x15\x92\x33\x12\x8b\x34\xaa\xab\x73\x52\xf3\xd2\x4b\x32\x14\xf4\xc0\x66\xeb\xc1\xa5\x6b\x6b\x35\x51\xf4\x05\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\x2a\x68\xc0\x6c\xd7\x81\xda\xa4\xc9\xa5\x69\xcd\x05\x08\x00\x00\xff\xff\x50\x09\x90\x07\xb3\x00\x00\x00")

func resources_ddl_postgres_create_table_route_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_route_stops_sql,
		"resources/ddl/postgres/create-table-route_stops.sql",
	)
}

func resources_ddl_postgres_create_table_route_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_route_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-route_stops.sql", size: 179, mode: os.FileMode(420), modTime: time.Unix(1429223652, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_routes_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\x2b\xca\x2f\x2d\x49\x2d\x56\xd0\xe0\x52\x50\x00\x33\xe3\x33\x53\x14\x32\xf3\x4a\x52\xd3\x53\x8b\x14\xfc\xfc\x43\x14\xfc\x42\x7d\x7c\x74\x80\x92\x89\xe9\xa9\x79\xc9\x95\xb8\x64\x21\x5a\x8b\x33\xf2\x8b\x4a\xe2\xf3\x12\x73\x53\x15\x92\x33\x12\x8b\x34\xaa\xab\x73\x52\xf3\xd2\x4b\x32\x14\xa0\xb6\xe8\xa1\x2b\xab\xad\xd5\x54\x70\x71\x75\x73\x0c\xf5\x41\x37\x2a\x27\x3f\x2f\x9d\xb0\x49\x70\x55\xb8\x0d\x4a\x49\x2d\x4e\xc6\x6b\x06\x48\x01\x6e\xed\x25\x95\x05\xa9\x70\x1f\x63\x57\x52\x5a\x94\x83\xd7\x02\xa0\x3c\x6e\xf3\x93\xf3\x73\xf2\x8b\x20\xda\xcd\x70\xba\x21\xb5\xa2\x84\x80\xc2\x80\x20\x4f\x5f\xc7\xa0\x48\x05\x6f\xd7\x48\x05\x0d\x58\x3c\x6a\x72\x69\x5a\x73\x01\x02\x00\x00\xff\xff\x03\xb8\xa9\xb1\xf0\x01\x00\x00")

func resources_ddl_postgres_create_table_routes_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_routes_sql,
		"resources/ddl/postgres/create-table-routes.sql",
	)
}

func resources_ddl_postgres_create_table_routes_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_routes_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-routes.sql", size: 496, mode: os.FileMode(420), modTime: time.Unix(1429220662, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_station_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\x2b\x2e\x49\x2c\xc9\xcc\xcf\x8b\xcf\xc9\xcc\x4b\x2d\x56\xd0\xe0\x52\x50\x80\x89\x64\xa6\x28\x64\xe6\x95\xa4\xa6\xa7\x16\x29\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\xe8\x00\xa5\x41\x0a\x71\xc9\x05\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\x2a\x68\x20\xcc\xd1\x81\x69\xd2\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\xc3\xf5\x39\xf3\x82\x00\x00\x00")

func resources_ddl_postgres_create_table_station_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_station_lines_sql,
		"resources/ddl/postgres/create-table-station_lines.sql",
	)
}

func resources_ddl_postgres_create_table_station_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_station_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-station_lines.sql", size: 130, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_station_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\x2b\x2e\x49\x2c\xc9\xcc\xcf\x8b\x2f\x2e\xc9\x2f\x28\x56\xd0\xe0\x52\x50\x80\x89\x64\xa6\x28\x64\xe6\x95\xa4\xa6\xa7\x16\x29\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\xe8\x80\xa5\xf3\x0b\x70\xc9\x05\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\x2a\x68\x20\xcc\xd1\x81\x69\xd2\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\xe4\x05\xea\xc1\x82\x00\x00\x00")

func resources_ddl_postgres_create_table_station_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_station_stops_sql,
		"resources/ddl/postgres/create-table-station_stops.sql",
	)
}

func resources_ddl_postgres_create_table_station_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_station_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-station_stops.sql", size: 130, mode: os.FileMode(420), modTime: time.Unix(1429222755, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_stations_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x50\x2d\xd6\x2b\x2e\x49\x2c\xc9\xcc\xcf\x2b\x56\xd0\xe0\x52\x50\x80\x72\xe2\x33\x53\x14\x82\x5d\x83\x3c\x1d\x7d\x74\x90\x04\xf3\x12\x73\x53\x15\x92\x33\x12\x8b\x34\xaa\xab\x73\x52\xf3\xd2\x4b\x32\x14\x80\xba\xf3\x0b\x8a\xc1\x24\x58\xba\xb6\x56\x53\xc1\xcf\x3f\x44\xc1\x2f\xd4\x07\x45\x6b\x4e\x62\x89\x42\x5a\x4e\x3e\x90\x74\x71\x75\x73\x0c\xf5\xc1\xa2\x22\x3f\x8f\x80\x8a\xf4\xd4\x7c\x05\x20\xce\x4d\x2d\x29\xaa\x44\xb1\x24\x20\xc8\xd3\xd7\x31\x28\x52\xc1\xdb\x35\x52\x41\x03\xe1\x03\x4d\x2e\x4d\x6b\x2e\x2e\x40\x00\x00\x00\xff\xff\x31\x6c\xb9\xa1\xef\x00\x00\x00")

func resources_ddl_postgres_create_table_stations_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_stations_sql,
		"resources/ddl/postgres/create-table-stations.sql",
	)
}

func resources_ddl_postgres_create_table_stations_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_stations_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-stations.sql", size: 239, mode: os.FileMode(420), modTime: time.Unix(1429223652, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_stop_times_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xcf\xcf\x0a\x82\x40\x10\x06\xf0\xbb\x4f\x31\x97\x40\x21\x7c\x81\x4e\x56\xdb\x69\x31\x08\x3d\xcb\xe2\x8e\xba\x64\xeb\x36\xbb\x06\x21\xbe\x7b\xcb\x1e\xfa\x43\x49\xc7\xe1\xfb\xf1\x7d\xcc\xee\xc4\xb2\x82\x41\x91\x6d\x39\x83\x95\x4d\xad\x1b\x4c\xe5\xd4\x05\x2d\xc4\x11\x80\x23\x65\x2a\x25\x41\x69\x87\x2d\x12\xe4\xc7\x02\xf2\x92\xf3\xb5\xcf\x04\x91\xba\x89\x3e\xe8\x00\xc8\x5f\xb0\x67\x87\xac\xe4\x2f\x25\xd1\x08\x72\x23\xe1\x1f\x17\x86\x17\x96\x42\x66\xf1\x3a\xa2\xae\xf1\x29\x7e\x36\x74\x28\x64\x65\x55\xab\xa1\xee\x04\xc5\xd3\xd4\xa3\x6e\x5d\x07\x6f\x8f\xa5\x9f\x70\x9e\x93\xaf\x2a\xa3\xea\xf3\xe8\xf9\xdd\x2c\xcf\x49\xf2\x2d\x43\xd3\x2c\xab\x28\xd9\x44\x8f\x00\x00\x00\xff\xff\xf2\xc6\xf8\xb2\x60\x01\x00\x00")

func resources_ddl_postgres_create_table_stop_times_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_stop_times_sql,
		"resources/ddl/postgres/create-table-stop_times.sql",
	)
}

func resources_ddl_postgres_create_table_stop_times_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_stop_times_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-stop_times.sql", size: 352, mode: os.FileMode(420), modTime: time.Unix(1429220662, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_stop_times_full_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x94\x3d\xaf\xda\x30\x14\x86\x77\x7e\x85\x97\x4a\x20\x15\xc6\x2e\x9d\x68\x9b\x76\x49\x41\x42\x61\xe8\x14\xb9\xce\x49\x62\x61\xec\x5c\xdb\x41\x97\x8b\xf8\xef\xd7\x71\x3e\x44\xe0\x38\x59\x32\xe4\x7d\xde\xf3\x61\x9f\xe3\x9f\x87\x68\x9b\x44\x24\xd9\xfe\x88\x23\xf2\xc5\x6c\x8c\x55\x55\x6a\xf9\x19\x4c\x9a\xd7\x42\x90\xe5\x82\x10\xff\x8f\x67\x84\x4b\x0b\x05\x68\xb2\xdb\x27\x64\x77\x8c\xe3\xaf\x8b\xf5\xba\x53\x99\xca\x80\xb0\x92\xea\xa5\x8f\xd0\xc5\x69\xfe\xae\xc8\xaf\xe8\xf7\xf6\x18\xf7\x96\xce\x20\xe9\xb9\x33\xdc\x6e\x02\x64\x61\x4b\xf2\xe8\x6c\xe4\xfb\x3d\xe0\xcd\xc0\xb0\x09\x6f\x23\x07\xbd\x82\x5a\x92\x0b\xe5\xbe\xb8\xac\x24\x2a\x0f\x8d\x16\xa0\xc8\x9f\x68\xff\x37\x4a\x0e\xff\x10\xe6\x43\x49\x68\x8e\x0a\xad\xae\x13\x5f\x6b\x1b\xa2\xd7\x5a\x4c\x34\xe6\x54\xac\x2f\xa1\x18\xb5\x5c\xc9\xd4\x5e\x2b\x18\x2e\xe9\x35\x43\x45\x35\x48\x9b\x1a\xeb\x69\x3c\xcf\x98\xc1\xb2\x51\xad\xf9\x85\x0a\x3f\x23\xc4\x7f\x9e\x89\x0c\x5c\x14\x5b\x6b\x98\x60\x7c\x3f\x06\xde\x6a\x90\x6c\xaa\x66\xcf\x95\x40\xb3\xd4\xf0\x22\x50\xf3\x98\xc1\x4f\xb7\xe2\xec\x54\x57\x73\x27\x94\x69\x17\x49\xe5\xf9\x1c\x47\x0b\x57\xf6\xf5\x71\x25\x5e\x19\xad\x6a\x0b\x61\xa4\x07\x4c\xa9\xb4\x45\xb7\xc1\xeb\x66\xf3\x8c\xe1\xfd\xb5\x94\x9b\xde\x62\x3e\xd6\x40\x4d\x85\xc2\x76\x6c\x14\x25\xb4\x65\xad\x3a\x73\x80\x2d\x84\x4c\xfb\x28\x45\x60\xde\x5b\x91\x29\xa1\x74\x6b\xff\x16\xac\x02\xde\xed\x0c\x68\x35\x0f\xbc\x6d\x6e\xf8\x40\x5f\x38\x9b\xb8\xc4\xa6\x15\x1f\xa0\x99\x3e\x6c\x40\x1b\xd1\x6c\x46\x08\xd6\x51\xc6\x35\x30\xbf\xc2\x81\x54\x3e\xd3\x7f\xb7\xe8\x27\xe4\x71\x69\x93\xf4\x6a\xe0\x75\x29\x69\x85\x3d\x4c\xad\xb7\x57\x9f\xbd\x8b\xd5\xf7\xcf\x00\x00\x00\xff\xff\x43\x92\x4a\xed\x20\x06\x00\x00")

func resources_ddl_postgres_create_table_stop_times_full_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_stop_times_full_sql,
		"resources/ddl/postgres/create-table-stop_times_full.sql",
	)
}

func resources_ddl_postgres_create_table_stop_times_full_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_stop_times_full_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-stop_times_full.sql", size: 1568, mode: os.FileMode(420), modTime: time.Unix(1429223652, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xc1\x4a\x03\x31\x10\x86\xef\xfb\x14\x73\x11\x76\x41\xfa\x02\x9e\xa2\xa6\x22\xa6\x5d\x59\xd2\x43\x4f\x4b\xc8\x0e\xdb\x85\x98\x2c\xc9\x78\xd0\xd2\x77\xb7\xd1\x5a\x70\x4d\x72\x0b\x7c\xff\xc7\xcc\x9f\xe4\xa1\xe3\x4c\x72\x90\xec\x5e\x70\xb8\x09\xab\x40\x6e\x0e\x50\x57\x00\xf1\xd4\x4f\x03\x4c\x96\x70\x44\x0f\xdb\x56\xc2\x76\x27\xc4\xed\x2f\xd3\x6e\x40\xd0\x07\xe5\xeb\xe3\xd1\xa0\x1d\xe9\x00\x3f\xfa\xea\x8a\x4f\xa7\x06\x1e\xf9\x9a\xed\xc4\xc2\xb5\xea\xad\xe4\x46\x9c\x75\x07\x0c\xba\xe0\x46\x9c\x75\x8d\x22\x58\x8b\x96\xc9\x0c\x76\xb6\x84\x47\x74\xf0\xc4\xdb\x0d\x97\xdd\xfe\x5f\xe2\xd3\x59\x8c\xf7\x95\xdc\xec\x02\xb3\x7b\xbd\x7b\x53\xa8\x74\xa6\x29\xd3\x38\xad\x68\x72\xb6\xa7\x8f\x19\xaf\xef\xb4\x4c\xcd\xca\xa3\xa5\x3e\xd0\x77\x36\x3d\xe5\x6f\x26\x35\xeb\xb5\x7b\xde\xb0\x73\xeb\x17\xbe\x87\xfa\xf2\x35\x9a\xaa\xb9\xab\xbe\x02\x00\x00\xff\xff\x67\xae\x43\x03\x41\x02\x00\x00")

func resources_ddl_postgres_create_table_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_stops_sql,
		"resources/ddl/postgres/create-table-stops.sql",
	)
}

func resources_ddl_postgres_create_table_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-stops.sql", size: 577, mode: os.FileMode(420), modTime: time.Unix(1429223652, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_transfers_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xcf\x41\x6b\x83\x30\x18\xc6\xf1\xbb\x9f\xe2\xb9\x0c\x15\x8c\xec\xbe\x53\xb6\x65\x30\x96\xb9\x21\xf1\xe0\x29\xc8\x16\x35\x50\xa3\x24\xb9\x14\xf1\xbb\xb7\x96\xb6\xb1\xf4\xd0\xfb\xef\xff\xf2\xbc\x6f\x25\xa3\x82\x41\xd0\x57\xce\xf0\xe4\x72\x6f\x1b\xe3\x5a\x65\x1d\x92\x08\x68\xed\x38\x48\xe7\xc7\x49\xea\x7f\xfc\xf5\x8d\x4d\xe6\x79\xa7\x4c\xe7\x7b\x04\x99\x6f\xd5\xb2\xa4\x28\x7e\x04\x8a\x8a\xf3\xec\x78\xc1\x8f\x8f\xfb\x60\xee\xea\xb3\x91\x7e\x3f\x29\x68\xe3\x55\xa7\xec\x55\xe0\x9d\x7d\xd0\x8a\x0b\xc4\xcf\xf1\xaa\x07\x6d\x64\x28\xf4\x10\x8a\x0b\x3c\x55\x84\xac\x98\x10\xfc\x96\x9f\xdf\xb4\xac\xf1\xc5\x6a\x24\xdb\x27\xb2\xb0\x28\xbb\x99\x90\x46\xe9\x4b\x74\x08\x00\x00\xff\xff\x88\x3d\x5d\x28\x33\x01\x00\x00")

func resources_ddl_postgres_create_table_transfers_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_transfers_sql,
		"resources/ddl/postgres/create-table-transfers.sql",
	)
}

func resources_ddl_postgres_create_table_transfers_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_transfers_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-transfers.sql", size: 307, mode: os.FileMode(420), modTime: time.Unix(1429220662, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_create_table_trips_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xcd\xaa\x83\x30\x10\x46\xf7\x79\x8a\xd9\x5c\x50\xb8\xf8\x02\x77\x95\xdb\xa6\x50\x9a\xda\x22\xba\x70\x55\x6c\x1c\xcc\x50\x89\x92\xa4\xdd\x88\xef\xde\x4a\x2b\xfd\x41\xbb\x3e\xdf\x9c\x03\xb3\x48\x04\x4f\x05\xa4\xfc\x5f\x0a\xf8\x71\x91\xb7\xd4\x3a\x08\x18\x80\x6d\xce\x1e\x0f\x54\x02\x19\x8f\x15\x5a\x88\x77\x29\xc4\x99\x94\xbf\x37\xe8\xd0\x5e\x48\xcd\xe2\xc1\xf2\x95\x69\x2c\x4a\x47\x95\x01\xa5\x0b\x1b\x74\x5d\x8d\xa6\xf2\x1a\xee\xf9\xe8\x6d\xd2\xf7\x21\x2c\xc5\x8a\x67\xf2\xe9\x28\xc9\xa2\xf2\xd4\x98\xd7\xc8\xe7\xe8\x58\x37\xea\x34\x0c\x26\x1b\x23\x9d\xd2\x3b\x5d\xb4\x38\x7b\x39\xd2\xa9\xcb\x7d\xb2\xde\xf2\x24\x87\x8d\xc8\x21\x78\x7c\x21\x64\xe1\x1f\x63\xd7\x00\x00\x00\xff\xff\xbf\x2d\x25\x17\x69\x01\x00\x00")

func resources_ddl_postgres_create_table_trips_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_create_table_trips_sql,
		"resources/ddl/postgres/create-table-trips.sql",
	)
}

func resources_ddl_postgres_create_table_trips_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_create_table_trips_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/create-table-trips.sql", size: 361, mode: os.FileMode(420), modTime: time.Unix(1429220662, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_delete_agency_by_key_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x71\xf5\x71\x0d\x71\x55\x70\x0b\xf2\xf7\x55\x48\x2f\x49\x2b\xd6\x4b\x4c\x4f\xcd\x4b\xce\x4c\x2d\x56\x08\xf7\x70\x0d\x72\x55\x00\x73\x2b\xe3\xb3\x53\x2b\x6d\x55\x0c\xad\x01\x01\x00\x00\xff\xff\x4a\x06\x5e\xfa\x2e\x00\x00\x00")

func resources_ddl_postgres_delete_agency_by_key_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_delete_agency_by_key_sql,
		"resources/ddl/postgres/delete-agency-by-key.sql",
	)
}

func resources_ddl_postgres_delete_agency_by_key_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_delete_agency_by_key_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/delete-agency-by-key.sql", size: 46, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_drop_table_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x50\x2d\xd6\x53\x2d\x06\x04\x00\x00\xff\xff\xd3\xf1\x57\xf5\x1a\x00\x00\x00")

func resources_ddl_postgres_drop_table_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_drop_table_sql,
		"resources/ddl/postgres/drop-table.sql",
	)
}

func resources_ddl_postgres_drop_table_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_drop_table_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/drop-table.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_insert_line_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd0\x31\x6e\xc5\x30\x08\x06\xe0\xf9\xf9\x14\x2c\x91\x5a\xa9\xf2\x05\xaa\x9c\x25\x6a\x63\xaa\x52\x39\x50\x01\x1d\x7a\xfb\x67\x3b\x4e\xbc\x38\xd3\x4f\xc8\x07\x0a\xc4\x86\xea\x40\xec\x02\x8b\xc5\x4c\x8c\x9b\xb9\xfc\x1a\xbc\xb4\x4c\xe9\x0d\x6a\x3d\xc2\x2e\x09\x5f\xc1\x30\xe3\xee\x90\xc8\x9c\x78\xf7\xf0\x80\x1c\x07\x88\x83\xc4\x1b\x85\x2f\x95\x23\x3c\x96\xf3\x95\x81\x95\xad\x8c\x0a\x3f\x42\x1c\xa0\x3c\xbd\xb3\x39\x1d\x58\xda\x0e\xc2\x63\xd4\x6a\x7e\x45\x98\x40\x57\x2a\x23\x1b\xf1\x56\x74\xd2\xe3\x8c\xa8\xfc\x79\xd9\xa3\xa7\x69\x55\x45\x7a\xc7\x89\xa9\xbf\x68\x90\x2b\xb9\xbe\xb3\x6f\x51\xdf\xf8\xe3\xc0\xb5\x9f\xa0\xe6\x20\x9a\x8a\xfd\xfc\x6f\x72\x76\x9b\xf7\xf0\x0c\x00\x00\xff\xff\x27\xc4\xae\x94\x7c\x01\x00\x00")

func resources_ddl_postgres_insert_line_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_insert_line_stops_sql,
		"resources/ddl/postgres/insert-line_stops.sql",
	)
}

func resources_ddl_postgres_insert_line_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_insert_line_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/insert-line_stops.sql", size: 380, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_insert_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcc\x31\x0e\xc2\x40\x0c\x44\xd1\x7e\x4f\x31\x0d\x12\x34\xb9\x00\x87\x89\x42\x62\x84\xa5\xc4\x96\x66\x4c\xc1\xed\x59\xb6\xa0\x8b\x1b\x37\x6f\xbe\x87\x8c\x05\x8f\x4a\x5c\x34\xed\x1e\x26\x5c\x7f\x6f\x8e\xe5\xb0\x1b\x64\xbb\xad\x85\xcd\x55\x1e\x6b\x35\xf4\xe3\xc4\x7c\x97\xcd\x7a\x25\x6b\x38\x2c\xc2\x7f\xd4\x9e\xcc\x63\xc0\x5e\x1c\x52\x60\x4b\x6e\x46\x3c\x3e\x27\x85\x7b\xfb\x06\x00\x00\xff\xff\x3f\xa3\xde\xa9\x8b\x00\x00\x00")

func resources_ddl_postgres_insert_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_insert_lines_sql,
		"resources/ddl/postgres/insert-lines.sql",
	)
}

func resources_ddl_postgres_insert_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_insert_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/insert-lines.sql", size: 139, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_insert_route_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\x41\x0a\xc3\x30\x0c\x04\xef\x79\xc5\x5e\x02\x2d\x14\x7f\xa0\xe4\x2d\x39\xc4\x2a\xa8\x34\x52\x91\xd4\xff\xd7\x8e\x49\x1c\x68\x7d\x30\x92\xd6\xb3\xd6\xb2\x38\x59\x80\x25\x14\xa3\x27\xd3\x4f\xd0\xec\xa1\x6f\xc7\xa5\x35\x9c\x6f\xa8\x83\x5e\x2c\x9a\xe9\x0a\xa7\x17\x2d\x81\xcc\x1e\x2c\x4b\x0c\x28\xc7\xd2\x89\x49\x9d\x4a\x07\x37\x3c\x4c\xd7\xed\xed\xd8\xa6\x0e\x2f\xbf\x0b\x19\x9e\xca\x72\x56\xe6\xe0\x95\x8a\x1c\x50\x29\xf7\x6e\x87\xa9\x5b\xff\x21\xc3\xb8\x78\xee\x4c\xed\x1a\xd3\xeb\x5f\x66\x5b\xda\x61\x15\x8a\x23\xc2\xd4\xd3\xdc\x87\x6f\x00\x00\x00\xff\xff\xf3\x1e\x03\xc1\x28\x01\x00\x00")

func resources_ddl_postgres_insert_route_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_insert_route_stops_sql,
		"resources/ddl/postgres/insert-route_stops.sql",
	)
}

func resources_ddl_postgres_insert_route_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_insert_route_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/insert-route_stops.sql", size: 296, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_insert_station_lines_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8f\xd1\xca\xc2\x30\x0c\x85\xef\xfb\x14\xe7\xe6\x87\x5f\x90\xbd\x80\xec\x59\x86\x6c\x15\x22\x35\x19\x3d\x79\x7f\x6c\x3b\x67\x55\xb4\x57\x49\xfb\x7d\x27\x8d\x28\x63\x76\x88\xba\xe1\x8f\x03\xfd\xec\x62\x3a\x25\xd1\x48\xfc\xef\xad\x2c\x47\xd4\xab\x52\x1c\xc0\x98\xe2\xec\x58\x84\x2e\x3a\x7b\x40\x39\xdd\x6c\x28\x87\x07\x1d\x2e\xd9\x6e\x8d\xe8\xe1\x04\xcb\x3c\x8d\x19\x57\x13\xfd\x78\x9c\xe8\xb6\x16\x82\x30\x7d\x4b\x1d\xf9\xda\x7d\x09\x68\x13\x37\x3b\x6d\x76\x15\x6c\xad\x6e\x7a\x96\x3f\xc4\xe2\x54\x25\xed\xff\x1e\xfb\x0a\xa7\x70\x0f\x00\x00\xff\xff\x58\x2d\xae\x1d\x24\x01\x00\x00")

func resources_ddl_postgres_insert_station_lines_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_insert_station_lines_sql,
		"resources/ddl/postgres/insert-station_lines.sql",
	)
}

func resources_ddl_postgres_insert_station_lines_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_insert_station_lines_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/insert-station_lines.sql", size: 292, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_insert_station_stops_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8d\x51\x0a\xc2\x40\x0c\x44\xff\xf7\x14\xf3\x53\x50\x90\xbd\x80\x78\x16\x29\x6d\x84\x88\x4d\x64\x67\xee\x8f\xdb\x6a\x51\xf7\x6b\xf3\xe6\x4d\xe2\x41\x6b\x82\x87\x12\x03\x2b\x35\xca\x33\xae\x54\x3e\x89\xc3\x3e\xfa\x7c\xc2\x8a\xfa\xe7\x08\xda\xc3\x26\x61\x76\xca\x63\x52\x41\x7f\x54\xfd\x73\xeb\xc7\x2e\xb7\x96\xcb\x66\x0c\x6f\x46\x10\xfd\x5a\x58\xc3\x3d\x3d\xbe\xd1\x56\xee\xa9\x90\xb1\xf7\x63\x5c\x0c\x97\xdf\xe5\x2b\x39\x97\x57\x00\x00\x00\xff\xff\x3c\x34\xb5\x60\xb5\x00\x00\x00")

func resources_ddl_postgres_insert_station_stops_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_insert_station_stops_sql,
		"resources/ddl/postgres/insert-station_stops.sql",
	)
}

func resources_ddl_postgres_insert_station_stops_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_insert_station_stops_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/insert-station_stops.sql", size: 181, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_insert_stations_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8e\x4d\x0a\xc2\x30\x10\x85\xf7\x39\xc5\xdb\x94\x24\x50\xbc\x80\x27\x70\xa3\x2e\xdc\x4b\x8c\xb1\x04\xda\x99\xd2\x19\x05\x6f\x6f\x7f\x22\x16\x3b\xab\x79\xdf\xf0\x3e\x26\x93\xa4\x41\x91\x49\x19\x95\xec\x44\x83\x66\x26\x81\x2b\xdb\x95\x42\x97\x6a\x7c\x53\x1b\x74\x15\x98\x7e\xa1\x49\xec\x21\xa9\x4d\x51\x71\xcf\xa2\x99\xa2\x1a\x8c\x33\x49\xb9\x5f\x3c\x33\x08\xaf\xc6\x15\x38\xea\xfc\x16\x32\x15\x18\x99\x62\x50\x67\xcf\xa7\xc3\xf1\xe2\x6c\xbd\xa9\xc2\xe2\x8f\x4e\x5d\x58\x6f\xbd\x79\x0c\xdc\xcd\x96\x6a\xb9\x09\xc4\x34\x03\x3f\x7b\xdc\xde\xeb\xa7\xf6\xe6\x13\x00\x00\xff\xff\x14\x5f\x59\x8e\x04\x01\x00\x00")

func resources_ddl_postgres_insert_stations_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_insert_stations_sql,
		"resources/ddl/postgres/insert-stations.sql",
	)
}

func resources_ddl_postgres_insert_stations_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_insert_stations_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/insert-stations.sql", size: 260, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_insert_stop_times_full_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\x4b\x96\xa3\x30\x0c\x5c\xc3\x29\xbc\xc9\xcb\x26\xe1\x06\x39\x0b\xcf\x63\x0b\xd0\xc4\xb1\x18\x49\x64\x3a\x7d\xfa\x36\xbf\x80\xd3\xb0\xc2\x25\x55\xd9\x2a\x15\x46\x01\x56\x83\x51\xc9\x9c\xa4\x12\xa5\xbe\x56\x7c\x80\xd4\xcd\x10\x42\x29\x10\xc0\xa9\x29\x4d\xfa\x96\x2a\xfa\x4b\x79\xbd\xee\x00\x47\x1e\x2e\xfb\x8e\x68\x1f\x39\xe0\x41\x5c\x06\x04\xab\xf9\x99\xe2\x87\x68\x0b\xb4\x21\xdf\x14\xe1\xf7\xbd\x03\x87\x55\x24\x90\xb3\x8a\x14\x6b\x7d\xf5\xb0\xb5\xf5\x96\x21\x6a\x2d\x3a\x15\x2f\x65\x21\x5a\x59\x66\x7c\xda\x30\x0d\xb9\xd0\xb5\xf2\x90\x5a\x75\x60\xc8\xe1\xe9\x1a\x81\x7f\x03\x44\xb7\xc9\x2e\x78\x07\xd6\xd7\x82\x6d\xdc\x15\x7a\x74\xf7\xa1\xcf\x5f\x91\xd4\x39\xb5\x53\xd3\x64\x38\x57\xb6\x4d\xb2\xaf\xdd\x5c\x5c\x31\x0d\x3a\x4f\x5a\xac\x07\xe9\x88\x75\xb1\x34\x6f\x4b\xa6\xb5\x87\x85\xcd\xee\x15\xc9\x2f\x9e\xb1\xb7\x7d\x2b\xe0\x28\x10\x7f\xf0\xe0\x4b\xf7\xb8\x56\xca\x38\x47\x60\x3e\xa6\xf0\x3c\xd1\xad\xcb\x29\x96\xfa\x68\xcd\xec\xcc\xdc\xe5\x91\x53\x8a\xc6\xfd\xa0\x5f\x5e\xa1\xd5\x9f\xb4\xb4\xfb\x6e\xfa\x24\xd6\xd9\x7e\x94\x2a\x1b\xa6\x47\x59\x2c\x71\x14\x23\x29\x9e\x11\xd8\xfc\x25\x8c\x93\x62\x16\xd4\xe4\xb0\xa1\xb8\xa5\xf3\xb6\x2e\x08\xfd\x01\x6f\x7c\x9f\x98\x99\xf1\x9e\xe6\xf6\xfe\x3b\x60\x4c\x4e\x88\xe1\x91\xb2\x6d\x28\x51\xd6\xdf\xf2\x7f\x07\x0c\x47\x1b\xbb\x9d\x4f\xcf\xf3\x4f\x00\x00\x00\xff\xff\x2e\xb4\x8f\x01\x62\x03\x00\x00")

func resources_ddl_postgres_insert_stop_times_full_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_insert_stop_times_full_sql,
		"resources/ddl/postgres/insert-stop_times_full.sql",
	)
}

func resources_ddl_postgres_insert_stop_times_full_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_insert_stop_times_full_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/insert-stop_times_full.sql", size: 866, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_select_agency_zone_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4e\xcd\x49\x4d\x2e\x51\xc8\xcd\xcc\xd3\x28\xd6\x2b\x2e\xc9\x2f\x88\xcf\x49\x2c\xd1\xd4\x51\xc8\x4d\xac\x40\x13\x40\x52\x91\x9f\x87\xa6\x02\x28\xa0\x90\x56\x94\x9f\xab\xa0\x0a\x11\x29\x56\x28\x06\x04\x00\x00\xff\xff\x5a\x01\xd4\x61\x59\x00\x00\x00")

func resources_ddl_postgres_select_agency_zone_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_select_agency_zone_sql,
		"resources/ddl/postgres/select-agency-zone.sql",
	)
}

func resources_ddl_postgres_select_agency_zone_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_select_agency_zone_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/select-agency-zone.sql", size: 89, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_select_trip_stop_times_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xce\x41\x0e\xc2\x30\x0c\x04\xc0\xaf\xf8\x52\xf5\x82\xf8\x01\x6f\xa9\x42\xb3\x08\xa3\xc6\x2e\xb6\x0b\xe2\xf7\x24\x40\xe0\xc0\xcd\x9a\xd5\xae\xec\x58\x30\x07\x79\xec\x93\x19\xdf\xd2\x32\x05\x17\xec\x1a\x64\xac\xc9\x62\x33\xfc\xc8\x43\xd7\xc9\x71\xdd\x20\x73\x93\x37\x48\x2a\xa0\x93\x69\xa1\xe1\x23\xad\xe0\xb5\x40\x2c\x02\xa3\x8b\xb2\xf4\xac\x32\xa9\x7c\xc7\x38\x1f\xbc\x5f\x74\x3f\xc3\xd0\xa2\x30\x7e\x45\xe3\xe0\x23\xa9\xe5\xba\x71\x7c\xfc\x3d\xf0\x0c\x00\x00\xff\xff\x13\x33\x4a\xe5\xbc\x00\x00\x00")

func resources_ddl_postgres_select_trip_stop_times_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_select_trip_stop_times_sql,
		"resources/ddl/postgres/select_trip_stop_times.sql",
	)
}

func resources_ddl_postgres_select_trip_stop_times_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_select_trip_stop_times_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/select_trip_stop_times.sql", size: 188, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_select_trips_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4e\xcd\x49\x4d\x2e\x51\x28\x29\xca\x2c\x88\xcf\x4c\x51\x48\x2b\xca\xcf\x55\x50\x2d\xd6\x03\xf1\x8b\x15\xf2\x8b\x52\x52\x8b\x14\x92\x2a\x61\xd2\x80\x00\x00\x00\xff\xff\x47\xab\x60\x69\x2d\x00\x00\x00")

func resources_ddl_postgres_select_trips_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_select_trips_sql,
		"resources/ddl/postgres/select_trips.sql",
	)
}

func resources_ddl_postgres_select_trips_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_select_trips_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/select_trips.sql", size: 45, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_update_agency_zone_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x2d\x48\x49\x2c\x49\x55\x50\x2d\xd6\x4b\x4c\x4f\xcd\x4b\xce\x4c\x2d\x56\x28\x4e\x2d\x51\x00\x73\x2a\xe3\x73\x33\xf3\xe2\x73\x12\x4b\x6c\x55\x0c\x75\xe0\x42\x89\x15\x10\x21\x23\x1d\x14\x55\xf9\x79\xb6\x2a\xc6\xa8\xaa\x40\x42\x26\x5c\x80\x00\x00\x00\xff\xff\x4e\x9b\x20\x4b\x62\x00\x00\x00")

func resources_ddl_postgres_update_agency_zone_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_update_agency_zone_sql,
		"resources/ddl/postgres/update-agency-zone.sql",
	)
}

func resources_ddl_postgres_update_agency_zone_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_update_agency_zone_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/update-agency-zone.sql", size: 98, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resources_ddl_postgres_update_gtfs_agency_zone_sql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x2d\x48\x49\x2c\x49\x55\x48\x2f\x49\x2b\xd6\x4b\x4c\x4f\xcd\x4b\xce\x4c\x2d\x56\x28\x4e\x2d\x51\x00\x73\x2a\xe3\x73\x33\xf3\xe2\x73\x12\x4b\x6c\x55\x0c\x75\xe0\x42\x89\x15\x10\x21\x23\x1d\x14\x55\xf9\x79\xb6\x2a\xc6\xa8\xaa\x40\x42\x26\x0a\xe5\x19\xa9\x45\xa9\x30\xf1\xec\xd4\x4a\x5b\x15\x53\x2e\x40\x00\x00\x00\xff\xff\x87\x0f\xe5\xb0\x78\x00\x00\x00")

func resources_ddl_postgres_update_gtfs_agency_zone_sql_bytes() ([]byte, error) {
	return bindata_read(
		_resources_ddl_postgres_update_gtfs_agency_zone_sql,
		"resources/ddl/postgres/update-gtfs-agency-zone.sql",
	)
}

func resources_ddl_postgres_update_gtfs_agency_zone_sql() (*asset, error) {
	bytes, err := resources_ddl_postgres_update_gtfs_agency_zone_sql_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "resources/ddl/postgres/update-gtfs-agency-zone.sql", size: 120, mode: os.FileMode(420), modTime: time.Unix(1428878138, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"resources/ddl/mysql/create-index.sql": resources_ddl_mysql_create_index_sql,
	"resources/ddl/mysql/create-schema.sql": resources_ddl_mysql_create_schema_sql,
	"resources/ddl/mysql/create-spatial-index.sql": resources_ddl_mysql_create_spatial_index_sql,
	"resources/ddl/mysql/create-table-agencies.sql": resources_ddl_mysql_create_table_agencies_sql,
	"resources/ddl/mysql/create-table-calendar_dates.sql": resources_ddl_mysql_create_table_calendar_dates_sql,
	"resources/ddl/mysql/create-table-calendars.sql": resources_ddl_mysql_create_table_calendars_sql,
	"resources/ddl/mysql/create-table-line_stops.sql": resources_ddl_mysql_create_table_line_stops_sql,
	"resources/ddl/mysql/create-table-lines.sql": resources_ddl_mysql_create_table_lines_sql,
	"resources/ddl/mysql/create-table-route_stops.sql": resources_ddl_mysql_create_table_route_stops_sql,
	"resources/ddl/mysql/create-table-routes.sql": resources_ddl_mysql_create_table_routes_sql,
	"resources/ddl/mysql/create-table-station_lines.sql": resources_ddl_mysql_create_table_station_lines_sql,
	"resources/ddl/mysql/create-table-station_stops.sql": resources_ddl_mysql_create_table_station_stops_sql,
	"resources/ddl/mysql/create-table-stations.sql": resources_ddl_mysql_create_table_stations_sql,
	"resources/ddl/mysql/create-table-stop_times.sql": resources_ddl_mysql_create_table_stop_times_sql,
	"resources/ddl/mysql/create-table-stop_times_full.sql": resources_ddl_mysql_create_table_stop_times_full_sql,
	"resources/ddl/mysql/create-table-stops.sql": resources_ddl_mysql_create_table_stops_sql,
	"resources/ddl/mysql/create-table-transfers.sql": resources_ddl_mysql_create_table_transfers_sql,
	"resources/ddl/mysql/create-table-trips.sql": resources_ddl_mysql_create_table_trips_sql,
	"resources/ddl/mysql/delete-agency-by-key.sql": resources_ddl_mysql_delete_agency_by_key_sql,
	"resources/ddl/mysql/drop-table.sql": resources_ddl_mysql_drop_table_sql,
	"resources/ddl/mysql/insert-line_stops.sql": resources_ddl_mysql_insert_line_stops_sql,
	"resources/ddl/mysql/insert-lines.sql": resources_ddl_mysql_insert_lines_sql,
	"resources/ddl/mysql/insert-route_stops.sql": resources_ddl_mysql_insert_route_stops_sql,
	"resources/ddl/mysql/insert-station_lines.sql": resources_ddl_mysql_insert_station_lines_sql,
	"resources/ddl/mysql/insert-station_stops.sql": resources_ddl_mysql_insert_station_stops_sql,
	"resources/ddl/mysql/insert-stations.sql": resources_ddl_mysql_insert_stations_sql,
	"resources/ddl/mysql/insert-stop_times_full.sql": resources_ddl_mysql_insert_stop_times_full_sql,
	"resources/ddl/mysql/select-agency-zone.sql": resources_ddl_mysql_select_agency_zone_sql,
	"resources/ddl/mysql/select_trip_stop_times.sql": resources_ddl_mysql_select_trip_stop_times_sql,
	"resources/ddl/mysql/select_trips.sql": resources_ddl_mysql_select_trips_sql,
	"resources/ddl/mysql/update-agency-zone.sql": resources_ddl_mysql_update_agency_zone_sql,
	"resources/ddl/mysql/update-gtfs-agency-zone.sql": resources_ddl_mysql_update_gtfs_agency_zone_sql,
	"resources/ddl/postgres/create-index.sql": resources_ddl_postgres_create_index_sql,
	"resources/ddl/postgres/create-schema.sql": resources_ddl_postgres_create_schema_sql,
	"resources/ddl/postgres/create-spatial-index.sql": resources_ddl_postgres_create_spatial_index_sql,
	"resources/ddl/postgres/create-table-agencies.sql": resources_ddl_postgres_create_table_agencies_sql,
	"resources/ddl/postgres/create-table-calendar_dates.sql": resources_ddl_postgres_create_table_calendar_dates_sql,
	"resources/ddl/postgres/create-table-calendars.sql": resources_ddl_postgres_create_table_calendars_sql,
	"resources/ddl/postgres/create-table-line_stops.sql": resources_ddl_postgres_create_table_line_stops_sql,
	"resources/ddl/postgres/create-table-lines.sql": resources_ddl_postgres_create_table_lines_sql,
	"resources/ddl/postgres/create-table-route_stops.sql": resources_ddl_postgres_create_table_route_stops_sql,
	"resources/ddl/postgres/create-table-routes.sql": resources_ddl_postgres_create_table_routes_sql,
	"resources/ddl/postgres/create-table-station_lines.sql": resources_ddl_postgres_create_table_station_lines_sql,
	"resources/ddl/postgres/create-table-station_stops.sql": resources_ddl_postgres_create_table_station_stops_sql,
	"resources/ddl/postgres/create-table-stations.sql": resources_ddl_postgres_create_table_stations_sql,
	"resources/ddl/postgres/create-table-stop_times.sql": resources_ddl_postgres_create_table_stop_times_sql,
	"resources/ddl/postgres/create-table-stop_times_full.sql": resources_ddl_postgres_create_table_stop_times_full_sql,
	"resources/ddl/postgres/create-table-stops.sql": resources_ddl_postgres_create_table_stops_sql,
	"resources/ddl/postgres/create-table-transfers.sql": resources_ddl_postgres_create_table_transfers_sql,
	"resources/ddl/postgres/create-table-trips.sql": resources_ddl_postgres_create_table_trips_sql,
	"resources/ddl/postgres/delete-agency-by-key.sql": resources_ddl_postgres_delete_agency_by_key_sql,
	"resources/ddl/postgres/drop-table.sql": resources_ddl_postgres_drop_table_sql,
	"resources/ddl/postgres/insert-line_stops.sql": resources_ddl_postgres_insert_line_stops_sql,
	"resources/ddl/postgres/insert-lines.sql": resources_ddl_postgres_insert_lines_sql,
	"resources/ddl/postgres/insert-route_stops.sql": resources_ddl_postgres_insert_route_stops_sql,
	"resources/ddl/postgres/insert-station_lines.sql": resources_ddl_postgres_insert_station_lines_sql,
	"resources/ddl/postgres/insert-station_stops.sql": resources_ddl_postgres_insert_station_stops_sql,
	"resources/ddl/postgres/insert-stations.sql": resources_ddl_postgres_insert_stations_sql,
	"resources/ddl/postgres/insert-stop_times_full.sql": resources_ddl_postgres_insert_stop_times_full_sql,
	"resources/ddl/postgres/select-agency-zone.sql": resources_ddl_postgres_select_agency_zone_sql,
	"resources/ddl/postgres/select_trip_stop_times.sql": resources_ddl_postgres_select_trip_stop_times_sql,
	"resources/ddl/postgres/select_trips.sql": resources_ddl_postgres_select_trips_sql,
	"resources/ddl/postgres/update-agency-zone.sql": resources_ddl_postgres_update_agency_zone_sql,
	"resources/ddl/postgres/update-gtfs-agency-zone.sql": resources_ddl_postgres_update_gtfs_agency_zone_sql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resources": &_bintree_t{nil, map[string]*_bintree_t{
		"ddl": &_bintree_t{nil, map[string]*_bintree_t{
			"mysql": &_bintree_t{nil, map[string]*_bintree_t{
				"create-index.sql": &_bintree_t{resources_ddl_mysql_create_index_sql, map[string]*_bintree_t{
				}},
				"create-schema.sql": &_bintree_t{resources_ddl_mysql_create_schema_sql, map[string]*_bintree_t{
				}},
				"create-spatial-index.sql": &_bintree_t{resources_ddl_mysql_create_spatial_index_sql, map[string]*_bintree_t{
				}},
				"create-table-agencies.sql": &_bintree_t{resources_ddl_mysql_create_table_agencies_sql, map[string]*_bintree_t{
				}},
				"create-table-calendar_dates.sql": &_bintree_t{resources_ddl_mysql_create_table_calendar_dates_sql, map[string]*_bintree_t{
				}},
				"create-table-calendars.sql": &_bintree_t{resources_ddl_mysql_create_table_calendars_sql, map[string]*_bintree_t{
				}},
				"create-table-line_stops.sql": &_bintree_t{resources_ddl_mysql_create_table_line_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-lines.sql": &_bintree_t{resources_ddl_mysql_create_table_lines_sql, map[string]*_bintree_t{
				}},
				"create-table-route_stops.sql": &_bintree_t{resources_ddl_mysql_create_table_route_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-routes.sql": &_bintree_t{resources_ddl_mysql_create_table_routes_sql, map[string]*_bintree_t{
				}},
				"create-table-station_lines.sql": &_bintree_t{resources_ddl_mysql_create_table_station_lines_sql, map[string]*_bintree_t{
				}},
				"create-table-station_stops.sql": &_bintree_t{resources_ddl_mysql_create_table_station_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-stations.sql": &_bintree_t{resources_ddl_mysql_create_table_stations_sql, map[string]*_bintree_t{
				}},
				"create-table-stop_times.sql": &_bintree_t{resources_ddl_mysql_create_table_stop_times_sql, map[string]*_bintree_t{
				}},
				"create-table-stop_times_full.sql": &_bintree_t{resources_ddl_mysql_create_table_stop_times_full_sql, map[string]*_bintree_t{
				}},
				"create-table-stops.sql": &_bintree_t{resources_ddl_mysql_create_table_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-transfers.sql": &_bintree_t{resources_ddl_mysql_create_table_transfers_sql, map[string]*_bintree_t{
				}},
				"create-table-trips.sql": &_bintree_t{resources_ddl_mysql_create_table_trips_sql, map[string]*_bintree_t{
				}},
				"delete-agency-by-key.sql": &_bintree_t{resources_ddl_mysql_delete_agency_by_key_sql, map[string]*_bintree_t{
				}},
				"drop-table.sql": &_bintree_t{resources_ddl_mysql_drop_table_sql, map[string]*_bintree_t{
				}},
				"insert-line_stops.sql": &_bintree_t{resources_ddl_mysql_insert_line_stops_sql, map[string]*_bintree_t{
				}},
				"insert-lines.sql": &_bintree_t{resources_ddl_mysql_insert_lines_sql, map[string]*_bintree_t{
				}},
				"insert-route_stops.sql": &_bintree_t{resources_ddl_mysql_insert_route_stops_sql, map[string]*_bintree_t{
				}},
				"insert-station_lines.sql": &_bintree_t{resources_ddl_mysql_insert_station_lines_sql, map[string]*_bintree_t{
				}},
				"insert-station_stops.sql": &_bintree_t{resources_ddl_mysql_insert_station_stops_sql, map[string]*_bintree_t{
				}},
				"insert-stations.sql": &_bintree_t{resources_ddl_mysql_insert_stations_sql, map[string]*_bintree_t{
				}},
				"insert-stop_times_full.sql": &_bintree_t{resources_ddl_mysql_insert_stop_times_full_sql, map[string]*_bintree_t{
				}},
				"select-agency-zone.sql": &_bintree_t{resources_ddl_mysql_select_agency_zone_sql, map[string]*_bintree_t{
				}},
				"select_trip_stop_times.sql": &_bintree_t{resources_ddl_mysql_select_trip_stop_times_sql, map[string]*_bintree_t{
				}},
				"select_trips.sql": &_bintree_t{resources_ddl_mysql_select_trips_sql, map[string]*_bintree_t{
				}},
				"update-agency-zone.sql": &_bintree_t{resources_ddl_mysql_update_agency_zone_sql, map[string]*_bintree_t{
				}},
				"update-gtfs-agency-zone.sql": &_bintree_t{resources_ddl_mysql_update_gtfs_agency_zone_sql, map[string]*_bintree_t{
				}},
			}},
			"postgres": &_bintree_t{nil, map[string]*_bintree_t{
				"create-index.sql": &_bintree_t{resources_ddl_postgres_create_index_sql, map[string]*_bintree_t{
				}},
				"create-schema.sql": &_bintree_t{resources_ddl_postgres_create_schema_sql, map[string]*_bintree_t{
				}},
				"create-spatial-index.sql": &_bintree_t{resources_ddl_postgres_create_spatial_index_sql, map[string]*_bintree_t{
				}},
				"create-table-agencies.sql": &_bintree_t{resources_ddl_postgres_create_table_agencies_sql, map[string]*_bintree_t{
				}},
				"create-table-calendar_dates.sql": &_bintree_t{resources_ddl_postgres_create_table_calendar_dates_sql, map[string]*_bintree_t{
				}},
				"create-table-calendars.sql": &_bintree_t{resources_ddl_postgres_create_table_calendars_sql, map[string]*_bintree_t{
				}},
				"create-table-line_stops.sql": &_bintree_t{resources_ddl_postgres_create_table_line_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-lines.sql": &_bintree_t{resources_ddl_postgres_create_table_lines_sql, map[string]*_bintree_t{
				}},
				"create-table-route_stops.sql": &_bintree_t{resources_ddl_postgres_create_table_route_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-routes.sql": &_bintree_t{resources_ddl_postgres_create_table_routes_sql, map[string]*_bintree_t{
				}},
				"create-table-station_lines.sql": &_bintree_t{resources_ddl_postgres_create_table_station_lines_sql, map[string]*_bintree_t{
				}},
				"create-table-station_stops.sql": &_bintree_t{resources_ddl_postgres_create_table_station_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-stations.sql": &_bintree_t{resources_ddl_postgres_create_table_stations_sql, map[string]*_bintree_t{
				}},
				"create-table-stop_times.sql": &_bintree_t{resources_ddl_postgres_create_table_stop_times_sql, map[string]*_bintree_t{
				}},
				"create-table-stop_times_full.sql": &_bintree_t{resources_ddl_postgres_create_table_stop_times_full_sql, map[string]*_bintree_t{
				}},
				"create-table-stops.sql": &_bintree_t{resources_ddl_postgres_create_table_stops_sql, map[string]*_bintree_t{
				}},
				"create-table-transfers.sql": &_bintree_t{resources_ddl_postgres_create_table_transfers_sql, map[string]*_bintree_t{
				}},
				"create-table-trips.sql": &_bintree_t{resources_ddl_postgres_create_table_trips_sql, map[string]*_bintree_t{
				}},
				"delete-agency-by-key.sql": &_bintree_t{resources_ddl_postgres_delete_agency_by_key_sql, map[string]*_bintree_t{
				}},
				"drop-table.sql": &_bintree_t{resources_ddl_postgres_drop_table_sql, map[string]*_bintree_t{
				}},
				"insert-line_stops.sql": &_bintree_t{resources_ddl_postgres_insert_line_stops_sql, map[string]*_bintree_t{
				}},
				"insert-lines.sql": &_bintree_t{resources_ddl_postgres_insert_lines_sql, map[string]*_bintree_t{
				}},
				"insert-route_stops.sql": &_bintree_t{resources_ddl_postgres_insert_route_stops_sql, map[string]*_bintree_t{
				}},
				"insert-station_lines.sql": &_bintree_t{resources_ddl_postgres_insert_station_lines_sql, map[string]*_bintree_t{
				}},
				"insert-station_stops.sql": &_bintree_t{resources_ddl_postgres_insert_station_stops_sql, map[string]*_bintree_t{
				}},
				"insert-stations.sql": &_bintree_t{resources_ddl_postgres_insert_stations_sql, map[string]*_bintree_t{
				}},
				"insert-stop_times_full.sql": &_bintree_t{resources_ddl_postgres_insert_stop_times_full_sql, map[string]*_bintree_t{
				}},
				"select-agency-zone.sql": &_bintree_t{resources_ddl_postgres_select_agency_zone_sql, map[string]*_bintree_t{
				}},
				"select_trip_stop_times.sql": &_bintree_t{resources_ddl_postgres_select_trip_stop_times_sql, map[string]*_bintree_t{
				}},
				"select_trips.sql": &_bintree_t{resources_ddl_postgres_select_trips_sql, map[string]*_bintree_t{
				}},
				"update-agency-zone.sql": &_bintree_t{resources_ddl_postgres_update_agency_zone_sql, map[string]*_bintree_t{
				}},
				"update-gtfs-agency-zone.sql": &_bintree_t{resources_ddl_postgres_update_gtfs_agency_zone_sql, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

