gtfs-importer
===============

This project aims to be an playground to explore General Transit Feed Specification Reference by enabling import GTFS data sets into a MySQL database and exposing an API to query data.

[![Build Status](https://travis-ci.org/helyx-io/gtfs-importer.svg?branch=master)](https://travis-ci.org/helyx-io/gtfs-importer)
[![Coverage Status](https://coveralls.io/repos/helyx-io/gtfs-importer/badge.png)](https://coveralls.io/r/helyx-io/gtfs-importer)



Dump Database
-------------

    mysqldump --no-data -hlocalhost -ugtfs -pgtfs gtfs > ddl/gtfs-ddl.sql  
    
    

Import Database
---------------

    mysql -ugtfs -pgtfs gtfs < gtfs-ddl.sql
    
 

REST Resources
--------------

 - **GET** http://localhost:3000/import/ - *Run import job*



Generate static resources:
--------------------------

Install go-bindata:
    
    go get -u github.com/jteeuwen/go-bindata/...
    
Run go-bindata:

    go-bindata -o data/data.go -pkg data resources/... 

