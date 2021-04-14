FROM golang:1.16.3-alpine

ENV GO111MODULE on

WORKDIR /go/src/app

COPY app/go.mod app/go.sum ./

RUN apk add --update --no-cache ca-certificates git
RUN go mod download
RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air
