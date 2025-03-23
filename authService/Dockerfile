FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./authService ./authService/
COPY ./libs/ ./libs/

WORKDIR /app/authService

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /auth-service ./cmd/auth/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN mkdir -p ./authService/config

COPY --from=builder /auth-service .
COPY --from=builder /app/authService/config/local.yaml /app/authService/config/config.yaml

EXPOSE 57442

CMD ["./auth-service", "--config", "/app/authService/config/config.yaml"]
