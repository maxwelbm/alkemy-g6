FROM golang:latest

WORKDIR /usr/src/alkemy-g6

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download
