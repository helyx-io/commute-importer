FROM golang:1.4
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install
RUN mkdir /var/log/commute-importer
CMD ["go-wrapper", "run"]
EXPOSE 3000