FROM golang:1.4-onbuild
RUN mkdir /var/log/gtfs-importer
EXPOSE 3000