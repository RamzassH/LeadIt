FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY ../go.work .
COPY . .
COPY libs/contracts ./libs/contracts
COPY libs/redis ./libs/redis
COPY libs/kafka ./libs/kafka

RUN go work use ./authService && \
    go work use ./libs/contracts && \
    go work use ./libs/redis && \
    go work use ./libs/kafka && \
    go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o auth-service ./authService/cmd/auth/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/auth-service .
COPY --from=builder /app/authService/config/local.yaml /app/authService/config/config.yaml

WORKDIR /

EXPOSE 57442

CMD ["./auth-service", "--config", "/app/authService/config/config.yaml"]
