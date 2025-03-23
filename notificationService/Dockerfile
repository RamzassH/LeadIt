FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./notificationService/go.mod ./notificationService/go.sum ./notificationService/
COPY ./libs/contracts ./libs/contracts
COPY ./libs/redis ./libs/redis
COPY ./libs/kafka ./libs/kafka

RUN cd notificationService && go mod download

COPY ./notificationService ./notificationService/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
     go build -o notification-service ./notificationService/cmd/notification/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir -p ./notificationService/config

COPY --from=builder /app/notification-service .
COPY --from=builder /app/notificationService/config/local.yaml /app/notificationService/config/config.yaml

WORKDIR /app
EXPOSE 57777

CMD ["./notification-service", "--config", "/app/notificationService/config/config.yaml"]
