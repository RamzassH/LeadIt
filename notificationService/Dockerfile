FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.work .
COPY authService ./authService
COPY gateway ./gateway
COPY libs/contracts ./libs/contracts
COPY libs/kafka ./libs/kafka
COPY notificationService ./notificationService

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notification-service ./notificationService/cmd/notification/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
RUN mkdir -p ./notificationService/config

COPY --from=builder /app/notification-service .
COPY --from=builder /app/notificationService/config/local.yaml ./notificationService/config/config.yaml

EXPOSE 57777
CMD ["./notification-service", "--config", "/app/notificationService/config/config.yaml"]
