FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./authService/go.mod ./authService/go.sum ./authService/
COPY libs/contracts ./libs/contracts
COPY libs/redis ./libs/redis
COPY libs/kafka ./libs/kafka

RUN cd authService && go mod download

COPY ./authService ./authService/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o auth-service ./authService/cmd/auth/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir -p ./authService/config

COPY --from=builder /app/auth-service .
COPY --from=builder /app/authService/config/local.yaml /app/authService/config/config.yaml

EXPOSE 57442

CMD ["./auth-service", "--config", "/app/authService/config/config.yaml"]
