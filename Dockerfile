FROM golang:1.17-alpine3.14 AS build
WORKDIR /app

COPY src /app/src
COPY main.go /app/main.go
COPY .env /app/.env
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

RUN go mod download
# RUN go build -o build ./src/main.go
ENTRYPOINT go run main.go httpserver