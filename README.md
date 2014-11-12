gtfs-playground
===============

This project aims to be an playground to explore General Transit Feed Specification Reference by enabling import GTFS data sets into a MySQL database and exposing an API to query data.

[![Build Status](https://travis-ci.org/helyx-io/gtfs-playground.svg?branch=master)](https://travis-ci.org/helyx-io/gtfs-playground)
[![Coverage Status](https://coveralls.io/repos/helyx-io/gtfs-playground/badge.png)](https://coveralls.io/r/helyx-io/gtfs-playground)



Dump Database
-------------

    mysqldump --no-data -hlocalhost -ugtfs -pgtfs gtfs > gtfs-ddl.sql
    
    

Import Database
---------------

    mysql -ugtfs -pgtfs gtfs < gtfs-ddl.sql
    
 

REST Resources
--------------

 - **GET** http://localhost:3000/import/ - *Run import job*
 - **GET** http://localhost:3000/agencies/ - *Get list of agencies*

        [
          {
            "key": "RATP",
            "agencyId": "RER",
            "name": "RER",
            "url": "http://...",
            "timezone": "Europe/Paris",
            "lang": "fr"
          }, ... ,
          {
            "key": "RATP",
            "agencyId": "Noctilien",
            "name": "Noctilien",
            "url": "http://...",
            "timezone": "Europe/Paris",
            "lang": "fr"
          }
        ]

 - **GET** http://localhost:3000/agencies/85 - *Get an agency by id*

        {
          "key": "RATP",
          "agencyId": "Noctilien",
          "name": "Noctilien",
          "url": "http://www.navitia.com",
          "timezone": "Europe/Paris",
          "lang": "fr"
        }



Perf Test
---------

Install Vegeta:
    
     go get github.com/tsenart/vegeta
     go install github.com/tsenart/vegeta

Run Vegeta on Routes URL:

    echo "GET http://localhost:3000/routes/" | vegeta attack -rate=20 -duration=10s | vegeta report -reporter=plot > report.html && open report.html