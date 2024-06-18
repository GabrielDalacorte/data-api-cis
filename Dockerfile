FROM golang:1.20-alpine

WORKDIR /app

RUN apk add --no-cache bash

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/server ./cmd/server
COPY ./internal ./internal

RUN go build -o /data-api-cis ./cmd/server

EXPOSE 8080

CMD ["sh", "-c", "sleep 10 && /data-api-cis"]
