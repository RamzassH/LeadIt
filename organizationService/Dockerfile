FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./organizationService/go.mod ./organizationService/
COPY ./organizationService/go.sum ./organizationService/
COPY libs/contracts/ ./libs/contracts/
COPY libs/redis/ ./libs/redis/
COPY libs/kafka/ ./libs/kafka/

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /organization-service ./organizationService/cmd/organization/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates
RUN mkdir -p ./organizationService/config

COPY --from=builder /app/organization-service .
COPY --from=builder /app/organizationService/config/local.yaml /app/authService/config/config.yaml

EXPOSE 57444

CMD ["/app/organization-service", "--config", "/app/organizationService/config.yaml"]