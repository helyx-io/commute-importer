gtfs-playground
===============

This project aims to be an playground to explore General Transit Feed Specification Reference by enabling import GTFS data sets into a MySQL database and exposing an API to query data.

[![Coverage Status](https://coveralls.io/repos/helyx-io/gtfs-playground/badge.png)](https://coveralls.io/r/helyx-io/gtfs-playground)


Dump Database
-------------

    mysqldump --no-data -hlocalhost -ugtfs -pgtfs gtfs > gtfs-ddl.sql
    
    

Impo
rt Database
---------------

    mysql -ugtfs -pgtfs gtfs < gtfs-ddl.sql
