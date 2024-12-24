# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Установите необходимые пакеты
RUN apk update && apk add --no-cache git

# Установите рабочую директорию
WORKDIR /app

# Копируйте файлы зависимостей и скачайте модули
COPY go.mod go.sum ./
RUN go mod download

# Копируйте исходный код
COPY . .

# Соберите приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-service ./cmd/auth/main.go

# Stage 2: Создайте минимальный образ для запуска приложения
FROM alpine:latest

# Установите необходимые пакеты
RUN apk --no-cache add ca-certificates

# Создайте директорию для приложения
WORKDIR /root/

# Копируйте скомпилированный бинарник из предыдущего этапа
COPY --from=builder /app/auth-service .

# Экспортируйте порт, который будет использоваться приложением
EXPOSE 50051

# Запустите приложение
CMD ["./auth-service"]
