#!/bin/bash
FROM golang:latest
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build
CMD ["./go-fiber-gorm-api-rest"]