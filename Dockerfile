FROM golang:latest

WORKDIR /usr/src/alkemy-g6

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.4

COPY . .

EXPOSE 8080