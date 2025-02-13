# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN pwd

# Скопируем зависимости
COPY authService/backend/go.mod authService/backend/go.sum ./

COPY libs/contracts/go.mod ./libs/contracts/go.sum ./
RUN go mod download

# Скопируем исходный код
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-service ./authService/backend/cmd/auth/main.go

# Минимальный образ
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir -p ./authService/backend/config

COPY --from=builder /app/auth-service .
COPY --from=builder /app/authService/backend/config/local.yaml /app/authService/backend/config/config.yaml

EXPOSE 57442

CMD ["./auth-service", "--config", "/app/authService/backend/config/config.yaml"]
