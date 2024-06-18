# Etapa de construção
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

ENV DB_HOST=localhost
ENV DB_USER=admin
ENV DB_PASSWORD=admin
ENV DB_NAME=cis

EXPOSE 8080

CMD ["./main"]
