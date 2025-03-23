FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./notificationService ./notificationService/
COPY ./libs/ ./libs/

WORKDIR /app/notificationService

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /notification-service ./cmd/notification/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN mkdir -p ./notificationService/config

COPY --from=builder /notification-service .

COPY --from=builder /app/notificationService/config/local.yaml /app/notificationService/config/config.yaml

EXPOSE 57777

CMD ["./notification-service", "--config", "/app/notificationService/config/config.yaml"]
