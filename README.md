commute-importer
=============

This project aims to be an playground to explore General Transit Feed Specification Reference by enabling import GTFS data sets into a SQL database and exposing an API to query data.

[![Build Status](https://travis-ci.org/helyx-io/commute-importer.svg?branch=master)](https://travis-ci.org/helyx-io/commute-importer)
[![Coverage Status](https://coveralls.io/repos/helyx-io/commute-importer/badge.png)](https://coveralls.io/r/helyx-io/commute-importer)



Dump Database
-------------

    mysqldump --no-data -hlocalhost -ucommute -pcommute commute > ddl/commute-ddl.sql  
    
    

Import Database
---------------

    mysql -ucommute -pcommute commute < commute-ddl.sql
    
 

REST Resources
--------------

 - **GET** http://localhost:3000/import/ - *Run import job*



Generate static resources:
--------------------------

Install go-bindata:
    
    go get -u github.com/jteeuwen/go-bindata/...
    
Run go-bindata:

    go-bindata -o data/data.go -pkg data resources/... 

