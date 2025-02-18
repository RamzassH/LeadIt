FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o notificationService ./cmd/main.go

EXPOSE 50052 8081
CMD ["./notificationService"]
