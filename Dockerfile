FROM golang:latest

WORKDIR /usr/src/alkemy-g6

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.4
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080