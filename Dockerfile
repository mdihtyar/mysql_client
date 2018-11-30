FROM golang:1.11.2-stretch

RUN mkdir -p ${GOPATH}/app

WORKDIR /go/src/app

RUN go get github.com/go-sql-driver/mysql

COPY main.go ./

RUN go build

ENTRYPOINT ["./app"]
