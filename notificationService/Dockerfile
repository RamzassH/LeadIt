FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY ../go.work .
COPY ../libs/contracts ./libs/contracts
COPY ../libs/redis ./libs/redis
COPY ../libs/kafka ./libs/kafka
COPY . ./

RUN go work use ./notificationService && \
    go work use ./libs/contracts && \
    go work use ./libs/redis && \
    go work use ./libs/kafka && \
    go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
     go build -o notification-service ./notificationService/cmd/notification/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/notification-service .
COPY --from=builder /app/notificationService/config/local.yaml ./notificationService/config/config.yaml

WORKDIR /app
EXPOSE 57777

CMD ["./notification-service", "--config", "/app/notificationService/config/config.yaml"]
