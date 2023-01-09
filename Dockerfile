FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o api cmd/api/main.go
