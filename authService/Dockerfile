# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-service ./backend/cmd/auth/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir -p ./backend/config

COPY --from=builder /app/auth-service .
COPY --from=builder /app/backend/config/local.yaml ./backend/config/local.yaml

EXPOSE 50051

CMD ["./auth-service", "--config", "./backend/config/local.yaml"]
