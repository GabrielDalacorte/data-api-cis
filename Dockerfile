FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/server ./cmd/server
COPY ./internal ./internal

RUN go build -o /data-api-cis ./cmd/server

EXPOSE 8080

CMD ["/data-api-cis"]
