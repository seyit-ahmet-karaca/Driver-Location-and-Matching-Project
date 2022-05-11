FROM golang:1.17 AS builder
WORKDIR /app

COPY go.mod go.mod
RUN go mod download
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

ENV APP_ENV=local_env
RUN GOOS=linux CGO_ENABLED=0 go build -o matchingService


FROM alpine:latest
WORKDIR /root

RUN mkdir .config

COPY .config/* .config/

COPY --from=builder /app/matchingService .
CMD ["./matchingService"]