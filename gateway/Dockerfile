FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

# Скопируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Скопируем исходный код
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway ./cmd/gateway/main.go

# Минимальный образ
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Скопируем бинарник
COPY --from=builder /app/gateway .

EXPOSE 8080

ENTRYPOINT ["./gateway", "-b"]
