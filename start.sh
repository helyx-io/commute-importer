#!/bin/bash

export BASE_URL="http://localhost/data"
export TMP_DIR="/tmp"
export HTTP_PORT=3000
export SESSION_SECRET="dsfglmfgsdhlj'130942%?"

export GOOGLE_AUTH_CLIENT_ID="1047611770782-fj4u1qo22pkqg50nca65446qcfk64loo.apps.googleusercontent.com"
export GOOGLE_AUTH_CLIENT_SECRET="pnXxoPGmqvJKerlavW2LI-Pj"

export DB_PORT=5432
export DB_HOSTNAME="127.0.0.1"
export DB_DIALECT="postgres"
export DB_DATABASE="commute"
export DB_USERNAME="commute"
export DB_PASSWORD="commute"
export DB_MIN_CNX=128
export DB_MAX_CNX=512

export REDIS_PORT=8888

go get
go clean
go build
go run app.go
