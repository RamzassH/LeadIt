# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Установите необходимые пакеты
RUN apk update && apk add --no-cache git

# Установите рабочую директорию
WORKDIR /app

# Копируйте файлы зависимостей и скачайте модули
COPY backend/go.mod backend/go.sum ./
COPY contracts ./
RUN go mod download

# Копируйте исходный код
COPY . .

# Соберите приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-service ./backend/cmd/auth/main.go

# Stage 2: Создайте минимальный образ для запуска приложения
FROM alpine:latest

# Установите необходимые пакеты
RUN apk --no-cache add ca-certificates

# Рабочая директория
WORKDIR /root/

# Создадим необходимые директории, чтобы соответствовать пути ./backend/config/
RUN mkdir -p ./backend/config

# Скопируем бинарник
COPY --from=builder /app/auth-service .

# Скопируем config.yaml (из того же места, куда положили в builder)
COPY --from=builder /app/backend/config/local.yaml ./local.yaml

EXPOSE 50051

CMD ["./auth-service", "--config", "./backend/config/local.yaml"]
