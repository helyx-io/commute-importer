gtfs-playground
===============

This project aims to be an playground to explore General Transit Feed Specification Reference by enabling import GTFS data sets into a MySQL database and exposing an API to query data.


Dump Database
-------------

    mysqldump --no-data -hlocalhost -ugtfs -pgtfs gtfs > gtfs-ddl.sql
    
    
Import Database
---------------

    mysql -ugtfs -pgtfs gtfs < gtfs-ddl.sql