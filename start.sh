#!/bin/bash

export GTFS_BASE_URL="http://localhost/data"
export GTFS_TMP_DIR="/tmp"
export SESSION_SECRET="dsfglmfgsdhlj'130942%?"

export GOOGLE_AUTH_CLIENT_ID="1047611770782-fj4u1qo22pkqg50nca65446qcfk64loo.apps.googleusercontent.com"
export GOOGLE_AUTH_CLIENT_SECRET="pnXxoPGmqvJKerlavW2LI-Pj"

export GTFS_DB_PORT=5432
export GTFS_DB_HOSTNAME="127.0.0.1"
export GTFS_DB_DIALECT="postgres"
export GTFS_DB_DATABASE="gtfs"
export GTFS_DB_USERNAME="gtfs"
export GTFS_DB_PASSWORD="gtfs"
export REDIS_PORT=8888
export HTTP_PORT=3000
export GTFS_DB_MIN_CNX=128
export GTFS_DB_MAX_CNX=512

go get
go clean
go build
go run app.go
