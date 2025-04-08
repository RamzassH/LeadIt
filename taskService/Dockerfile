FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./taskService ./taskService/
COPY ./libs/ ./libs/

WORKDIR /app/taskService

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /task-service ./cmd/task/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN mkdir -p ./taskService/config

COPY --from=builder /task-service .
COPY --from=builder /app/taskService/config/local.yaml /app/taskService/config/config.yaml

EXPOSE 57446

CMD ["/app/task-service", "--config", "/app/taskService/config/config.yaml"]